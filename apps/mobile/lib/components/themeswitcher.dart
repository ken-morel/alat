import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:alat/theme.dart';

class ThemeSwitcher extends StatelessWidget {
  const ThemeSwitcher({super.key});

  @override
  Widget build(BuildContext context) {
    final themeProvider = Provider.of<ThemeProvider>(context);

    return RadioGroup(
      onChanged: (ThemeMode? value) {
        if (value != null) {
          themeProvider.setThemeMode(value);
        }
      },
      groupValue: themeProvider.themeMode,
      child: Column(
        children: [
          RadioListTile<ThemeMode>(
            title: const Text('Light'),
            value: ThemeMode.light,
          ),
          RadioListTile<ThemeMode>(
            title: const Text('Dark'),
            value: ThemeMode.dark,
          ),
          RadioListTile<ThemeMode>(
            title: const Text('System'),
            value: ThemeMode.system,
          ),
        ],
      ),
    );
  }
}
