import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;

class DeviceIcon extends StatelessWidget {
  final dalat.DeviceType deviceType;
  final dalat.DeviceColor color;
  final double size;

  const DeviceIcon({
    super.key,
    required this.deviceType,
    required this.color,
    this.size = 20,
  });

  @override
  Widget build(BuildContext context) {
    return Icon(
      DeviceIcon.iconFor(deviceType),
      size: size,
      color: Color.fromRGBO(color.r, color.g, color.b, 1),
    );
  }

  static IconData iconFor(dalat.DeviceType deviceType) {
    switch (deviceType) {
      case dalat.deviceTypeMobile:
        return Icons.phone_rounded;
      case dalat.deviceTypeDesktop:
        return Icons.computer_rounded;
      case dalat.deviceTypeTV:
        return Icons.tv_rounded;
      case dalat.deviceTypeWeb:
        return Icons.web_rounded;
      case dalat.deviceTypeArduino:
        return Icons.device_hub_rounded;
      default:
        return Icons.devices_rounded;
    }
  }
}
