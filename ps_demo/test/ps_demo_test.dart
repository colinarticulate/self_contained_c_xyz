import 'package:flutter_test/flutter_test.dart';
import 'package:ps_demo/ps_demo.dart';
import 'package:ps_demo/ps_demo_platform_interface.dart';
import 'package:ps_demo/ps_demo_method_channel.dart';
import 'package:plugin_platform_interface/plugin_platform_interface.dart';

class MockPsDemoPlatform
    with MockPlatformInterfaceMixin
    implements PsDemoPlatform {

  @override
  Future<String?> getPlatformVersion() => Future.value('42');
}

void main() {
  final PsDemoPlatform initialPlatform = PsDemoPlatform.instance;

  test('$MethodChannelPsDemo is the default instance', () {
    expect(initialPlatform, isInstanceOf<MethodChannelPsDemo>());
  });

  test('getPlatformVersion', () async {
    PsDemo psDemoPlugin = PsDemo();
    MockPsDemoPlatform fakePlatform = MockPsDemoPlatform();
    PsDemoPlatform.instance = fakePlatform;

    expect(await psDemoPlugin.getPlatformVersion(), '42');
  });
}
