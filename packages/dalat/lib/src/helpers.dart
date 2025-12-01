import 'dart:convert';
import 'dart:ffi';

import 'package:dalat/dalat.dart';
import 'package:dalat/src/bindings.dart';
import 'package:dalat/src/isolate_helper.dart';
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
    String methodName,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    final jsonStr = await FfiIsolate.run(methodName, [handle]) as String?;
    if (jsonStr == null) {
      throw Exception(
          'Failed to get data from Go core: function returned null pointer. ${getAlatError()}');
    }
    return fromJson(jsonDecode(jsonStr));
  }

  Future<void> helper(String methodName) async {
    await FfiIsolate.run(methodName, [handle]);
  }

  Future<List<T>> jsonListGetterHelper<T>(
    String methodName,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    final jsonStr = await FfiIsolate.run(methodName, [handle]) as String?;
    if (jsonStr == null) {
      // An empty list can be represented by a null pointer.
      return [];
    }
    final List<dynamic> decoded = jsonDecode(jsonStr) ?? [];
    return decoded
        .map((item) => fromJson(item as Map<String, dynamic>))
        .toList();
  }

  Future<void> jsonSetterHelper(
    String methodName,
    dynamic jsonData,
  ) async {
    await setterHelper(methodName, jsonEncode(jsonData));
  }

  Future<void> setterHelper(
    String methodName,
    String data,
  ) async {
    await FfiIsolate.run(methodName, [handle, data]);
  }

  Future<List<DeviceColor>> getAlatColors() async {
    final jsonStr =
        await FfiIsolate.run('get_alat_device_colors_json', []) as String?;
    if (jsonStr == null) {
      return [];
    }
    final List<dynamic> decoded = jsonDecode(jsonStr) ?? [];
    return decoded
        .map((item) => DeviceColor.fromJson(item as Map<String, dynamic>))
        .toList();
  }
}
