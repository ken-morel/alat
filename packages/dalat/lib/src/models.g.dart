// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'models.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

SysInfo _$SysInfoFromJson(Map<String, dynamic> json) => SysInfo(
  hostname: json['hostname'] as String,
  os: json['os'] as String,
  platform: json['platform'] as String,
  memTotal: (json['memTotal'] as num).toInt(),
  memUsed: (json['memUsed'] as num).toInt(),
  diskTotal: (json['diskTotal'] as num).toInt(),
  diskUsed: (json['diskUsed'] as num).toInt(),
  batteryCharging: json['batteryCharging'] as bool,
  batteryPercent: (json['batteryPercent'] as num).toDouble(),
  cpuUsage: (json['cpuUsage'] as num).toDouble(),
);

Map<String, dynamic> _$SysInfoToJson(SysInfo instance) => <String, dynamic>{
  'hostname': instance.hostname,
  'os': instance.os,
  'platform': instance.platform,
  'memTotal': instance.memTotal,
  'memUsed': instance.memUsed,
  'diskTotal': instance.diskTotal,
  'diskUsed': instance.diskUsed,
  'batteryCharging': instance.batteryCharging,
  'batteryPercent': instance.batteryPercent,
  'cpuUsage': instance.cpuUsage,
};

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

AppConfig _$AppConfigFromJson(Map<String, dynamic> json) => AppConfig(
  setupComplete: json['setupComplete'] as bool,
  deviceName: json['deviceName'] as String,
  deviceColor: DeviceColor.fromJson(
    json['deviceColor'] as Map<String, dynamic>,
  ),
  certificate: const Uint8ListConverter().fromJson(json['certificate'] as List),
  deviceType: json['deviceType'] as String,
);

Map<String, dynamic> _$AppConfigToJson(AppConfig instance) => <String, dynamic>{
  'setupComplete': instance.setupComplete,
  'deviceName': instance.deviceName,
  'deviceColor': instance.deviceColor,
  'certificate': const Uint8ListConverter().toJson(instance.certificate),
  'deviceType': instance.deviceType,
};

SysInfoConfig _$SysInfoConfigFromJson(Map<String, dynamic> json) =>
    SysInfoConfig(
      enabled: json['enabled'] as bool,
      cacheSeconds: (json['cacheSeconds'] as num).toInt(),
    );

Map<String, dynamic> _$SysInfoConfigToJson(SysInfoConfig instance) =>
    <String, dynamic>{
      'enabled': instance.enabled,
      'cacheSeconds': instance.cacheSeconds,
    };

FileSendConfig _$FileSendConfigFromJson(Map<String, dynamic> json) =>
    FileSendConfig(
      enabled: json['enabled'] as bool,
      maxSize: (json['maxSize'] as num).toInt(),
      saveFolder: json['saveFolder'] as String,
    );

Map<String, dynamic> _$FileSendConfigToJson(FileSendConfig instance) =>
    <String, dynamic>{
      'enabled': instance.enabled,
      'maxSize': instance.maxSize,
      'saveFolder': instance.saveFolder,
    };

ServiceConfig _$ServiceConfigFromJson(Map<String, dynamic> json) =>
    ServiceConfig(
      sysInfo: SysInfoConfig.fromJson(json['sysinfo'] as Map<String, dynamic>),
      fileSend: FileSendConfig.fromJson(
        json['filesend'] as Map<String, dynamic>,
      ),
    );

Map<String, dynamic> _$ServiceConfigToJson(ServiceConfig instance) =>
    <String, dynamic>{
      'sysinfo': instance.sysInfo,
      'filesend': instance.fileSend,
    };

FoundDevice _$FoundDeviceFromJson(Map<String, dynamic> json) => FoundDevice(
  ip: json['ip'] as String,
  port: (json['port'] as num).toInt(),
  info: DeviceInfo.fromJson(json['info'] as Map<String, dynamic>),
);

Map<String, dynamic> _$FoundDeviceToJson(FoundDevice instance) =>
    <String, dynamic>{
      'ip': instance.ip,
      'port': instance.port,
      'info': instance.info,
    };

