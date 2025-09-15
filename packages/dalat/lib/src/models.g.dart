// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'models.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

DeviceColor _$DeviceColorFromJson(Map<String, dynamic> json) => DeviceColor(
  name: json['name'] as String,
  r: (json['r'] as num).toInt(),
  g: (json['g'] as num).toInt(),
  b: (json['b'] as num).toInt(),
  hex: json['hex'] as String,
);

Map<String, dynamic> _$DeviceColorToJson(DeviceColor instance) =>
    <String, dynamic>{
      'name': instance.name,
      'hex': instance.hex,
      'r': instance.r,
      'g': instance.g,
      'b': instance.b,
    };

AppSettings _$AppSettingsFromJson(Map<String, dynamic> json) => AppSettings(
  setupComplete: json['setupComplete'] as bool,
  deviceName: json['deviceName'] as String,
  deviceColor: DeviceColor.fromJson(
    json['deviceColor'] as Map<String, dynamic>,
  ),
);

Map<String, dynamic> _$AppSettingsToJson(AppSettings instance) =>
    <String, dynamic>{
      'setupComplete': instance.setupComplete,
      'deviceName': instance.deviceName,
      'deviceColor': instance.deviceColor,
    };

SysInfoSettings _$SysInfoSettingsFromJson(Map<String, dynamic> json) =>
    SysInfoSettings(
      enabled: json['enabled'] as bool,
      cacheSeconds: (json['cacheSeconds'] as num).toInt(),
    );

Map<String, dynamic> _$SysInfoSettingsToJson(SysInfoSettings instance) =>
    <String, dynamic>{
      'enabled': instance.enabled,
      'cacheSeconds': instance.cacheSeconds,
    };

FileSendSettings _$FileSendSettingsFromJson(Map<String, dynamic> json) =>
    FileSendSettings(
      enabled: json['enabled'] as bool,
      maxSize: (json['maxSize'] as num).toInt(),
      saveFolder: json['saveFolder'] as String,
    );

Map<String, dynamic> _$FileSendSettingsToJson(FileSendSettings instance) =>
    <String, dynamic>{
      'enabled': instance.enabled,
      'maxSize': instance.maxSize,
      'saveFolder': instance.saveFolder,
    };

ServiceSettings _$ServiceSettingsFromJson(
  Map<String, dynamic> json,
) => ServiceSettings(
  sysInfo: SysInfoSettings.fromJson(json['sysinfo'] as Map<String, dynamic>),
  fileSend: FileSendSettings.fromJson(json['filesend'] as Map<String, dynamic>),
);

Map<String, dynamic> _$ServiceSettingsToJson(ServiceSettings instance) =>
    <String, dynamic>{
      'sysinfo': instance.sysInfo,
      'filesend': instance.fileSend,
    };

FoundDevice _$FoundDeviceFromJson(Map<String, dynamic> json) => FoundDevice();

Map<String, dynamic> _$FoundDeviceToJson(FoundDevice instance) =>
    <String, dynamic>{};

PairedDevice _$PairedDeviceFromJson(Map<String, dynamic> json) =>
    PairedDevice();

Map<String, dynamic> _$PairedDeviceToJson(PairedDevice instance) =>
    <String, dynamic>{};

ConnectedDevice _$ConnectedDeviceFromJson(Map<String, dynamic> json) =>
    ConnectedDevice();

Map<String, dynamic> _$ConnectedDeviceToJson(ConnectedDevice instance) =>
    <String, dynamic>{};

NodeStatus _$NodeStatusFromJson(Map<String, dynamic> json) => NodeStatus();

Map<String, dynamic> _$NodeStatusToJson(NodeStatus instance) =>
    <String, dynamic>{};
