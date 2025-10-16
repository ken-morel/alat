import 'dart:async';

import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

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
  double? targetPercent;
  String? error;
  late Timer timer;
  late Timer updateTimer;
  late AppState _appState;
  bool charging = false;
  void _fetchData() async {
    try {
      final info = _appState.node!.queryConnectedDeviceSysInfo(
        widget.connectedDevice.info.id,
      );
      setState(() {
        error = null;
        charging = info.batteryCharging;
        targetPercent = info.batteryPercent;
      });
    } catch (e) {
      setState(() {
        error = e.toString();
        percent = null;
        targetPercent = null;
      });
    }
  }

  @override
  void initState() {
    _appState = context.read<AppState>();
    _fetchData();
    timer = Timer.periodic(Duration(seconds: 2), (_) => _fetchData());
    updateTimer = Timer.periodic(Duration(milliseconds: 100), (_) {
      if (targetPercent == null) return;
      percent ??= 0;
      setState(() {
        if (percent! >= targetPercent!) {
          percent = targetPercent;
        } else {
          final diff = (targetPercent! - percent!);
          percent = percent! + (diff > 10 ? diff : 10) / 5;
        }
      });
    });
    super.initState();
  }

  @override
  void dispose() {
    timer.cancel();
    updateTimer.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
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
                color: Color.fromRGBO(
                  widget.connectedDevice.info.color.r,
                  widget.connectedDevice.info.color.g,
                  widget.connectedDevice.info.color.b,
                  1,
                ),
              ),
              CircularProgressIndicator(
                value: (percent ?? 100) / 100,
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
