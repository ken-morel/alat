import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

import '../../state.dart';

class ConnectingPage extends StatefulWidget {
  final dalat.FoundDevice device;

  const ConnectingPage({super.key, required this.device});

  @override
  State<ConnectingPage> createState() => _ConnectingPageState();
}

class _ConnectingPageState extends State<ConnectingPage> {
  late Future<dalat.RequestPairResponse> _pairRequestPromise;

  @override
  void initState() {
    super.initState();
    _initiatePairing();
  }

  void _initiatePairing() {
    final appState = context.read<AppState>();
    setState(() {
      _pairRequestPromise = Future.delayed(
        Duration(seconds: 2),
        () => appState.node!.requestPair(widget.device.info.id),
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Pairing'),
        // The back button will be hidden while connecting
        automaticallyImplyLeading: false,
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: FutureBuilder(
            future: _pairRequestPromise,
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return _buildConnectingView();
              } else if (snapshot.hasError) {
                return _buildFailedView(
                  'Connection Failed',
                  error: snapshot.error.toString(),
                );
              } else if (snapshot.hasData) {
                final response = snapshot.data!;
                if (response.status != 0) {
                  return _buildFailedView(
                    'Connection Failed',
                    error: response.error,
                  );
                } else if (response.accepted) {
                  return _buildSuccessView();
                } else {
                  return _buildFailedView(
                    'Connection Refused',
                    error: response.reason,
                  );
                }
              }
              // Should not happen
              return _buildFailedView('An unknown error occurred.');
            },
          ),
        ),
      ),
    );
  }

  Widget _buildConnectingView() {
    return Column(
      key: const ValueKey('connecting'),
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const CircularProgressIndicator(),
        const SizedBox(height: 32),
        Text(
          'Requesting to pair with...',
          style: Theme.of(context).textTheme.titleMedium,
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 8),
        Text(
          widget.device.info.name,
          style: Theme.of(context).textTheme.headlineSmall,
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 16),
        Text(
          'Please accept the request on the other device.',
          style: Theme.of(context).textTheme.bodyMedium,
          textAlign: TextAlign.center,
        ),
      ],
    );
  }

  Widget _buildSuccessView() {
    return Column(
      key: const ValueKey('success'),
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const Icon(Icons.check_circle, color: Colors.green, size: 80),
        const SizedBox(height: 24),
        Text(
          'Successfully Paired!',
          style: Theme.of(context).textTheme.headlineSmall,
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 8),
        Text(
          'You can now share services with ${widget.device.info.name}.',
          textAlign: TextAlign.center,
        ),
        const SizedBox(height: 32),
        ElevatedButton(
          onPressed: () {
            // Pop twice: once for this page, once for the pair page.
            Navigator.of(context).pop();
          },
          child: const Text('Done'),
        ),
      ],
    );
  }

  Widget _buildFailedView(String message, {String? error}) {
    return Column(
      key: const ValueKey('failed'),
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Icon(Icons.error, color: Theme.of(context).colorScheme.error, size: 80),
        const SizedBox(height: 24),
        Text(
          message,
          style: Theme.of(context).textTheme.headlineSmall,
          textAlign: TextAlign.center,
        ),
        if (error != null && error.isNotEmpty) ...[
          const SizedBox(height: 8),
          Text(
            error,
            textAlign: TextAlign.center,
            style: Theme.of(context).textTheme.bodySmall,
          ),
        ],
        const SizedBox(height: 32),
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            OutlinedButton(
              onPressed: () => Navigator.of(context).pop(),
              child: const Text('Cancel'),
            ),
            const SizedBox(width: 12),
            ElevatedButton(
              onPressed: _initiatePairing,
              child: const Text('Try Again'),
            ),
          ],
        ),
      ],
    );
  }
}
