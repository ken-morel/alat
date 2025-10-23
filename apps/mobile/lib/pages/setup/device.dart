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
  late Future<List<dalat.DeviceColor>> _colorsFuture;
  dalat.DeviceColor color = dalat.DeviceColor(
    name: "blue",
    hex: "#0000FF",
    r: 0,
    g: 0,
    b: 255,
  );
  @override
  void initState() {
    super.initState();
    _colorsFuture = widget.setupState.appState.node!.getAlatColors();
    color = widget.setupState.appState.settings?.deviceColor ?? color;
    nameController.text =
        widget.setupState.appState.settings?.deviceName ??
        AppLocalizations.of(context)!.nameMe;
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
    return FutureBuilder<List<dalat.DeviceColor>>(
      future: _colorsFuture,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Center(child: CircularProgressIndicator());
        }
        if (snapshot.hasError) {
          return Center(child: Text('Error: ${snapshot.error}'));
        }
        if (!snapshot.hasData || snapshot.data!.isEmpty) {
          return const Center(child: Text('No colors available.'));
        }

        final colors = snapshot.data!;
        return GridView.builder(
          shrinkWrap: true,
          physics: const NeverScrollableScrollPhysics(),
          gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
            crossAxisCount: 5,
            crossAxisSpacing: 10,
            mainAxisSpacing: 10,
          ),
          itemCount: colors.length,
          itemBuilder: (context, index) {
            final tileColor = colors[index];
            final isSelected = color.name == tileColor.name;
            final tileColorRgb =
                Color.fromRGBO(tileColor.r, tileColor.g, tileColor.b, 1);
            final iconColor =
                ThemeData.estimateBrightnessForColor(tileColorRgb) ==
                        Brightness.dark
                    ? Colors.white
                    : Colors.black;

            return GestureDetector(
              onTap: () {
                setState(() {
                  color = tileColor;
                });
                _save();
              },
              child: Container(
                width: 65,
                height: 65,
                decoration: BoxDecoration(
                  color: tileColorRgb,
                  shape: BoxShape.circle,
                  border: isSelected
                      ? Border.all(
                          color: Theme.of(context).colorScheme.primary,
                          width: 4,
                        )
                      : null,
                ),
                child: isSelected
                    ? Icon(
                        Icons.check,
                        color: iconColor,
                      )
                    : null,
              ),
            );
          },
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
