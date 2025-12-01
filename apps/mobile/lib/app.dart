import 'package:alat/components/pairing_dialog.dart';
import 'package:alat/pages/dashboard/devicefilesend.dart';
import 'package:alat/pages/dashboard/pair.dart';
import 'package:alat/pages/dashboard/webshare.dart';
import 'package:alat/pages/share_handler_page.dart';
import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:alat/pages/dashboard/dashboard.dart';
import 'package:alat/pages/setup/setup.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:alat/pages/start.dart';
import 'package:alat/theme.dart';
import 'package:provider/provider.dart';

class AlatApplication extends StatelessWidget {
  final GlobalKey<NavigatorState> navigatorKey;
  const AlatApplication({super.key, required this.navigatorKey});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => ThemeProvider(),
      child: Consumer<ThemeProvider>(
        builder: (context, themeProvider, child) {
          return MaterialApp(
            title: 'Alat',
            theme: createLightTheme(),
            darkTheme: createDarkTheme(),
            localizationsDelegates: AppLocalizations.localizationsDelegates,
            supportedLocales: AppLocalizations.supportedLocales,
            themeMode: themeProvider.themeMode,
            navigatorKey: navigatorKey, // Assign the global key
            home: Consumer<AppState>(
              builder: (context, appState, _) {
                // This builder will now also handle showing the pairing dialog.
                return ValueListenableBuilder<PairRequestState?>(
                  valueListenable: appState.pendingPairRequest,
                  builder: (context, pairRequest, child) {
                    // If there's a pending request and the dialog isn't already shown, show it.
                    if (pairRequest != null) {
                      // Use a post-frame callback to show the dialog after the build is complete.
                      WidgetsBinding.instance.addPostFrameCallback((_) {
                        showDialog(
                          context: context,
                          barrierDismissible: false, // User must respond
                          builder: (BuildContext dialogContext) {
                            return PairingDialog(pairRequestState: pairRequest);
                          },
                        );
                      });
                    }

                    // Return the main page content based on the app state.
                    if (!appState.isReady) {
                      return const StartPage(); // Or a loading indicator
                    }
                    if (appState.settings?.setupComplete ?? false) {
                      return const DashboardPage();
                    } else {
                      return SetupAssistantPageView();
                    }
                  },
                );
              },
            ),
            // Use onGenerateRoute to handle showing the dialog from a notification.
            onGenerateRoute: (settings) {
              // Handle other routes normally.
              switch (settings.name) {
                case '/dashboard':
                  return MaterialPageRoute(
                    builder: (_) => const DashboardPage(),
                  );
                case '/dashboard/pair':
                  return MaterialPageRoute(
                    builder: (_) => const PairDevicePage(),
                  );
                case '/setup':
                  return MaterialPageRoute(
                    builder: (_) => SetupAssistantPageView(),
                  );
                case '/sendfiles':
                  return MaterialPageRoute(
                    builder: (_) => const SendFilesPage(),
                  );
                case '/webshare':
                  return MaterialPageRoute(
                    builder: (_) => const WebSharePage(),
                  );
                case ShareHandlerPage.routeName:
                  return MaterialPageRoute(
                    builder: (_) => const ShareHandlerPage(),
                  );
                default:
                  return MaterialPageRoute(builder: (_) => const StartPage());
              }
            },
          );
        },
      ),
    );
  }
}
