import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class WebShareWidget extends StatefulWidget {
  const WebShareWidget({super.key});

  @override
  State<WebShareWidget> createState() => _WebShareWidgetState();
}

class _WebShareWidgetState extends State<WebShareWidget> {
  final _passcodeController = TextEditingController();

  @override
  void dispose() {
    _passcodeController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final appState = context.watch<AppState>();
    final status = appState.webShareStatus;
    final isRunning = status?.running ?? false;

    return ListView(
      padding: const EdgeInsets.all(16.0),
      children: [
        Card(
          child: ClipRRect(
            borderRadius: BorderRadius.circular(
              12.0,
            ), // Match Card's default border radius
            child: ExpansionTile(
              title: const Text('Server Controls'),
              initiallyExpanded: true,
              children: [
                ListTile(
                  title: const Text('Status'),
                  trailing: Text(
                    isRunning ? 'Running' : 'Stopped',
                    style: TextStyle(
                      color: isRunning ? Colors.green : Colors.red,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                if (isRunning && status != null) ...[
                  ListTile(
                    title: const Text('Share URL'),
                    subtitle: Text(status.shareURLs.first),
                  ),
                  ListTile(
                    title: const Text('Passcode'),
                    subtitle: Text(status.passcode),
                  ),
                ],
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: ElevatedButton(
                    onPressed: () {
                      if (isRunning) {
                        appState.webShareStop();
                      } else {
                        appState.webShareStart();
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: isRunning ? Colors.red : Colors.green,
                    ),
                    child: Text(isRunning ? 'Stop Server' : 'Start Server'),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 16.0,
                    vertical: 16.0,
                  ),
                  child: TextField(
                    controller: _passcodeController,
                    decoration: InputDecoration(
                      labelText: 'Set New Passcode',
                      suffixIcon: IconButton(
                        icon: const Icon(Icons.send),
                        onPressed: () {
                          if (_passcodeController.text.isNotEmpty) {
                            appState.webShareSetPasscode(
                              _passcodeController.text,
                            );
                            _passcodeController.clear();
                            FocusScope.of(context).unfocus();
                          }
                        },
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
        const SizedBox(height: 16),
        Card(
          child: ClipRRect(
            borderRadius: BorderRadius.circular(
              12.0,
            ), // Match Card's default border radius
            child: ExpansionTile(
              title: const Text('Shared Files'),
              initiallyExpanded: true,
              children: [
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Row(
                    children: [
                      Expanded(
                        child: ElevatedButton.icon(
                          onPressed: () => appState.webShareAddFiles(),
                          icon: const Icon(Icons.add),
                          label: const Text('Add Files'),
                        ),
                      ),
                      const SizedBox(width: 16),
                      Expanded(
                        child: ElevatedButton.icon(
                          onPressed: status?.sharedFiles.isNotEmpty ?? false
                              ? () => appState.webShareClearFiles()
                              : null,
                          icon: const Icon(Icons.delete_sweep),
                          label: const Text('Clear All'),
                          style: ElevatedButton.styleFrom(
                            backgroundColor: Colors.orange,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                if (status?.sharedFiles.isEmpty ?? true)
                  const ListTile(
                    title: Center(child: Text('No files are shared.')),
                  )
                else
                  for (final file in status!.sharedFiles)
                    ListTile(
                      title: Text(file.name),
                      subtitle: Text(
                        '${(file.size / 1024).toStringAsFixed(2)} KB',
                      ),
                      trailing: IconButton(
                        icon: const Icon(Icons.close, color: Colors.red),
                        onPressed: () => appState.webShareRemoveFile(file.uuid),
                      ),
                    ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
