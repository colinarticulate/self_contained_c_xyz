import 'package:flutter/material.dart';
import 'widgets/async_calls.dart';

class Demo extends StatelessWidget {
  const Demo({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
        style: Theme.of(context).textTheme.displayLarge!,
        textAlign: TextAlign.center,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ElevatedButton(
              onPressed: () {
                showModalBottomSheet(
                    isScrollControlled: true,
                    context: context,
                    builder: (BuildContext context) {
                      return pronCallBody();
                    });
              },
              child: const Text('pron'),
            ),
          ],
        ));
  }
}
