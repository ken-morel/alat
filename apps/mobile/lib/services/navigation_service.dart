import 'package:flutter/material.dart';

/// A service to allow navigation from outside the widget tree,
/// for example, from a notification handler.
class NavigationService {
  // A global key that will be assigned to the MaterialApp's navigator.
  final GlobalKey<NavigatorState> navigatorKey = GlobalKey<NavigatorState>();

  /// Navigates to a named route.
  Future<dynamic> navigateTo(String routeName, {Object? arguments}) {
    return navigatorKey.currentState!.pushNamed(routeName, arguments: arguments);
  }

  /// Provides access to the current context if needed.
  BuildContext get currentContext => navigatorKey.currentContext!;
}
