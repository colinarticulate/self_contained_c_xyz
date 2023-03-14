import 'package:plugin_platform_interface/plugin_platform_interface.dart';

import 'ps_demo_method_channel.dart';

abstract class PsDemoPlatform extends PlatformInterface {
  /// Constructs a PsDemoPlatform.
  PsDemoPlatform() : super(token: _token);

  static final Object _token = Object();

  static PsDemoPlatform _instance = MethodChannelPsDemo();

  /// The default instance of [PsDemoPlatform] to use.
  ///
  /// Defaults to [MethodChannelPsDemo].
  static PsDemoPlatform get instance => _instance;

  /// Platform-specific implementations should set this with their own
  /// platform-specific class that extends [PsDemoPlatform] when
  /// they register themselves.
  static set instance(PsDemoPlatform instance) {
    PlatformInterface.verifyToken(instance, _token);
    _instance = instance;
  }

  Future<String?> getPlatformVersion() {
    throw UnimplementedError('platformVersion() has not been implemented.');
  }
}
