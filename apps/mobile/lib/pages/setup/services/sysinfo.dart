import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'dart:async'; // Required for Timer

class SysInfoSetupPage extends StatefulWidget {
  final SetupState setupState;
  const SysInfoSetupPage({super.key, required this.setupState});
  @override
  State<SysInfoSetupPage> createState() => _SysInfoSetupPageState();
}

class _SysInfoSetupPageState extends State<SysInfoSetupPage> {
  late bool enabled;
  late TextEditingController _cacheSecondsController;
  Timer? _debounce;

  @override
  void initState() {
    super.initState();
    enabled =
        widget.setupState.appState.serviceSettings?.sysInfo.enabled ?? true;
    _cacheSecondsController = TextEditingController(
      text:
          (widget.setupState.appState.serviceSettings?.sysInfo.cacheSeconds ??
                  10)
              .toString(),
    );
  }

  @override
  void dispose() {
    _cacheSecondsController.dispose();
    _debounce?.cancel();
    super.dispose();
  }

  void _onCacheSecondsChanged(String value) {
    if (_debounce?.isActive ?? false) _debounce!.cancel();
    _debounce = Timer(const Duration(milliseconds: 500), () {
      final int? newCacheSeconds = int.tryParse(value);
      if (newCacheSeconds != null && newCacheSeconds >= 0) {
        // Here you would update your setupState
        // widget.setupState.appState.serviceSettings?.sysInfo.cacheSeconds = newCacheSeconds;
        print("Cache seconds updated to: $newCacheSeconds");
      } else {
        // Optionally show an error or reset to a valid value
        if (newCacheSeconds != null && newCacheSeconds < 0) {
          print("Cache seconds cannot be negative.");
        }
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "System Information and Stats",
            style: Theme.of(context).textTheme.headlineMedium,
          ),
          const SizedBox(height: 10),
          Text(
            "This permits other devices to display this device's battery, memory, and other system information.",
            style: Theme.of(context).textTheme.bodyLarge,
          ),
          const SizedBox(height: 20),
          _ServiceSwitch(
            value: enabled,
            onChanged: (value) {
              setState(() {
                enabled = value;
                // Update setupState here
                // widget.setupState.appState.serviceSettings?.sysInfo.enabled = value;
              });
            },
          ),
          if (enabled) ...[
            const SizedBox(height: 30),
            Text("Options", style: Theme.of(context).textTheme.headlineSmall),
            const SizedBox(height: 20),
            _CacheSecondsInput(
              controller: _cacheSecondsController,
              onChanged: _onCacheSecondsChanged,
            ),
          ],
        ],
      ),
    );
  }
}

class _ServiceSwitch extends StatelessWidget {
  final bool value;
  final ValueChanged<bool> onChanged;

  const _ServiceSwitch({required this.value, required this.onChanged});

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Checkbox(
          value: value,
          onChanged: (newValue) {
            if (newValue != null) {
              onChanged(newValue);
            }
          },
        ),
        Text(value ? "Service Enabled" : "Service Disabled"),
      ],
    );
  }
}

class _CacheSecondsInput extends StatelessWidget {
  final TextEditingController controller;
  final ValueChanged<String> onChanged;

  const _CacheSecondsInput({required this.controller, required this.onChanged});

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      keyboardType: TextInputType.number,
      decoration: const InputDecoration(
        labelText: "Cache Refresh Interval",
        hintText: "e.g., 10",
        suffixText: "seconds",
        border: OutlineInputBorder(),
      ),
      inputFormatters: [FilteringTextInputFormatter.digitsOnly],
      onChanged: onChanged,
    );
  }
}
