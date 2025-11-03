import 'dart:convert';
import 'dart:ffi';

import 'package:dalat/dalat.dart';
import 'package:dalat/src/bindings.dart';
import 'package:ffi/ffi.dart';

mixin InstanceHelpers {
  int get handle;

  static String getAlatError() {
    final msgPointer = bindings.get_error();
    final error = msgPointer == nullptr
        ? "Unknown error"
        : msgPointer.cast<Utf8>().toDartString();
    bindings.free_string(msgPointer);
    return error;
  }

  Future<T> jsonGetterHelper<T>(
    Pointer<Char> Function(int) ffiFunc,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    return fromJson(jsonDecode(await getterHelper(ffiFunc)));
  }

  Future<String> getterHelper<T>(Pointer<Char> Function(int) ffiFunc) async {
    final ptr = ffiFunc(handle);
    if (ptr == nullptr) {
      throw Exception(
        'Failed to get data from Go core: function returned null pointer. ${getAlatError()}',
      );
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      return jsonStr;
    } finally {
      bindings.free_string(ptr);
    }
  }

  Future<void> helper(int Function(int) ffiFunc) async {
    final result = ffiFunc(handle);
    if (result < 0) {
      throw Exception(
        'Failed to set data in Go core. Code: $result ${getAlatError()}',
      );
    }
  }

  Future<List<T>> jsonListGetterHelper<T>(
    Pointer<Char> Function(int) ffiFunc,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    final ptr = ffiFunc(handle);
    if (ptr == nullptr) {
      // An empty list is represented by a null pointer in this API
      return [];
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final List<dynamic> decoded = jsonDecode(jsonStr) ?? [];
      return decoded
          .map((item) => fromJson(item as Map<String, dynamic>))
          .toList();
    } finally {
      bindings.free_string(ptr);
    }
  }

  Future<void> jsonSetterHelper(
    dynamic jsonData,
    int Function(int, Pointer<Char>) ffiFunc,
  ) async {
    return setterHelper(jsonEncode(jsonData), ffiFunc);
  }

  Future<void> setterHelper(
    String data,
    int Function(int, Pointer<Char>) ffiFunc,
  ) async {
    final jsonStrC = data.toNativeUtf8();
    try {
      final result = ffiFunc(handle, jsonStrC.cast());
      if (result < 0) {
        throw Exception(
          'Failed to set data in Go core. Code: $result ${getAlatError()}',
        );
      }
    } finally {
      malloc.free(jsonStrC);
    }
  }

  Future<List<DeviceColor>> getAlatColors() {
    final ptr = bindings.get_alat_device_colors_json();
    if (ptr == nullptr) {
      return Future.value([]);
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final List<dynamic> decoded = jsonDecode(jsonStr);
      return Future.value(
        decoded
            .map((item) => DeviceColor.fromJson(item as Map<String, dynamic>))
            .toList(),
      );
    } finally {
      bindings.free_string(ptr);
    }
  }
}
