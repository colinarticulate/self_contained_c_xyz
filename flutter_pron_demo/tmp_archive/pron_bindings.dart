// AUTO GENERATED FILE, DO NOT EDIT.
//
// Generated by `package:ffigen`.
// ignore_for_file: type=lint
import 'dart:convert';
import 'dart:ffi' as ffi;
import 'dart:ffi';
import 'package:ffi/ffi.dart';

/// Bindings to flutter_pron
class PronGO {
  /// Holds the symbol lookup function.
  final ffi.Pointer<T> Function<T extends ffi.NativeType>(String symbolName)
      _lookup;

  /// The symbols are looked up in [dynamicLibrary].
  PronGO(ffi.DynamicLibrary dynamicLibrary) : _lookup = dynamicLibrary.lookup;

  /// The symbols are looked up with [lookup].
  PronGO.fromLookup(
      ffi.Pointer<T> Function<T extends ffi.NativeType>(String symbolName)
          lookup)
      : _lookup = lookup;

  GoString Pron(
    GoString audiofile,
    GoString word,
    GoString outputfolder,
    GoString dictfile,
    GoString phdictfile,
    GoString featparams,
    GoString hmm,
    GoString proffile,
  ) {
    return _Pron(
      audiofile,
      word,
      outputfolder,
      dictfile,
      phdictfile,
      featparams,
      hmm,
    );
  }

  late final _PronPtr = _lookup<
      ffi.NativeFunction<
          GoString Function(GoString, GoString, GoString, GoString, GoString,
              GoString, GoString)>>('Pron');
  late final _Pron = _PronPtr.asFunction<
      GoString Function(GoString, GoString, GoString, GoString, GoString,
          GoString, GoString)>();

  GoString MockPron(
    GoString audiofile,
    GoString word,
    GoString outputfolder,
    GoString dictfile,
    GoString phdictfile,
    GoString featparams,
    GoString hmm,
    GoString proffile,
  ) {
    return _Pron(
      audiofile,
      word,
      outputfolder,
      dictfile,
      phdictfile,
      featparams,
      hmm,
    );
  }

  late final _MockPronPtr = _lookup<
      ffi.NativeFunction<
          GoString Function(GoString, GoString, GoString, GoString, GoString,
              GoString, GoString)>>('MockPron');
  late final _MockPron = _PronPtr.asFunction<
      GoString Function(GoString, GoString, GoString, GoString, GoString,
          GoString, GoString)>();
}

class max_align_t extends ffi.Opaque {}

class _GoString_ extends ffi.Struct {
  external ffi.Pointer<ffi.Char> p;

  @ptrdiff_t()
  external int n;
}

typedef ptrdiff_t = ffi.Long;

class GoInterface extends ffi.Struct {
  external ffi.Pointer<ffi.Void> t;

  external ffi.Pointer<ffi.Void> v;
}

class GoSlice extends ffi.Struct {
  external ffi.Pointer<ffi.Void> data;

  @GoInt()
  external int len;

  @GoInt()
  external int cap;
}

typedef GoInt = GoInt64;
typedef GoInt64 = ffi.LongLong;
typedef GoString = _GoString_;

const int NULL = 0;

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

//We could also add slices, but we currently don't need them
