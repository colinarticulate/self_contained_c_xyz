import '../plugins/ps_demo.dart';

class PluginRepository {
  Future<String> mockHeavyCall() async {
    return Future<String>.delayed(
      const Duration(seconds: 2),
      () => 'Data Loaded ',//+FFIBridge.add(2,3).toString(),
    );
  }

  Future<String> psCall() async {
    return FFIBridge.ps_demo();
    // return mockHeavyCall();
  }

  Future<String> psBatchCall() async {
    return FFIBridge.ps_batch_demo();
    // return mockHeavyCall();
  }
}
