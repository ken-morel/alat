import 'package:json_annotation/json_annotation.dart';

part 'models.g.dart';

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
class AppSettings {
  @JsonKey(name: 'setupComplete')
  bool setupComplete;
  @JsonKey(name: 'deviceName')
  String deviceName;
  @JsonKey(name: 'deviceColor')
  DeviceColor deviceColor;

  AppSettings({
    required this.setupComplete,
    required this.deviceName,
    required this.deviceColor,
  });

  factory AppSettings.fromJson(Map<String, dynamic> json) =>
      _$AppSettingsFromJson(json);
  Map<String, dynamic> toJson() => _$AppSettingsToJson(this);
}

@JsonSerializable()
class SysInfoSettings {
  @JsonKey(name: 'enabled')
  bool enabled;
  @JsonKey(name: 'cacheSeconds')
  int cacheSeconds;

  SysInfoSettings({required this.enabled, required this.cacheSeconds});
  factory SysInfoSettings.fromJson(Map<String, dynamic> json) =>
      _$SysInfoSettingsFromJson(json);
  Map<String, dynamic> toJson() => _$SysInfoSettingsToJson(this);
}

@JsonSerializable()
class FileSendSettings {
  @JsonKey(name: 'enabled')
  bool enabled;
  @JsonKey(name: 'maxSize')
  int maxSize;
  @JsonKey(name: 'saveFolder')
  String saveFolder;

  FileSendSettings({
    required this.enabled,
    required this.maxSize,
    required this.saveFolder,
  });
  factory FileSendSettings.fromJson(Map<String, dynamic> json) =>
      _$FileSendSettingsFromJson(json);
  Map<String, dynamic> toJson() => _$FileSendSettingsToJson(this);
}

@JsonSerializable()
class ServiceSettings {
  @JsonKey(name: 'sysinfo')
  SysInfoSettings sysInfo;
  @JsonKey(name: 'filesend')
  FileSendSettings fileSend;

  ServiceSettings({required this.sysInfo, required this.fileSend});
  factory ServiceSettings.fromJson(Map<String, dynamic> json) =>
      _$ServiceSettingsFromJson(json);
  Map<String, dynamic> toJson() => _$ServiceSettingsToJson(this);
}

@JsonSerializable()
class FoundDevice {
  FoundDevice(); // Add fields that match the JSON output from Go
  factory FoundDevice.fromJson(Map<String, dynamic> json) =>
      _$FoundDeviceFromJson(json);
  Map<String, dynamic> toJson() => _$FoundDeviceToJson(this);
}

@JsonSerializable()
class PairedDevice {
  PairedDevice(); // Add fields that match the JSON output from Go
  factory PairedDevice.fromJson(Map<String, dynamic> json) =>
      _$PairedDeviceFromJson(json);
  Map<String, dynamic> toJson() => _$PairedDeviceToJson(this);
}

@JsonSerializable()
class ConnectedDevice {
  ConnectedDevice(); // Add fields that match the JSON output from Go
  factory ConnectedDevice.fromJson(Map<String, dynamic> json) =>
      _$ConnectedDeviceFromJson(json);
  Map<String, dynamic> toJson() => _$ConnectedDeviceToJson(this);
}

@JsonSerializable()
class NodeStatus {
  NodeStatus(); // Add fields that match the JSON output from Go
  factory NodeStatus.fromJson(Map<String, dynamic> json) =>
      _$NodeStatusFromJson(json);
  Map<String, dynamic> toJson() => _$NodeStatusToJson(this);
}

enum DeviceType { unspecified, mobile, desktop, tv }
