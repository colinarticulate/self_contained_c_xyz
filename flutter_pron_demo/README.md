# flutter_plugin3

A new Flutter project.

## Getting Started

This project is a starting point for a Flutter application.

A few resources to get you started if this is your first Flutter project:

- [Lab: Write your first Flutter app](https://docs.flutter.dev/get-started/codelab)
- [Cookbook: Useful Flutter samples](https://docs.flutter.dev/cookbook)

For help getting started with Flutter development, view the
[online documentation](https://docs.flutter.dev/), which offers tutorials,
samples, guidance on mobile development, and a full API reference.


Once ffigen generates flutter bindings for Go, add the following:
//These are the only 3 things that needed to be added:
Pointer<GoString> toGoString(String dartString) {
//https://dart.dev/tools/diagnostic-messages?utm_source=dartdev&utm_medium=redir&utm_id=diagcode&utm_content=creation_of_struct_or_union#creation_of_struct_or_union
  final pointer = calloc.allocate<GoString>(sizeOf<GoString>());
  pointer.ref.p = dartString.toNativeUtf8().cast<Char>();
  pointer.ref.n = dartString.length;
  // return pointer.ref;
  return pointer;
}

void freeGoString(Pointer<GoString> goString) {
  malloc.free(goString.ref.p);
  calloc.free(goString);
}

String fromGoString(GoString goString) {
  final string = goString.p.cast<Utf8>().toDartString(length: goString.n);
  return string;
}
