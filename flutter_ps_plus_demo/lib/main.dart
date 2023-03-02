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
  final stopwatch = Stopwatch()..start();
  Pointer<ArrayOfStrings> result =
      FFIBridge.c_ps_demo(device_path.toNativeUtf8());
  await Future.delayed(const Duration(seconds: 1))!;
  final timing = stopwatch.elapsed;
  print('ps_demo executed in ${timing}');
  final n = result.ref.num_arrays;
  final params = List.empty(growable: true);
  for (var i = 0; i < n; i++) {
    params.add(result.ref.array.elementAt(i).value.toDartString());
  }
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

void main() {
  //FFIBridge.initialize();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;
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

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return FutureBuilder<String>(
      future: _localPath,
      builder: (BuildContext context, AsyncSnapshot<String> snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        }
        if (snapshot.hasData) {
          String local_path = snapshot.data!;

          return Column(
            children: <Widget>[
              FutureBuilder<String>(
                future: ps_demo(local_path),
                builder:
                    (BuildContext context, AsyncSnapshot<String> snapshot) {
                  if (snapshot.hasData) {
                    String _result = snapshot.data!;
                    return Text('Result: ${snapshot.data}');
                  } else {
                    return Text('Executing ps_demo ...');
                  }
                },
              ),
              ElevatedButton(
                onPressed: _retry,
                child: Text('Retry'),
              )
            ],
          );

          // return Scaffold(
          //   appBar: AppBar(
          //     // Here we take the value from the MyHomePage object that was created by
          //     // the App.build method, and use it to set our appbar title.
          //     title: Text(widget.title),
          //   ),
          //   body: Center(
          //     // Center is a layout widget. It takes a single child and positions it
          //     // in the middle of the parent.
          //     child: Column(
          //         mainAxisAlignment: MainAxisAlignment.center,
          //         children: [
          //           Text(
          //               'capitalize flutter=${FFIBridge.capitalize('flutter')}',
          //               style: TextStyle(fontSize: 40)),
          //           Text('1+2=${FFIBridge.add(1, 2)}',
          //               style: TextStyle(fontSize: 40)),
          //           Text('From C = ${FFIBridge.helloworld()}',
          //               style: TextStyle(fontSize: 40)),
          //           Text(
          //               'params = ${FFIBridge.get_some_parameters(local_path)}',
          //               style: TextStyle(fontSize: 10)),
          //         ]),
          //     // child: Column(
          //     //   // Column is also a layout widget. It takes a list of children and
          //     //   // arranges them vertically. By default, it sizes itself to fit its
          //     //   // children horizontally, and tries to be as tall as its parent.
          //     //   //
          //     //   // Invoke "debug painting" (press "p" in the console, choose the
          //     //   // "Toggle Debug Paint" action from the Flutter Inspector in Android
          //     //   // Studio, or the "Toggle Debug Paint" command in Visual Studio Code)
          //     //   // to see the wireframe for each widget.
          //     //   //
          //     //   // Column has various properties to control how it sizes itself and
          //     //   // how it positions its children. Here we use mainAxisAlignment to
          //     //   // center the children vertically; the main axis here is the vertical
          //     //   // axis because Columns are vertical (the cross axis would be
          //     //   // horizontal).
          //     //   mainAxisAlignment: MainAxisAlignment.center,

          //     //   children: <Widget>[
          //     //     //Text('capitalize flutter=${FFIBridge.capitalize('flutter')}'),
          //     //     //Text('1+2=${FFIBridge.add(1, 2)}'),
          //     //     const Text(
          //     //       'You have pushed the button this many times:',
          //     //     ),
          //     //     Text(
          //     //       '$_counter',
          //     //       style: Theme.of(context).textTheme.headlineMedium,
          //     //     ),
          //     //   ],
          //     // ),
          //   ),
          //   floatingActionButton: FloatingActionButton(
          //     onPressed: _incrementCounter,
          //     tooltip: 'Increment',
          //     child: const Icon(Icons.add),
          //   ), // This trailing comma makes auto-formatting nicer for build methods.
          // );
        }
        return const Text("no data");
      },
    );
  }
}
