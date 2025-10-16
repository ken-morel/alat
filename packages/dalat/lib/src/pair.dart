import 'package:dalat/dalat.dart';
import 'package:json_annotation/json_annotation.dart';
part 'pair.g.dart';

@JsonSerializable()
class PairResponse {
  @JsonKey(name: 'accepted')
  bool accepted;
  @JsonKey(name: 'reason')
  String reason;
  PairResponse({required this.accepted, required this.reason});
  Map<String, dynamic> toJson() => _$PairResponseToJson(this);
  factory PairResponse.fromJson(Map<String, dynamic> json) =>
      _$PairResponseFromJson(json);
}

@JsonSerializable()
class PairRequest {
  @JsonKey(name: 'requestId')
  String requestid;
  @JsonKey(name: 'token')
  @Uint8ListConverter()
  PairToken token;
  @JsonKey(name: 'device')
  DeviceDetails device;
  PairRequest({
    required this.requestid,
    required this.token,
    required this.device,
  });
  Map<String, dynamic> toJson() => _$PairRequestToJson(this);
  factory PairRequest.fromJson(Map<String, dynamic> json) =>
      _$PairRequestFromJson(json);
}
