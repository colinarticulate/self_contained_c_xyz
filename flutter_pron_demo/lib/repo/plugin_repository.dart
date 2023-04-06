import '../data/storage.dart';
import '../plugins/ps_demo.dart';
import 'dart:async';
import 'dart:isolate';

class PluginRepository {
  late final String devicePath;
  Completer storageInitCompleted = Completer();

  PluginRepository() {
    init();
  }

  void init() async {
    devicePath = await Storage.init();
    storageInitCompleted.complete();
  }

  Future<String> psCall() async {
    await storageInitCompleted.future;
    ReceivePort receivePort = ReceivePort();
    Completer<String> completer = Completer<String>();
    Isolate isolate =
        await Isolate.spawn(_psCallIsolateFunction, receivePort.sendPort);
    receivePort.listen((data) {
      completer.complete(data);
      isolate.kill(priority: Isolate.immediate);
      receivePort.close();
    });
    return completer.future;
  }

  Future<String> _psCallDirect() async {
    return FFIBridge.ps_demo();
    // return mockHeavyCall();
  }

  void _psCallIsolateFunction(SendPort sendPort) async {
    FFIBridge.initialize(devicePath);
    String result = await _psCallDirect();
    sendPort.send(result);
  }

  //psBatchCall

  Future<String> psBatchCall() async {
    await storageInitCompleted.future;
    ReceivePort receivePort = ReceivePort();
    Completer<String> completer = Completer<String>();
    Isolate isolate =
        await Isolate.spawn(_psBatchCallIsolateFunction, receivePort.sendPort);
    receivePort.listen((data) {
      completer.complete(data);
      isolate.kill(priority: Isolate.immediate);
      receivePort.close();
    });
    return completer.future;
  }

  void _psBatchCallIsolateFunction(SendPort sendPort) async {
    FFIBridge.initialize(devicePath);
    String result = await _psBatchCallDirect();
    sendPort.send(result);
  }

  Future<String> _psBatchCallDirect() async {
    return FFIBridge.ps_batch_demo();
    // return mockHeavyCall();
  }

  Future<String> mockHeavyCall() async {
    return Future<String>.delayed(
      const Duration(seconds: 2),
      () => 'Data Loaded ', //+FFIBridge.add(2,3).toString(),
    );
  }

//Pron
  Future<String> _pronCallDirect() async {
    return FFIBridge.pron_demo();
    // return mockHeavyCall();
  }

  void _pronCallIsolateFunction(SendPort sendPort) async {
    FFIBridge.initialize(devicePath);
    String result = await _pronCallDirect();
    sendPort.send(result);
  }

  Future<String> pronCall() async {
    await storageInitCompleted.future;
    ReceivePort receivePort = ReceivePort();
    Completer<String> completer = Completer<String>();
    Isolate isolate =
        await Isolate.spawn(_pronCallIsolateFunction, receivePort.sendPort);
    receivePort.listen((data) {
      completer.complete(data);
      isolate.kill(priority: Isolate.immediate);
      receivePort.close();
    });
    return completer.future;
  }
}
