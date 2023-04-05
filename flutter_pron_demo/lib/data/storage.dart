import 'dart:convert';
import 'dart:io';
import 'dart:async';

import 'package:flutter/services.dart';
import 'package:path_provider/path_provider.dart';
import 'package:path/path.dart' as p;

class Storage {
  static String path = "";
  static String _demo_path = "";

  //This should be static Future<bool>, and we check that the directory exists
  //in the device returning False if not
  static Future<String> init() async {
    // final directory = await getApplicationDocumentsDirectory();
    //final directory = await getApplicationSupportDirectory();
    final directory = Platform.isAndroid
        ? await getExternalStorageDirectory()
        : await getApplicationSupportDirectory();
    // final directory =
    //     await getApplicationSupportDirectory(); //Both Linux and iOS
    print('Persistent data path: $directory');

    //get info about the assets
    final manifestContent = await rootBundle.loadString('AssetManifest.json');

    final Map<String, dynamic> manifestMap = json.decode(manifestContent);
    // >> To get paths you need these 2 lines

    //Choose specifice assets
    final imagePaths = manifestMap.keys
        .where((String key) => key.contains('assets/ps_plus/'))
        .toList();
    //print(imagePaths);

    //final persistent_path = directory.path;
    //String asset_path = imagePaths[0];

    //new Directory(p.join(directory!.path,"ps_plus")).create();
    for (var i = 0; i < imagePaths.length; i++) {
      final pathfile =
          p.join(directory!.path, imagePaths[i].replaceAll("assets/", ""));
      final byteData = await rootBundle.load(imagePaths[i]);
      print('-->: $pathfile');
      final file = await File(pathfile).create(recursive: true);
      await file.writeAsBytes(byteData.buffer
          .asUint8List(byteData.offsetInBytes, byteData.lengthInBytes));
    }
    path = directory!.path;

    return path;
    //return "/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/data/";
  }
}
