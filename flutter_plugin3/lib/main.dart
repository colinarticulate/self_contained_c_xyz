import 'package:flutter/material.dart';
import 'package:flutter_plugin3/repo/plugin_repository.dart';
import 'package:get_it/get_it.dart';

// import 'data/storage.dart';
// import 'plugins/ps_demo.dart';
import 'app.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  GetIt.I.registerSingleton<PluginRepository>(PluginRepository());
  await GetIt.I<PluginRepository>().init();

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
