import 'dart:convert';
import 'dart:ffi';

import 'package:dalat/dalat.dart';
import 'package:dalat/src/bindings.dart';
import 'package:ffi/ffi.dart';

mixin InstanceConfig {
  Future<void> jsonSetterHelper(
    dynamic jsonData,
    int Function(int, Pointer<Char>) ffiFunc,
  );
  Future<T> jsonGetterHelper<T>(
    Pointer<Char> Function(int) ffiFunc,
    T Function(Map<String, dynamic>) fromJson,
  );

  static AppConfig createAppConfig() {
    final ptr = bindings.default_app_config();
    if (ptr == nullptr) {
      throw "Could not create default app settings, backend sent invalid null response";
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final Map<String, dynamic> decoded = jsonDecode(jsonStr);
      return AppConfig.fromJson(decoded);
    } finally {
      bindings.free_string(ptr);
    }
  }

  static ServiceConfig createServiceConfig() {
    final ptr = bindings.default_service_config();
    if (ptr == nullptr) {
      throw "Could not create default service configuration, backend sent invalid null response";
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final Map<String, dynamic> decoded = jsonDecode(jsonStr);
      return ServiceConfig.fromJson(decoded);
    } finally {
      bindings.free_string(ptr);
    }
  }

  int get handle;

  Future<AppConfig> getAppConfig() async {
    return jsonGetterHelper(bindings.get_app_config_json, AppConfig.fromJson);
  }

  Future<void> setAppConfig(AppConfig settings) async {
    return jsonSetterHelper(settings.toJson(), bindings.set_app_config_json);
  }

  Future<ServiceConfig> getServiceConfig() {
    return jsonGetterHelper(
      bindings.get_service_config_json,
      ServiceConfig.fromJson,
    );
  }

  Future<void> setServiceConfig(ServiceConfig settings) {
    return jsonSetterHelper(
      settings.toJson(),
      bindings.set_service_config_json,
    );
  }
}
