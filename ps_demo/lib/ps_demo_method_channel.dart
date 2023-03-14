import 'package:flutter/foundation.dart';
import 'package:flutter/services.dart';

import 'ps_demo_platform_interface.dart';

/// An implementation of [PsDemoPlatform] that uses method channels.
class MethodChannelPsDemo extends PsDemoPlatform {
  /// The method channel used to interact with the native platform.
  @visibleForTesting
  final methodChannel = const MethodChannel('ps_demo');

  @override
  Future<String?> getPlatformVersion() async {
    final version = await methodChannel.invokeMethod<String>('getPlatformVersion');
    return version;
  }
}
