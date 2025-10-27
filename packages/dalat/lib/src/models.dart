import 'package:json_annotation/json_annotation.dart';
import 'converters.dart';

part 'models.g.dart';

@JsonSerializable()
class SysInfo {
  @JsonKey(name: 'hostname')
  String hostname;
  @JsonKey(name: 'os')
  String os;
  @JsonKey(name: 'platform')
  String platform;
  @JsonKey(name: 'memTotal')
  int memTotal;
  @JsonKey(name: 'memUsed')
  int memUsed;
  @JsonKey(name: 'diskTotal')
  int diskTotal;
  @JsonKey(name: 'diskUsed')
  int diskUsed;
  @JsonKey(name: 'batteryCharging')
  bool batteryCharging;
  @JsonKey(name: 'batteryPercent')
  double batteryPercent;
  @JsonKey(name: 'cpuUsage')
  double cpuUsage;
  SysInfo({
    required this.hostname,
    required this.os,
    required this.platform,
    required this.memTotal,
    required this.memUsed,
    required this.diskTotal,
    required this.diskUsed,
    required this.batteryCharging,
    required this.batteryPercent,
    required this.cpuUsage,
  });

  factory SysInfo.fromJson(Map<String, dynamic> json) =>
      _$SysInfoFromJson(json);
  Map<String, dynamic> toJson() => _$SysInfoToJson(this);
}

@JsonSerializable()
class DeviceColor {
  @JsonKey(name: 'name')
  final String name;
  @JsonKey(name: 'hex')
  final String hex;
  @JsonKey(name: 'r')
  final int r;
  @JsonKey(name: 'g')
  final int g;
  @JsonKey(name: 'b')
  final int b;

  DeviceColor({
    required this.name,
    required this.r,
    required this.g,
    required this.b,
    required this.hex,
  });
  factory DeviceColor.fromJson(Map<String, dynamic> json) =>
      _$DeviceColorFromJson(json);

  Map<String, dynamic> toJson() => _$DeviceColorToJson(this);
}

@JsonSerializable()
@Uint8ListConverter()
class AppConfig {
  @JsonKey(name: 'setupComplete')
  bool setupComplete;
  @JsonKey(name: 'deviceName')
  String deviceName;
  @JsonKey(name: 'deviceColor')
  DeviceColor deviceColor;
  @JsonKey(name: 'certificate')
  Certificate certificate;
  @JsonKey(name: 'deviceType')
  DeviceType deviceType;

  AppConfig({
    required this.setupComplete,
    required this.deviceName,
    required this.deviceColor,
    required this.certificate,
    required this.deviceType,
  });

  factory AppConfig.fromJson(Map<String, dynamic> json) =>
      _$AppConfigFromJson(json);
  Map<String, dynamic> toJson() => _$AppConfigToJson(this);
}

@JsonSerializable()
class SysInfoConfig {
  @JsonKey(name: 'enabled')
  bool enabled;
  @JsonKey(name: 'cacheSeconds')
  int cacheSeconds;

  SysInfoConfig({required this.enabled, required this.cacheSeconds});
  factory SysInfoConfig.fromJson(Map<String, dynamic> json) =>
      _$SysInfoConfigFromJson(json);
  Map<String, dynamic> toJson() => _$SysInfoConfigToJson(this);
}

@JsonSerializable()
class FileSendConfig {
  @JsonKey(name: 'enabled')
  bool enabled;
  @JsonKey(name: 'maxSize')
  int maxSize;
  @JsonKey(name: 'saveFolder')
  String saveFolder;

  FileSendConfig({
    required this.enabled,
    required this.maxSize,
    required this.saveFolder,
  });
  factory FileSendConfig.fromJson(Map<String, dynamic> json) =>
      _$FileSendConfigFromJson(json);
  Map<String, dynamic> toJson() => _$FileSendConfigToJson(this);
}

@JsonSerializable()
class ServiceConfig {
  @JsonKey(name: 'sysinfo')
  SysInfoConfig sysInfo;
  @JsonKey(name: 'filesend')
  FileSendConfig fileSend;

