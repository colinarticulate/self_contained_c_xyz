import 'package:flutter/material.dart';
import 'dart:ffi';
import 'package:ffi/ffi.dart';
import 'package:flutter/services.dart';
import 'package:path_provider/path_provider.dart';
import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'package:path/path.dart' as p;

class ArrayOfStrings extends Struct {
  external Pointer<Pointer<Utf8>> array;

  @Int32()
  external int num_arrays;
}

Future<String> get _localPath async {
  // final directory = await getApplicationDocumentsDirectory();
  //final directory = await getApplicationSupportDirectory();
  final directory = Platform.isAndroid
      ? await getExternalStorageDirectory()
      : await getApplicationSupportDirectory();
  // final directory =
  //     await getApplicationSupportDirectory(); //Both Linux and iOS
  print('Persistent data path: $directory');

  //get info about the assets
  final manifestContent = await rootBundle.loadString('AssetManifest.json');

  final Map<String, dynamic> manifestMap = json.decode(manifestContent);
  // >> To get paths you need these 2 lines

  //Choose specifice assets
  final imagePaths = manifestMap.keys
      .where((String key) => key.contains('assets/ps_plus/'))
      .toList();
  //print(imagePaths);

  //final persistent_path = directory.path;
  //String asset_path = imagePaths[0];

  for (var i = 0; i < imagePaths.length; i++) {
    final pathfile =
        p.join(directory!.path, imagePaths[i].replaceAll("assets/", ""));
    final byteData = await rootBundle.load(imagePaths[i]);
    print('-->: $pathfile');
    final file = await File(pathfile).create(recursive: true);
    await file.writeAsBytes(byteData.buffer
        .asUint8List(byteData.offsetInBytes, byteData.lengthInBytes));
  }

  return directory!.path;
  //return "/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/data/";
}

Future<String> ps_demo(String device_path) async {
  device_path =
      "/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/data/";
  final c_path = device_path.toNativeUtf8();
  print(c_path);
  print(c_path.toDartString());
  final stopwatch = Stopwatch()..start();
  Pointer<ArrayOfStrings> result = FFIBridge.c_ps_demo(c_path);
  //await Future.delayed(const Duration(seconds: 1))!;
  final timing = stopwatch.elapsed;
  print('ps_demo executed in ${timing}');
  final n = result.ref.num_arrays;
  final params = List.empty(growable: true);
  for (var i = 0; i < n; i++) {
    params.add(result.ref.array.elementAt(i).value.toDartString());
  }

  malloc.free(c_path);
  //now we could delete result invoking delete method

  return "ps_demo finished in ${timing} ";
}

class FFIBridge {
  static bool initialize() {
    nativeApiLib = (DynamicLibrary.open('libps_plus.so')); // android and linux

    final _add = nativeApiLib
        .lookup<NativeFunction<Int32 Function(Int32, Int32)>>('add');
    add = _add.asFunction<int Function(int, int)>();

    final _cap = nativeApiLib.lookup<
        NativeFunction<Pointer<Utf8> Function(Pointer<Utf8>)>>('capitalize');
    _capitalize = _cap.asFunction<Pointer<Utf8> Function(Pointer<Utf8>)>();

    final __helloWorld = nativeApiLib
        .lookup<NativeFunction<Pointer<Utf8> Function()>>('hello_world');
    _helloworld = __helloWorld.asFunction<Pointer<Utf8> Function()>();

    final _some_parameters = nativeApiLib.lookup<
            NativeFunction<Pointer<ArrayOfStrings> Function(Pointer<Utf8>)>>(
        'get_some_parameters');
    c_get_some_parameters = _some_parameters
        .asFunction<Pointer<ArrayOfStrings> Function(Pointer<Utf8>)>();

    final _ps_demo = nativeApiLib.lookup<
            NativeFunction<Pointer<ArrayOfStrings> Function(Pointer<Utf8>)>>(
        'ps_demo');
    c_ps_demo =
        _ps_demo.asFunction<Pointer<ArrayOfStrings> Function(Pointer<Utf8>)>();

    return true;
  }

  static late DynamicLibrary nativeApiLib;
  static late Function add;
  static late Function _capitalize;
  static late Function _helloworld;
  static late Function c_get_some_parameters;
  static late Function c_ps_demo;

  static String capitalize(String str) {
    final _str = str.toNativeUtf8();
    Pointer<Utf8> res = _capitalize(_str);
    calloc.free(_str);
    return res.toDartString();
  }

  static String helloworld() {
    Pointer<Utf8> res = _helloworld();
    return res.toDartString();
  }

  static List<dynamic> get_some_parameters(String path) {
    Pointer<ArrayOfStrings> result = c_get_some_parameters(path.toNativeUtf8());
    // print(result);
    // print(result.ref);
    // print(result.ref.num_arrays);
    // print(result.ref.array.elementAt(0).value.toDartString());
    final n = result.ref.num_arrays;
    final params = List.empty(growable: true);
    for (var i = 0; i < n; i++) {
      params.add(result.ref.array.elementAt(i).value.toDartString());
    }

    return params;
  }
}

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  static const String _title = 'Flutter Code Sample';

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      title: _title,
      home: MyStatefulWidget(),
    );
  }
}

class MyStatefulWidget extends StatefulWidget {
  const MyStatefulWidget({super.key});

  @override
  State<MyStatefulWidget> createState() => _MyStatefulWidgetState();
}

class _MyStatefulWidgetState extends State<MyStatefulWidget> {
  final Future<String> _calculation = Future<String>.delayed(
    const Duration(seconds: 2),
    () => 'Data Loaded',
  );

  String local_path = "";
  //late Future<String> _result;
  String _result = "";

  @override
  void initState() {
    FFIBridge.initialize();
    _localPath;
    super.initState();
  }

  void _retry() {
    setState(() {
      ps_demo(local_path);
    });
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
      style: Theme.of(context).textTheme.displayMedium!,
      textAlign: TextAlign.center,
      child: FutureBuilder<String>(
        future: _calculation, // a previously-obtained Future<String> or null
        builder: (BuildContext context, AsyncSnapshot<String> snapshot) {
          List<Widget> children;
          if (snapshot.hasData) {
            children = <Widget>[
              const Icon(
                Icons.check_circle_outline,
                color: Colors.green,
                size: 60,
              ),
              Padding(
                padding: const EdgeInsets.only(top: 16),
                child: Text('Result: ${snapshot.data}'),
              ),
            ];
          } else if (snapshot.hasError) {
            children = <Widget>[
              const Icon(
                Icons.error_outline,
                color: Colors.red,
                size: 60,
              ),
              Padding(
                padding: const EdgeInsets.only(top: 16),
                child: Text('Error: ${snapshot.error}'),
              ),
              ElevatedButton(
                onPressed: _retry,
                child: Text('Retry'),
              )
            ];
          } else {
            children = const <Widget>[
              SizedBox(
                width: 60,
                height: 60,
                child: CircularProgressIndicator(),
              ),
              Padding(
                padding: EdgeInsets.only(top: 16),
                child: Text('Awaiting result...'),
              ),
            ];
          }
          return Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: children,
            ),
          );
        },
      ),
    );
  }
}
