import 'package:flutter/material.dart';

// import 'data/storage.dart';
// import 'plugins/ps_demo.dart';
import 'app.dart';

void main() async {
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