  ServiceConfig({required this.sysInfo, required this.fileSend});
  factory ServiceConfig.fromJson(Map<String, dynamic> json) =>
      _$ServiceConfigFromJson(json);
  Map<String, dynamic> toJson() => _$ServiceConfigToJson(this);
}

@JsonSerializable()
@Uint8ListConverter()
class FoundDevice {
  @JsonKey(name: 'ip')
  Ip ip;
  @JsonKey(name: 'port')
  Port port;
  @JsonKey(name: 'info')
  DeviceInfo info;

  FoundDevice({required this.ip, required this.port, required this.info});
  factory FoundDevice.fromJson(Map<String, dynamic> json) =>
      _$FoundDeviceFromJson(json);
  Map<String, dynamic> toJson() => _$FoundDeviceToJson(this);
}

@JsonSerializable()
@Uint8ListConverter()
class DeviceDetails {
  @JsonKey(name: 'color')
  DeviceColor color;
  @JsonKey(name: 'name')
  String name;
  @JsonKey(name: 'type')
  DeviceType type;
  @JsonKey(name: 'certificate')
  Certificate certificate;
  DeviceDetails({
    required this.color,
    required this.name,
    required this.type,
    required this.certificate,
  });
  factory DeviceDetails.fromJson(Map<String, dynamic> json) =>
      _$DeviceDetailsFromJson(json);
  Map<String, dynamic> toJson() => _$DeviceDetailsToJson(this);
}

@JsonSerializable()
@Uint8ListConverter()
class PairedDevice {
  @JsonKey(name: 'certificate')
  Certificate certificate;
  @JsonKey(name: 'token')
  PairToken token;
  PairedDevice({
    required this.certificate,
    required this.token,
  }); // Add fields that match the JSON output from Go
  factory PairedDevice.fromJson(Map<String, dynamic> json) =>
      _$PairedDeviceFromJson(json);
  Map<String, dynamic> toJson() => _$PairedDeviceToJson(this);
}

@JsonSerializable()
class DeviceInfo {
  @JsonKey(name: 'id')
  String id;
  @JsonKey(name: 'name')
  String name;
  @JsonKey(name: 'color')
  DeviceColor color;
  @JsonKey(name: 'type')
  DeviceType type;
  DeviceInfo({
    required this.id,
    required this.name,
    required this.color,
    required this.type,
  });
  factory DeviceInfo.fromJson(Map<String, dynamic> json) =>
      _$DeviceInfoFromJson(json);
  Map<String, dynamic> toJson() => _$DeviceInfoToJson(this);
}

@JsonSerializable()
@Uint8ListConverter()
class ConnectedDevice {
  @JsonKey(name: 'info')
  DeviceInfo info;
  @JsonKey(name: 'pairedDevice')
  PairedDevice pairedDevice;
  @JsonKey(name: 'ip')
  Ip ip;
  @JsonKey(name: 'port')
  Port port;
  ConnectedDevice({
    required this.info,
    required this.pairedDevice,
    required this.ip,
    required this.port,
  });
  factory ConnectedDevice.fromJson(Map<String, dynamic> json) =>
      _$ConnectedDeviceFromJson(json);
  Map<String, dynamic> toJson() => _$ConnectedDeviceToJson(this);
}

@JsonSerializable()
class NodeStatus {
  @JsonKey(name: 'discoveryRunning')
  bool discoveryRunning;
  @JsonKey(name: 'serverRunning')
  bool serverRunning;
  @JsonKey(name: 'workerRunning')
  bool workerRunning;
  @JsonKey(name: 'port')
  Port port;

  NodeStatus({
    required this.serverRunning,
    required this.workerRunning,
    required this.discoveryRunning,
    required this.port,
  }); // Add fields that match the JSON output from Go
  factory NodeStatus.fromJson(Map<String, dynamic> json) =>
      _$NodeStatusFromJson(json);
  Map<String, dynamic> toJson() => _$NodeStatusToJson(this);
}