DeviceDetails _$DeviceDetailsFromJson(Map<String, dynamic> json) =>
    DeviceDetails(
      color: DeviceColor.fromJson(json['color'] as Map<String, dynamic>),
      name: json['name'] as String,
      type: json['type'] as String,
      certificate: const Uint8ListConverter().fromJson(
        json['certificate'] as List,
      ),
    );

Map<String, dynamic> _$DeviceDetailsToJson(DeviceDetails instance) =>
    <String, dynamic>{
      'color': instance.color,
      'name': instance.name,
      'type': instance.type,
      'certificate': const Uint8ListConverter().toJson(instance.certificate),
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
  type: json['type'] as String,
);

Map<String, dynamic> _$DeviceInfoToJson(DeviceInfo instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'color': instance.color,
      'type': instance.type,
    };

ConnectedDevice _$ConnectedDeviceFromJson(Map<String, dynamic> json) =>
    ConnectedDevice(
      info: DeviceInfo.fromJson(json['info'] as Map<String, dynamic>),
      pairedDevice: PairedDevice.fromJson(
        json['pairedDevice'] as Map<String, dynamic>,
      ),
      ip: json['ip'] as String,
      port: (json['port'] as num).toInt(),
    );

Map<String, dynamic> _$ConnectedDeviceToJson(ConnectedDevice instance) =>
    <String, dynamic>{
      'info': instance.info,
      'pairedDevice': instance.pairedDevice,
      'ip': instance.ip,
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

RequestPairResponse _$RequestPairResponseFromJson(Map<String, dynamic> json) =>
    RequestPairResponse(
      status: (json['status'] as num).toInt(),
      error: json['error'] as String,
      accepted: json['accepted'] as bool,
      reason: json['reason'] as String,
    );

Map<String, dynamic> _$RequestPairResponseToJson(
  RequestPairResponse instance,
) => <String, dynamic>{
  'status': instance.status,
  'error': instance.error,
  'accepted': instance.accepted,
  'reason': instance.reason,
};

SingleFileTransferStatus _$SingleFileTransferStatusFromJson(
  Map<String, dynamic> json,
) => SingleFileTransferStatus(
  fileName: json['fileName'] as String,
  percent: (json['percent'] as num).toDouble(),
  fileSize: (json['fileSize'] as num).toInt(),
  status: json['status'] as String,
);

Map<String, dynamic> _$SingleFileTransferStatusToJson(
  SingleFileTransferStatus instance,
) => <String, dynamic>{
  'fileName': instance.fileName,
  'percent': instance.percent,
  'fileSize': instance.fileSize,
  'status': instance.status,
};

DeviceFileTransferStatus _$DeviceFileTransferStatusFromJson(
  Map<String, dynamic> json,
) => DeviceFileTransferStatus(
  device: DeviceInfo.fromJson(json['device'] as Map<String, dynamic>),
  transfers: (json['transfers'] as List<dynamic>?)
      ?.map((e) => SingleFileTransferStatus.fromJson(e as Map<String, dynamic>))
      .toList(),
  percent: (json['percent'] as num).toDouble(),
);

Map<String, dynamic> _$DeviceFileTransferStatusToJson(
  DeviceFileTransferStatus instance,
) => <String, dynamic>{
  'device': instance.device,
  'transfers': instance.transfers,
  'percent': instance.percent,
};

FileTransfersStatus _$FileTransfersStatusFromJson(
  Map<String, dynamic> json,
) => FileTransfersStatus(
  percentSending: (json['percentSending'] as num).toDouble(),
  percentReceiving: (json['percentReceiving'] as num).toDouble(),
  sending: (json['sending'] as List<dynamic>?)
      ?.map((e) => DeviceFileTransferStatus.fromJson(e as Map<String, dynamic>))
      .toList(),
  receiving: (json['receiving'] as List<dynamic>?)
      ?.map((e) => DeviceFileTransferStatus.fromJson(e as Map<String, dynamic>))
      .toList(),
);

Map<String, dynamic> _$FileTransfersStatusToJson(
  FileTransfersStatus instance,
) => <String, dynamic>{
  'percentSending': instance.percentSending,
  'percentReceiving': instance.percentReceiving,
  'sending': instance.sending,
  'receiving': instance.receiving,
};
