import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;

class AlatStatusWidget extends StatelessWidget {
  final dalat.AlatInstance node;
  const AlatStatusWidget({super.key, required this.node});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<dalat.NodeStatus>(
      future: node.getNodeStatus(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return _statusView(context, snapshot.data!);
        } else if (snapshot.hasError) {
          return const Icon(Icons.error, color: Colors.redAccent);
        } else {
          return const CircularProgressIndicator();
        }
      },
    );
  }

  Widget _statusView(BuildContext context, dalat.NodeStatus status) {
    final theme = Theme.of(context);
    final okay =
        status.discoveryRunning && status.serverRunning && status.workerRunning;
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text("Port", style: Theme.of(context).textTheme.headlineSmall),
              Text("Status", style: Theme.of(context).textTheme.headlineSmall),
            ],
          ),
          Divider(),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text(
                '${status.port}',
                style: theme.textTheme.bodyLarge?.copyWith(
                  fontWeight: FontWeight.bold,
                ),
              ),
              Row(
                children: [
                  const Text("W"),
                  const SizedBox(width: 2),
                  CircleAvatar(
                    backgroundColor: status.workerRunning
                        ? Colors.greenAccent.shade700
                        : Colors.redAccent,
                    radius: 12,
                    child: Icon(
                      okay ? Icons.check : Icons.close,
                      color: Colors.white,
                      size: 16,
                    ),
                  ),
                  const SizedBox(width: 5),
                  const Text("D"),
                  const SizedBox(width: 2),
                  CircleAvatar(
                    backgroundColor: status.discoveryRunning
                        ? Colors.greenAccent.shade700
                        : Colors.redAccent,
                    radius: 12,
                    child: Icon(
                      okay ? Icons.check : Icons.close,
                      color: Colors.white,
                      size: 16,
                    ),
                  ),
                  const SizedBox(width: 5),
                  const Text("S"),
                  const SizedBox(width: 2),
                  CircleAvatar(
                    backgroundColor: status.serverRunning
                        ? Colors.greenAccent.shade700
                        : Colors.redAccent,
                    radius: 12,
                    child: Icon(
                      okay ? Icons.check : Icons.close,
                      color: Colors.white,
                      size: 16,
                    ),
                  ),
                ],
              ),
            ],
          ),
        ],
      ),
    );
  }
}
