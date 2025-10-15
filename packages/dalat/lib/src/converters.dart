import 'dart:typed_data';

import 'package:json_annotation/json_annotation.dart';

typedef DeviceType = String;

const DeviceType deviceTypeMobile = "mobile",
    deviceTypeDesktop = "desktop",
    deviceTypeTV = "tv",
    deviceTypeWeb = "web",
    deviceTypeArduino = "arduino",
    deviceTypeUnspecified = "unspecified";

typedef Certificate = Uint8List;
typedef Ip = String;
typedef Port = int;

class Uint8ListConverter implements JsonConverter<Uint8List, List<dynamic>> {
  const Uint8ListConverter();

  @override
  Uint8List fromJson(json) {
    return Uint8List.fromList(json.cast<int>());
  }

  @override
  toJson(Uint8List object) {
    return object.toList();
  }
}
