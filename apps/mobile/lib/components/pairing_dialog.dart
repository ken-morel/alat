import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';

import '../state.dart';

class PairingDialog extends StatelessWidget {
  final PairRequestState pairRequestState;

  const PairingDialog({super.key, required this.pairRequestState});

  @override
  Widget build(BuildContext context) {
    final deviceDetails = pairRequestState.request.device;

    return AlertDialog(
      title: const Text('Pairing Request'),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text('A device wants to connect:'),
          const SizedBox(height: 16),
          ListTile(
            leading: Icon(
              deviceDetails.type == dalat.deviceTypeDesktop
                  ? Icons.desktop_windows
                  : Icons.phone_iphone,
              color: Color.fromRGBO(
                deviceDetails.color.r,
                deviceDetails.color.g,
                deviceDetails.color.b,
                1,
              ),
              size: 40,
            ),
            title: Text(
              deviceDetails.name,
              style: const TextStyle(fontWeight: FontWeight.bold),
            ),
            subtitle: Text('Type: ${deviceDetails.type}'),
          ),
          const SizedBox(height: 16),
          const Text(
            'Only accept if you recognize this device and initiated the connection.',
            style: TextStyle(fontSize: 12, fontStyle: FontStyle.italic),
          ),
        ],
      ),
      actions: <Widget>[
        TextButton(
          child: const Text('Deny'),
          onPressed: () {
            // Complete the future with a 'denied' response
            pairRequestState.completer.complete(
              dalat.PairResponse(accepted: false, reason: 'User denied'),
            );
            Navigator.of(context).pop(); // Close the dialog
          },
        ),
        FilledButton(
          child: const Text('Accept'),
          onPressed: () {
            // Complete the future with an 'accepted' response
            pairRequestState.completer.complete(
              dalat.PairResponse(accepted: true, reason: 'User accepted'),
            );
            Navigator.of(context).pop(); // Close the dialog
          },
        ),
      ],
    );
  }
}

