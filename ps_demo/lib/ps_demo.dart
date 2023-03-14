
import 'ps_demo_platform_interface.dart';

class PsDemo {
  Future<String?> getPlatformVersion() {
    return PsDemoPlatform.instance.getPlatformVersion();
  }
}
