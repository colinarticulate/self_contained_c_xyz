import '../plugins/ps_demo.dart';

class PluginRepository {
  Future<String> mockHeavyCall() async {
    return Future<String>.delayed(
      const Duration(seconds: 2),
      () => 'Data Loaded',
    );
  }

  Future<String> psCall() async {
    return FFIBridge.ps_demo();
  }

  Future<String> psBatchCall() async {
    return FFIBridge.ps_batch_demo();
  }
}
