import 'dart:convert';
import 'dart:ffi';
import 'dart:io';
import 'package:ffi/ffi.dart';

class Binding<P, R> {
  late final Pointer<Utf8> Function(Pointer<Utf8>) _func;
  Binding(DynamicLibrary lib, String name) {
    _func = lib
        .lookup<NativeFunction<Pointer<Utf8> Function(Pointer<Utf8>)>>(name)
        .asFunction();
  }
  R call(P args) {
    final jsonString = jsonEncode(args);
    final nativePointer = jsonString.toNativeUtf8();
    final resultPtr = _func(nativePointer);
    final result = resultPtr.toDartString();
    return jsonDecode(result) as R;
  }
}

class ApiBridge {
  late final DynamicLibrary lib;

  late final Binding<Map, List> searchDevices;

  ApiBridge() {
    lib = Platform.isAndroid
        ? DynamicLibrary.open('libalat.so')
        : DynamicLibrary.open('build/libalat.so'); // for linux testing

    searchDevices = Binding(lib, 'SearchDevices');
  }
}
