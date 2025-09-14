import 'package:flutter/material.dart';

// Generated from the desktop application's theme.css
// OKLCH values have been converted to standard RGB hex values.

const _lightColorScheme = ColorScheme(
  brightness: Brightness.light,
  primary: Color(0xFF4F6385), // --color-primary-500
  onPrimary: Color(0xFFFFFFFF),
  secondary: Color(0xFF5A7D6C), // --color-secondary-500
  onSecondary: Color(0xFFFFFFFF),
  error: Color(0xFFD95F32), // --color-error-500
  onError: Color(0xFFFFFFFF),
  background: Color(0xFFFFFFFF), // --body-background-color
  onBackground: Color(0xFF1A1C1E),
  surface: Color(0xFFFDFBFF),
  onSurface: Color(0xFF1A1C1E),
);

const _darkColorScheme = ColorScheme(
  brightness: Brightness.dark,
  primary: Color(0xFF4F6385), // --color-primary-500
  onPrimary: Color(0xFFFFFFFF),
  secondary: Color(0xFF5A7D6C), // --color-secondary-500
  onSecondary: Color(0xFFFFFFFF),
  error: Color(0xFFD95F32), // --color-error-500
  onError: Color(0xFFFFFFFF),
  background: Color(0xFF2A3757), // --body-background-color-dark
  onBackground: Color(0xFFE2E2E6),
  surface: Color(0xFF1A1C1E),
  onSurface: Color(0xFFE2E2E6),
);

final ThemeData lightTheme = ThemeData(
  useMaterial3: true,
  colorScheme: _lightColorScheme,
  appBarTheme: const AppBarTheme(
    centerTitle: true,
    backgroundColor: Colors.transparent,
    elevation: 0,
  ),
);

final ThemeData darkTheme = ThemeData(
  useMaterial3: true,
  colorScheme: _darkColorScheme,
  appBarTheme: const AppBarTheme(
    centerTitle: true,
    backgroundColor: Colors.transparent,
    elevation: 0,
  ),
);
