import 'dart:ffi';
import 'dart:io';
import 'package:ffi/ffi.dart';

// This should be in another file
class ArrayOfStrings extends Struct {
  external Pointer<Pointer<Utf8>> array;

  @Int32()
  external int num_arrays;
}

class FFIBridge {
  static String _path = "";

  // This should be:
  // static Future<bool>, so we could check if we have all the files we need
  // so the plugin doesn't crash, return False if something is missing.
  static bool initialize(String device_path) {
    _path =
        device_path + "/ps_plus/"; // this need fixing to make it fault-proof.
    print('demo path from FFIBridge: ${_path}');

    //nativeApiLib = Platform.isMacOS || Platform.isIOS ? DynamicLibrary.process() : (DynamicLibrary.open('libapi.so')); // android and linux
    //nativeApiLib = DynamicLibrary.open('libapi.dylib');
    nativeApiLib = DynamicLibrary.open('libps_plus.dylib');

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

    final _ps_batch_demo = nativeApiLib.lookup<
            NativeFunction<Pointer<ArrayOfStrings> Function(Pointer<Utf8>)>>(
        'ps_batch_demo');
    c_ps_batch_demo = _ps_batch_demo
        .asFunction<Pointer<ArrayOfStrings> Function(Pointer<Utf8>)>();

    return true;
  }

  static late DynamicLibrary nativeApiLib;
  static late Function add;
  static late Function _capitalize;
  static late Function _helloworld;
  static late Function c_get_some_parameters;
  static late Function c_ps_demo;
  static late Function c_ps_batch_demo;

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

  static List<dynamic> get_some_parameters() {
    Pointer<ArrayOfStrings> result =
        c_get_some_parameters(_path.toNativeUtf8());
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

  static Future<String> ps_demo() async {
    final c_path = _path.toNativeUtf8();
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

    return "${timing}";
  }

  static Future<String> ps_batch_demo() async {
    final c_path = _path.toNativeUtf8();
    print(c_path);
    print(c_path.toDartString());
    final stopwatch = Stopwatch()..start();
    Pointer<ArrayOfStrings> result = FFIBridge.c_ps_batch_demo(c_path);
    //await Future.delayed(const Duration(seconds: 1))!;
    final timing = stopwatch.elapsed;
    print('ps_BATCH_demo executed in ${timing}');
    final n = result.ref.num_arrays;
    final params = List.empty(growable: true);
    for (var i = 0; i < n; i++) {
      params.add(result.ref.array.elementAt(i).value.toDartString());
    }

    malloc.free(c_path);
    //now we could delete result invoking delete method

    return "${timing}";
  }
} // FFIBridge
//------------------------------------------------------------