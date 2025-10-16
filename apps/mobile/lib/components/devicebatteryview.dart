import 'dart:async';

import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;

class DeviceBatteryView extends StatefulWidget {
  final dalat.ConnectedDevice connectedDevice;
  final double size;
  const DeviceBatteryView({
    super.key,
    required this.connectedDevice,
    this.size = 100,
  });

  @override
  State<DeviceBatteryView> createState() => _DeviceBatteryViewState();
}

class _DeviceBatteryViewState extends State<DeviceBatteryView> {
  double? percent;
  String? error;
  late Timer timer;
  bool charging = false;
  @override
  Widget build(BuildContext context) {
    final textTheme = Theme.of(context).textTheme;
    return SizedBox(
      width: widget.size,
      height: widget.size,
      child: Container(
        decoration: BoxDecoration(shape: BoxShape.circle, border: Border()),
        child: Center(
          child: Stack(
            alignment: AlignmentGeometry.xy(0, 0),
            children: [
              Icon(
                error != null
                    ? Icons.battery_alert_rounded
                    : charging
                    ? Icons.battery_charging_full
                    : percent == 100
                    ? Icons.battery_full_rounded
                    : Icons.battery_0_bar_rounded,
                size: widget.size / 3,
                color: textTheme.bodySmall?.color,
              ),
              CircularProgressIndicator(
                value: 0.5,
                color: Theme.of(context).colorScheme.primary,
                backgroundColor: Theme.of(context).colorScheme.primaryContainer,
                valueColor: Animation.fromValueListenable(
                  ValueNotifier(
                    error != null
                        ? Colors.red
                        : percent != null
                        ? percent == 100
                              ? Colors.green
                              : Colors.blue
                        : Colors.yellow,
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
