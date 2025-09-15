import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;

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
    r: 255,
    g: 0,
    b: 0,
  );
  @override
  void initState() {
    color = widget.setupState.appState.settings?.deviceColor ?? color;
    nameController.text =
        widget.setupState.appState.settings?.deviceName ?? "Name me";
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text("Device color", style: Theme.of(context).textTheme.headlineMedium),
        SizedBox(height: 10),
        Text("A color to more easily identify your device"),
        SizedBox(height: 20),
        _buildColorSelect(context),
        SizedBox(height: 20),
        _buildNameSelect(context),
      ],
    );
  }

  Widget _buildColorSelect(BuildContext context) {
    return FutureBuilder(
      future: widget.setupState.appState.alat!.getAlatColors(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          final colors = snapshot.data!;
          const cols = [0, 1, 2, 3, 4];
          const rows = [0, 1, 2, 3];
          return Column(
            children: rows
                .map(
                  (rowid) => Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: cols
                        .map(
                          (colid) => Container(
                            decoration: BoxDecoration(
                              borderRadius: BorderRadiusGeometry.circular(50),
                              color:
                                  color.name ==
                                      colors[colid + rowid * cols.length].name
                                  ? Theme.of(context).colorScheme.primary
                                  : null,
                            ),
                            child: Padding(
                              padding: EdgeInsetsGeometry.all(5),
                              child: FilledButton(
                                onPressed: () {
                                  setState(() {
                                    color = colors[colid + rowid * cols.length];
                                  });
                                  _save();
                                },
                                style: FilledButton.styleFrom(
                                  backgroundColor: Color.fromRGBO(
                                    colors[colid + rowid * cols.length].r,
                                    colors[colid + rowid * cols.length].g,
                                    colors[colid + rowid * cols.length].b,
                                    1,
                                  ),
                                ),
                                child: const SizedBox(width: 0, height: 55),
                              ),
                            ),
                          ),
                        )
                        .toList(),
                  ),
                )
                .toList(),
          );
        } else {
          return CircularProgressIndicator();
        }
      },
    );
  }

  Widget _buildNameSelect(BuildContext context) {
    return Column(
      children: [
        Text("Device name", style: Theme.of(context).textTheme.headlineMedium),
        SizedBox(height: 20),
        Text("Because each of your devices also merits a name"),
        SizedBox(height: 30),
        TextField(
          onChanged: (_) => _save(),
          controller: nameController,
          textCapitalization: TextCapitalization.words,
          maxLines: 1,
          maxLength: 63,
          decoration: InputDecoration(
            label: Text("Name"),
            helperText: "A short, memorable name",
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
