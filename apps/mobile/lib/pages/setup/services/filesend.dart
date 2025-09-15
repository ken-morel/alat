import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:flutter/services.dart';

class FileSendSetupPage extends StatefulWidget {
  final SetupState setupState;
  const FileSendSetupPage({super.key, required this.setupState});
  @override
  State<FileSendSetupPage> createState() => _FileSendSetupPageState();
}

class _FileSendSetupPageState extends State<FileSendSetupPage> {
  late bool enabled;
  late int maxSize;
  late final String saveFolder;
  final TextEditingController maxSizeController = TextEditingController();
  @override
  void initState() {
    enabled =
        widget.setupState.appState.serviceSettings?.fileSend.enabled ?? true;
    maxSize =
        ((widget.setupState.appState.serviceSettings?.fileSend.maxSize ?? 0) /
                (1024 * 1024))
            .round();
    maxSizeController.text = maxSize.toString();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(
          AppLocalizations.of(context)!.fileReceive,
          style: Theme.of(context).textTheme.headlineMedium,
        ),
        Text(
          AppLocalizations.of(
            context,
          )!.permitConnectedDevicesToSendFilesToThisDevicesWithouthPriorConfirmation,
        ),
        SizedBox(height: 20),
        Row(
          children: [
            Checkbox(
              value: enabled,
              onChanged: (newValue) {
                if (newValue != null) {
                  setState(() {
                    enabled = newValue;
                  });
                  _save();
                }
              },
            ),
            Text(
              enabled
                  ? AppLocalizations.of(context)!.serviceEnabled
                  : AppLocalizations.of(context)!.serviceDisabled,
            ),
          ],
        ),
        SizedBox(height: 20),
        Text(
          AppLocalizations.of(context)!.options,
          style: Theme.of(context).textTheme.headlineSmall,
        ),
        SizedBox(height: 10),
        _MaxSizeInput(
          controller: maxSizeController,
          onChanged: (String val) {
            maxSize = int.parse(val);
            _save();
          },
        ),
      ],
    );
  }

  Future<void> _save() async {
    widget.setupState.appState.serviceSettings?.fileSend.enabled = enabled;
    widget.setupState.appState.serviceSettings?.fileSend.maxSize =
        maxSize * 1024 * 1024;
    await widget.setupState.appState.saveSettings();
  }
}

class _MaxSizeInput extends StatelessWidget {
  final TextEditingController controller;
  final ValueChanged<String> onChanged;

  const _MaxSizeInput({required this.controller, required this.onChanged});

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      keyboardType: TextInputType.number,
      decoration: InputDecoration(
        labelText: AppLocalizations.of(context)!.maximulFileSize,
        hintText: "e.g., 10",
        suffixText: AppLocalizations.of(context)!.megaBytes,
        border: OutlineInputBorder(),
      ),
      inputFormatters: [FilteringTextInputFormatter.digitsOnly],
      onChanged: onChanged,
    );
  }
}
