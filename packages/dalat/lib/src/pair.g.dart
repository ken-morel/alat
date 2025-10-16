// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'pair.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

PairResponse _$PairResponseFromJson(Map<String, dynamic> json) => PairResponse(
  accepted: json['accepted'] as bool,
  reason: json['reason'] as String,
);

Map<String, dynamic> _$PairResponseToJson(PairResponse instance) =>
    <String, dynamic>{'accepted': instance.accepted, 'reason': instance.reason};

PairRequest _$PairRequestFromJson(Map<String, dynamic> json) => PairRequest(
  requestid: json['requestId'] as String,
  token: const Uint8ListConverter().fromJson(json['token'] as List),
  device: DeviceDetails.fromJson(json['device'] as Map<String, dynamic>),
);

Map<String, dynamic> _$PairRequestToJson(PairRequest instance) =>
    <String, dynamic>{
      'requestId': instance.requestid,
      'token': const Uint8ListConverter().toJson(instance.token),
      'device': instance.device,
    };
