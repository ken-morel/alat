import 'package:alat/theme/theme.dart';
import 'package:flutter/material.dart';

class ThemeProvider with ChangeNotifier {
  final ThemeData _themeData = AppTheme.darkTheme;

  ThemeData get themeData => _themeData;

  // In the future, you can add logic to switch to a light theme
  // void toggleTheme() {
  //   if (_themeData == AppTheme.darkTheme) {
  //     _themeData = AppTheme.lightTheme;
  //   } else {
  //     _themeData = AppTheme.darkTheme;
  //   }
  //   notifyListeners();
  // }
}
