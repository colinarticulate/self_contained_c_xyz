import 'package:flutter/material.dart';
import '../repo/plugin_repository.dart';

class psCallBody extends StatelessWidget {
  psCallBody({super.key});
  final PluginRepository _psPlugin = PluginRepository();

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
        style: Theme.of(context).textTheme.displayMedium!,
        textAlign: TextAlign.center,
        child: FutureBuilder<String>(
            future: _psPlugin
                .psCall(), // a previously-obtained Future<String> or null
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
            }));
  }
}

class psBatchCallBody extends StatelessWidget {
  psBatchCallBody({super.key});
  final PluginRepository _psPlugin = PluginRepository();

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
        style: Theme.of(context).textTheme.displayMedium!,
        textAlign: TextAlign.center,
        child: FutureBuilder<String>(
            future: _psPlugin
                .psBatchCall(), // a previously-obtained Future<String> or null
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
            }));
  }
}

class psMockCallBody extends StatelessWidget {
  psMockCallBody({super.key});
  final PluginRepository _psPlugin = PluginRepository();

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
        style: Theme.of(context).textTheme.displayMedium!,
        textAlign: TextAlign.center,
        child: FutureBuilder<String>(
            future: _psPlugin
                .mockHeavyCall(), // a previously-obtained Future<String> or null
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
            }));
  }
}
