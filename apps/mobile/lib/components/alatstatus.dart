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
              Text('Alat Status', style: theme.textTheme.titleLarge),
              CircleAvatar(
                backgroundColor: okay
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
          const SizedBox(height: 16),
          _buildStatusRow(context, 'Discovery', status.discoveryRunning),
          const Divider(),
          _buildStatusRow(context, 'Server', status.serverRunning),
          const Divider(),
          _buildStatusRow(context, 'Worker', status.workerRunning),
          const Divider(),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text('Port', style: theme.textTheme.bodyMedium),
              Text(
                '${status.port}',
                style: theme.textTheme.bodyLarge?.copyWith(
                  fontWeight: FontWeight.bold,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildStatusRow(BuildContext context, String title, bool isRunning) {
    final theme = Theme.of(context);
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(title, style: theme.textTheme.bodyMedium),
          Row(
            children: [
              Icon(
                isRunning ? Icons.check_circle : Icons.cancel,
                color: isRunning ? Colors.green : Colors.red,
                size: 20,
              ),
              const SizedBox(width: 8),
              Text(
                isRunning ? 'Running' : 'Stopped',
                style: theme.textTheme.bodyMedium?.copyWith(
                  color: isRunning ? Colors.green : Colors.red,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
