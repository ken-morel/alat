import 'dart:typed_data';

import 'package:json_annotation/json_annotation.dart';

typedef Certificate = Uint8List;
typedef Ip = Uint8List;
typedef Port = int;

class Uint8ListConverter implements JsonConverter<Uint8List, List<dynamic>> {
  const Uint8ListConverter();

  @override
  Uint8List fromJson(List<dynamic> json) {
    return Uint8List.fromList(json.cast<int>());
  }

  @override
  List<dynamic> toJson(Uint8List object) {
    return object.toList();
  }
}
