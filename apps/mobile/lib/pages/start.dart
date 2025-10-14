import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:alat/state.dart';
import 'package:alat/l10n/app_localizations.dart';

class StartPage extends StatelessWidget {
  const StartPage({super.key});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: FutureBuilder(
          future: context.read<AppState>().initialize(),
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              WidgetsBinding.instance.addPostFrameCallback((_) {
                if (snapshot.data!) {
                  Navigator.of(context).pushReplacementNamed("/dashboard");
                } else {
                  Navigator.of(context).pushReplacementNamed("/setup");
                }
              });
              return Container();
            } else if (snapshot.connectionState == ConnectionState.waiting) {
              return Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  CircularProgressIndicator(),
                  SizedBox(height: 30),
                  Text(
                    AppLocalizations.of(context)!.initializingAlatCore,
                    style: Theme.of(context).textTheme.headlineSmall,
                  ),
                ],
              );
            } else {
              return Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Icon(Icons.error, size: 200, color: Colors.redAccent),
                  SizedBox(height: 30),
                  Padding(
                    padding: EdgeInsetsGeometry.symmetric(
                      horizontal: 50,
                      vertical: 10,
                    ),
                    child: Text(
                      AppLocalizations.of(
                        context,
                      )!.errorInitializingAlatCoreError(
                        snapshot.hasError ? snapshot.error.toString() : "...",
                      ),
                      style: Theme.of(context).textTheme.headlineSmall,
                    ),
                  ),
                ],
              );
            }
          },
        ),
      ),
    );
  }
}
