import 'package:flutter/material.dart';

class AppTheme {
  static final ThemeData darkTheme = ThemeData(
    brightness: Brightness.dark,
    primaryColor: const Color(0xFF5a756e), // $sage
    scaffoldBackgroundColor: const Color(0xFF121212), // $background
    cardColor: const Color(0xFF1a1a1a), // $black
    dividerColor: const Color(0xFF2c3e2c), // $dark-green

    colorScheme: const ColorScheme.dark(
      primary: Color(0xFF5a756e), // $sage
      secondary: Color(0xFF456b45), // $secondary
      surface: Color(0xFF1a1a1a), // $black
      background: Color(0xFF121212), // $background
      error: Colors.redAccent,
      onPrimary: Colors.white,
      onSecondary: Colors.white,
      onSurface: Color(0xFFe0e0e0), // $text-primary
      onBackground: Color(0xFFe0e0e0),
      onError: Colors.white,
    ),

    appBarTheme: const AppBarTheme(
      backgroundColor: Color(0xFF1a1a1a), // $black
      elevation: 0,
      titleTextStyle: TextStyle(
        color: Color(0xFFe0e0e0),
        fontSize: 20,
        fontWeight: FontWeight.w500,
      ),
      iconTheme: IconThemeData(color: Color(0xFFe0e0e0)),
    ),

    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: Color(0xFFe0e0e0)),
      bodyMedium: TextStyle(color: Color(0xFFb0bec5)),
      displayLarge: TextStyle(color: Color(0xFFe0e0e0)),
      headlineMedium: TextStyle(color: Color(0xFFe0e0e0)),
    ),

    buttonTheme: ButtonThemeData(
      buttonColor: const Color(0xFF5a756e), // $sage
      textTheme: ButtonTextTheme.primary,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(8.0),
      ),
    ),

    floatingActionButtonTheme: const FloatingActionButtonThemeData(
      backgroundColor: Color(0xFF5a756e), // $sage
    ),

    // A placeholder for a light theme if needed in the future
    // static final ThemeData lightTheme = ThemeData.light();
  );
}
