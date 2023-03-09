import 'package:flutter/material.dart';

import 'data/storage.dart';
import 'plugins/ps_demo.dart';
import 'app.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  String device_path = await Storage.init();
  FFIBridge.initialize(device_path);

  runApp(
    const MaterialApp(
      home: Scaffold(
        body: Center(
          child: Demo(),
        ),
      ),
    ),
  );
}
