import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:alat/l10n/app_localizations.dart';

class SetupDevice extends StatefulWidget {
  final SetupState setupState;
  const SetupDevice({super.key, required this.setupState});
  @override
  State<StatefulWidget> createState() => _SetupDeviceState();
}

class _SetupDeviceState extends State<SetupDevice> {
  final TextEditingController nameController = TextEditingController();
  dalat.DeviceColor color = dalat.DeviceColor(
    name: "blue",
    hex: "#0000FF",
    r: 0,
    g: 0,
    b: 255,
  );
  @override
  void initState() {
    color = widget.setupState.appState.settings?.deviceColor ?? color;
    nameController.text =
        widget.setupState.appState.settings?.deviceName ??
        AppLocalizations.of(context)!.nameMe;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(
          AppLocalizations.of(context)!.deviceColor,
          style: Theme.of(context).textTheme.headlineMedium,
        ),
        SizedBox(height: 10),
        Text(
          AppLocalizations.of(context)!.aColorToMoreEasilyIdentifyYourDevice,
        ),
        SizedBox(height: 20),
        _buildColorSelect(context),
        SizedBox(height: 20),
        _buildNameSelect(context),
      ],
    );
  }

  Widget _buildColorSelect(BuildContext context) {
    return FutureBuilder(
      future: Future.delayed(
        Duration(milliseconds: 500),
        () => widget.setupState.appState.node!.getAlatColors(),
      ),
      builder: (context, snapshot) {
        List<dalat.DeviceColor>? colors;
        if (snapshot.hasData) {
          colors = snapshot.data!;
        }
        const cols = [0, 1, 2, 3, 4];
        const rows = [0, 1, 2, 3];
        return Column(
          children: rows
              .map(
                (rowid) => Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: cols.map((colid) {
                    dalat.DeviceColor tileColor = colors == null
                        ? color
                        : colors[colid + rowid * cols.length];
                    return Container(
                      decoration: BoxDecoration(
                        borderRadius: BorderRadiusGeometry.circular(50),
                        color: color.name == tileColor.name
                            ? Theme.of(context).colorScheme.primary
                            : null,
                      ),
                      child: Padding(
                        padding: EdgeInsetsGeometry.all(5),
                        child: FilledButton(
                          onPressed: () {
                            setState(() {
                              color = tileColor;
                            });
                            _save();
                          },
                          style: FilledButton.styleFrom(
                            backgroundColor: Color.fromRGBO(
                              tileColor.r,
                              tileColor.g,
                              tileColor.b,
                              1,
                            ),
                          ),
                          child: const SizedBox(width: 0, height: 55),
                        ),
                      ),
                    );
                  }).toList(),
                ),
              )
              .toList(),
        );
      },
    );
  }

  Widget _buildNameSelect(BuildContext context) {
    return Column(
      children: [
        Text(
          AppLocalizations.of(context)!.deviceName,
          style: Theme.of(context).textTheme.headlineMedium,
        ),
        SizedBox(height: 20),
        Text(
          AppLocalizations.of(context)!.becauseEachOfYourDevicesAlsoMeritsAName,
        ),
        SizedBox(height: 30),
        TextField(
          onChanged: (_) => _save(),
          controller: nameController,
          textCapitalization: TextCapitalization.words,
          maxLines: 1,
          maxLength: 63,
          decoration: InputDecoration(
            label: Text(AppLocalizations.of(context)!.name),
            helperText: AppLocalizations.of(context)!.aShortMemorableName,
            filled: true,
          ),
        ),
      ],
    );
  }

  void _save() {
    widget.setupState.appState.settings?.deviceName = nameController.text;
    widget.setupState.appState.settings?.deviceColor = color;
    widget.setupState.appState.saveSettings();
  }
}
