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
  certificate: const Uint8ListConverter().fromJson(json['certificate'] as List),
);

Map<String, dynamic> _$AppSettingsToJson(AppSettings instance) =>
    <String, dynamic>{
      'setupComplete': instance.setupComplete,
      'deviceName': instance.deviceName,
      'deviceColor': instance.deviceColor,
      'certificate': const Uint8ListConverter().toJson(instance.certificate),
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

FoundDevice _$FoundDeviceFromJson(Map<String, dynamic> json) => FoundDevice(
  ip: const Uint8ListConverter().fromJson(json['ip'] as List),
  port: (json['port'] as num).toInt(),
  info: DeviceInfo.fromJson(json['info'] as Map<String, dynamic>),
);

Map<String, dynamic> _$FoundDeviceToJson(FoundDevice instance) =>
    <String, dynamic>{
      'ip': const Uint8ListConverter().toJson(instance.ip),
      'port': instance.port,
      'info': instance.info,
    };

DeviceDetails _$DeviceDetailsFromJson(Map<String, dynamic> json) =>
    DeviceDetails(
      color: DeviceColor.fromJson(json['color'] as Map<String, dynamic>),
      name: json['name'] as String,
      type: $enumDecode(_$DeviceTypeEnumMap, json['type']),
      certificate: const Uint8ListConverter().fromJson(
        json['certificate'] as List,
      ),
    );

Map<String, dynamic> _$DeviceDetailsToJson(DeviceDetails instance) =>
    <String, dynamic>{
      'color': instance.color,
      'name': instance.name,
      'type': _$DeviceTypeEnumMap[instance.type]!,
      'certificate': const Uint8ListConverter().toJson(instance.certificate),
    };

const _$DeviceTypeEnumMap = {
  DeviceType.unspecified: 'unspecified',
  DeviceType.mobile: 'mobile',
  DeviceType.desktop: 'desktop',
  DeviceType.tv: 'tv',
};

PairedDevice _$PairedDeviceFromJson(Map<String, dynamic> json) => PairedDevice(
  certificate: const Uint8ListConverter().fromJson(json['certificate'] as List),
  token: const Uint8ListConverter().fromJson(json['token'] as List),
);

Map<String, dynamic> _$PairedDeviceToJson(PairedDevice instance) =>
    <String, dynamic>{
      'certificate': const Uint8ListConverter().toJson(instance.certificate),
      'token': const Uint8ListConverter().toJson(instance.token),
    };

DeviceInfo _$DeviceInfoFromJson(Map<String, dynamic> json) => DeviceInfo(
  id: json['id'] as String,
  name: json['name'] as String,
  color: DeviceColor.fromJson(json['color'] as Map<String, dynamic>),
  type: $enumDecode(_$DeviceTypeEnumMap, json['type']),
);

Map<String, dynamic> _$DeviceInfoToJson(DeviceInfo instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'color': instance.color,
      'type': _$DeviceTypeEnumMap[instance.type]!,
    };

ConnectedDevice _$ConnectedDeviceFromJson(Map<String, dynamic> json) =>
    ConnectedDevice(
      info: DeviceInfo.fromJson(json['info'] as Map<String, dynamic>),
      pairedDevice: PairedDevice.fromJson(
        json['pairedDevice'] as Map<String, dynamic>,
      ),
      ip: const Uint8ListConverter().fromJson(json['ip'] as List),
      port: (json['port'] as num).toInt(),
    );

Map<String, dynamic> _$ConnectedDeviceToJson(ConnectedDevice instance) =>
    <String, dynamic>{
      'info': instance.info,
      'pairedDevice': instance.pairedDevice,
      'ip': const Uint8ListConverter().toJson(instance.ip),
      'port': instance.port,
    };

NodeStatus _$NodeStatusFromJson(Map<String, dynamic> json) => NodeStatus(
  serverRunning: json['serverRunning'] as bool,
  workerRunning: json['workerRunning'] as bool,
  discoveryRunning: json['discoveryRunning'] as bool,
  port: (json['port'] as num).toInt(),
);

Map<String, dynamic> _$NodeStatusToJson(NodeStatus instance) =>
    <String, dynamic>{
      'discoveryRunning': instance.discoveryRunning,
      'serverRunning': instance.serverRunning,
      'workerRunning': instance.workerRunning,
      'port': instance.port,
    };
