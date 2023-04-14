import 'package:flutter/material.dart';
import 'package:flutter_pron_demo/repo/plugin_repository.dart';

// import 'data/storage.dart';
// import 'plugins/ps_demo.dart';
import 'app.dart';

void main() async {
  // PluginRepository();
  WidgetsFlutterBinding.ensureInitialized();

  runApp(
    MaterialApp(
      home: Scaffold(
        body: Center(
          child: Demo(),
        ),
      ),
    ),
  );
}
