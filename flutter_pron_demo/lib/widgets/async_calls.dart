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

class pronCallBody extends StatelessWidget {
  pronCallBody({super.key});
  final PluginRepository _psPlugin = PluginRepository();

  @override
  Widget build(BuildContext context) {
    // return DefaultTextStyle(
    //     style: Theme.of(context).textTheme.displaySmall!,
    //     textAlign: TextAlign.center,
    return new Container(
        width: MediaQuery.of(context).size.width * 0.8,
        height: MediaQuery.of(context).size.height * 0.8,
        child: FutureBuilder<String>(
            future: _psPlugin
                .pronCall(), // a previously-obtained Future<String> or null
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
