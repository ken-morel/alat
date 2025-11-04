import 'dart:isolate';
import 'dart:async';

import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class AlatStatusWidget extends StatefulWidget {
  final dalat.AlatInstance node;
  const AlatStatusWidget({super.key, required this.node});
  @override
  State<StatefulWidget> createState() => _AlatStatusWidgetState();
}

class _AlatStatusWidgetState extends State<AlatStatusWidget> {
  dalat.NodeStatus? status;
  String? error;
  Timer? timer;
  late AppState appState;

  @override
  void initState() {
    appState = context.read();
    timer = Timer.periodic(
      Duration(milliseconds: 200),
      (_) async => _fetchStatus(),
    );
    _fetchStatus();
    super.initState();
  }

  void _fetchStatus() async {
    dalat.NodeStatus? stat;
    String? err;
    try {
      stat = await appState.node?.getNodeStatus();
    } catch (e) {
      err = e.toString();
      stat = null;
    } finally {
      setState(() {
        status = stat;
        error = err;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    if (error != null) {
      return Column(
        children: [
          Icon(Icons.error, color: Colors.redAccent),
          Text(error!),
        ],
      );
    } else if (status == null) {
      return const CircularProgressIndicator();
    } else {
      return _statusView(context, status!);
    }
  }

  Widget _statusView(BuildContext context, dalat.NodeStatus status) {
    final node = widget.node;
    final theme = Theme.of(context);
    final okay =
        status.discoveryRunning && status.serverRunning && status.workerRunning;
    final bool anyRunning =
        status.discoveryRunning || status.serverRunning || status.workerRunning;
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
          SizedBox(height: 10),
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
          FilledButton.tonal(
            onPressed: () {
              Isolate.run(() {
                if (anyRunning) {
                  node.stop();
                } else {
                  node.start();
                }
              });
            },
            child: Text(anyRunning ? "Stop" : "Start"),
          ),
        ],
      ),
    );
  }

  @override
  void dispose() {
    timer?.cancel();
    super.dispose();
  }
}
