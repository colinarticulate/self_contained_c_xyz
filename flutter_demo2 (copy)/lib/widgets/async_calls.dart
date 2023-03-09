import 'package:flutter/material.dart';
import 'package:flutter_demo2/repo/plugin_repository.dart';

class HeavyCallBody extends StatelessWidget {
  HeavyCallBody({super.key});
  final PluginRepository _psPlugin = PluginRepository();

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<String>(
        future: _psPlugin
            .heavyCall(), // a previously-obtained Future<String> or null
        builder: (BuildContext context, AsyncSnapshot<String> snapshot) {
          //List<Widget> children;
          if (snapshot.hasData) {
            return Padding(
              padding: const EdgeInsets.all(8.0),
              child: Center(child: Text('Returned: ${snapshot.data!}')),
            );
          } else if (snapshot.hasError) {
            return Padding(
              padding: const EdgeInsets.all(8.0),
              child: Center(child: Text('Error: ${snapshot.error}')),
            );
          } else {
            return Padding(
              padding: const EdgeInsets.all(8.0),
              child: Center(child: CircularProgressIndicator()),
            );
          }
        });
  }
}