@JsonSerializable()
class RequestPairResponse {
  @JsonKey(name: 'status')
  int status;
  @JsonKey(name: 'error')
  String error;
  @JsonKey(name: 'accepted')
  bool accepted;
  @JsonKey(name: 'reason')
  String reason;
  RequestPairResponse({
    required this.status,
    required this.error,
    required this.accepted,
    required this.reason,
  });
  factory RequestPairResponse.fromJson(Map<String, dynamic> json) =>
      _$RequestPairResponseFromJson(json);
  Map<String, dynamic> toJson() => _$RequestPairResponseToJson(this);
}

@JsonSerializable()
class SingleFileTransferStatus {
  @JsonKey(name: 'fileName')
  String fileName;
  @JsonKey(name: 'percent')
  double percent;
  @JsonKey(name: 'fileSize')
  int fileSize;
  @JsonKey(name: 'status')
  FileTransferStatus status;
  SingleFileTransferStatus({
    required this.fileName,
    required this.percent,
    required this.fileSize,
    required this.status,
  });
  factory SingleFileTransferStatus.fromJson(Map<String, dynamic> json) =>
      _$SingleFileTransferStatusFromJson(json);
  Map<String, dynamic> toJson() => _$SingleFileTransferStatusToJson(this);
}

@JsonSerializable()
class DeviceFileTransferStatus {
  @JsonKey(name: 'device')
  DeviceInfo device;
  @JsonKey(name: 'transfers')
  List<SingleFileTransferStatus>? transfers;
  @JsonKey(name: 'percent')
  double percent;
  DeviceFileTransferStatus({
    required this.device,
    required this.transfers,
    required this.percent,
  });
  factory DeviceFileTransferStatus.fromJson(Map<String, dynamic> json) =>
      _$DeviceFileTransferStatusFromJson(json);
  Map<String, dynamic> toJson() => _$DeviceFileTransferStatusToJson(this);
}

@JsonSerializable()
class FileTransfersStatus {
  @JsonKey(name: 'percentSending')
  double percentSending;
  @JsonKey(name: 'percentReceiving')
  double percentReceiving;
  @JsonKey(name: 'sending')
  List<DeviceFileTransferStatus>? sending;
  @JsonKey(name: 'receiving')
  List<DeviceFileTransferStatus>? receiving;
  FileTransfersStatus({
    required this.percentSending,
    required this.percentReceiving,
    required this.sending,
    required this.receiving,
  });

  factory FileTransfersStatus.fromJson(Map<String, dynamic> json) =>
      _$FileTransfersStatusFromJson(json);
  Map<String, dynamic> toJson() => _$FileTransfersStatusToJson(this);
}

@JsonSerializable()
class SharedFile {
  @JsonKey(name: 'uuid')
  String uuid;
  @JsonKey(name: 'path')
  String path;
  @JsonKey(name: 'name')
  String name;
  @JsonKey(name: 'size')
  int size;
  SharedFile({
    required this.uuid,
    required this.path,
    required this.name,
    required this.size,
  });
  factory SharedFile.fromJson(Map<String, dynamic> json) =>
      _$SharedFileFromJson(json);
  Map<String, dynamic> toJson() => _$SharedFileToJson(this);
}

@JsonSerializable()
class WebshareStatus {
  @JsonKey(name: 'running')
  bool running;
  @JsonKey(name: 'port')
  int port;
  @JsonKey(name: 'passcode')
  String passcode;
  @JsonKey(name: 'sharedFiles')
  List<SharedFile> sharedFiles;
  @JsonKey(name: 'shareURLs')
  List<String> shareURLs;
  WebshareStatus({
    required this.running,
    required this.port,
    required this.passcode,
    required this.sharedFiles,
    required this.shareURLs,
  });
  factory WebshareStatus.fromJson(Map<String, dynamic> json) =>
      _$WebshareStatusFromJson(json);
  Map<String, dynamic> toJson() => _$WebshareStatusToJson(this);
}
