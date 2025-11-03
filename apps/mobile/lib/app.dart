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
                if (!appState.isReady) {
                  return const StartPage(); // Or a loading indicator
                }
                if (appState.sharedFiles.value.isNotEmpty) {
                  return const ShareHandlerPage();
                } else if (appState.settings?.setupComplete ?? false) {
                  return const DashboardPage();
                } else {
                  return SetupAssistantPageView();
                }
              },
            ),
            // Use onGenerateRoute to handle showing the dialog from a notification.
            onGenerateRoute: (settings) {
              if (settings.name == '/pair-request') {
                return MaterialPageRoute(
                  builder: (context) {
                    // This page will be the one that shows the dialog.
                    return const PairingRequestHandlerPage();
                  },
                  settings: settings,
                );
              }
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

/// A helper page that listens for the pairing request and shows the dialog.
class PairingRequestHandlerPage extends StatelessWidget {
  const PairingRequestHandlerPage({super.key});

  @override
  Widget build(BuildContext context) {
    final appState = Provider.of<AppState>(context, listen: false);

    return ValueListenableBuilder<PairRequestState?>(
      valueListenable: appState.pendingPairRequest,
      builder: (context, value, child) {
        if (value == null) {
          // If the request is cleared, pop the page.
          WidgetsBinding.instance.addPostFrameCallback((_) {
            if (Navigator.canPop(context)) {
              Navigator.pop(context);
            }
          });
          return const Scaffold(
            body: Center(child: Text('No active pairing request.')),
          );
        }

        return PairingDialog(pairRequestState: value);
      },
    );
  }
}
