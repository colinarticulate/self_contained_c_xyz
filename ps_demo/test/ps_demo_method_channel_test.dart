import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:ps_demo/ps_demo_method_channel.dart';

void main() {
  MethodChannelPsDemo platform = MethodChannelPsDemo();
  const MethodChannel channel = MethodChannel('ps_demo');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await platform.getPlatformVersion(), '42');
  });
}
