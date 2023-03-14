import 'package:flutter/material.dart';
import 'widgets/async_calls.dart';

class Demo extends StatelessWidget {
  const Demo({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
        style: Theme.of(context).textTheme.displayLarge!,
        textAlign: TextAlign.center,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ElevatedButton(
              onPressed: () {
                showModalBottomSheet(
                    context: context,
                    builder: (BuildContext context) {
                      return psCallBody();
                    });
              },
              child: const Text('PS parallel'),
            ),
            SizedBox(width: 50),
            ElevatedButton(
              onPressed: () {
                showModalBottomSheet(
                    context: context,
                    builder: (BuildContext context) {
                      return psBatchCallBody();
                    });
              },
              child: const Text('Batch parallel'),
            ),
          ],
        ));
  }
}
