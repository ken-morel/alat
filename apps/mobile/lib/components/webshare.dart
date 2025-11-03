import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
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
    final theme = Theme.of(context);

    return Scaffold(
      body: ListView(
        padding: const EdgeInsets.all(8.0),
        children: [
          Card(
            elevation: 2.0,
            child: Column(
              children: [
                SwitchListTile(
                  title: const Text('WebShare Server'),
                  value: isRunning,
                  onChanged: (bool value) {
                    if (value) {
                      appState.webShareStart();
                    } else {
                      appState.webShareStop();
                    }
                  },
                  secondary: Icon(
                    isRunning ? Icons.public : Icons.public_off,
                    color: isRunning
                        ? theme.colorScheme.primary
                        : theme.colorScheme.onSurface.withOpacity(0.6),
                  ),
                ),
                if (isRunning && status != null) ...[
                  const Divider(),
                  ListTile(
                    leading: const Icon(Icons.link),
                    title: const Text('Share URL'),
                    subtitle: Text(status.shareURLs.first),
                    onTap: () {
                      Clipboard.setData(
                        ClipboardData(text: status.shareURLs.first),
                      );
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(
                          content: Text('URL Copied to Clipboard'),
                        ),
                      );
                    },
                  ),
                  ListTile(
                    leading: const Icon(Icons.password),
                    title: const Text('Passcode'),
                    subtitle: Text(
                      status.passcode,
                      style: const TextStyle(fontWeight: FontWeight.bold),
                    ),
                    onTap: () {
                      Clipboard.setData(ClipboardData(text: status.passcode));
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(
                          content: Text('Passcode Copied to Clipboard'),
                        ),
                      );
                    },
                  ),
                  Padding(
                    padding: const EdgeInsets.fromLTRB(16, 0, 16, 16),
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
              ],
            ),
          ),
          const SizedBox(height: 16),
          Card(
            elevation: 2.0,
            child: Column(
              children: [
                ListTile(
                  title: const Text('Shared Files'),
                  trailing: (status?.sharedFiles.isNotEmpty ?? false)
                      ? IconButton(
                          icon: Icon(
                            Icons.delete_sweep,
                            color: theme.colorScheme.error,
                          ),
                          onPressed: () => appState.webShareClearFiles(),
                          tooltip: 'Clear All Files',
                        )
                      : null,
                ),
                const Divider(),
                if (status?.sharedFiles.isEmpty ?? true)
                  const ListTile(
                    title: Center(
                      child: Padding(
                        padding: EdgeInsets.symmetric(vertical: 24.0),
                        child: Text('No files are shared.'),
                      ),
                    ),
                  )
                else
                  ListView.builder(
                    shrinkWrap: true,
                    physics: const NeverScrollableScrollPhysics(),
                    itemCount: status!.sharedFiles.length,
                    itemBuilder: (context, index) {
                      final file = status.sharedFiles[index];
                      return ListTile(
                        title: Text(file.name),
                        subtitle: Text(
                          '${(file.size / 1024).toStringAsFixed(2)} KB',
                        ),
                        trailing: IconButton(
                          icon: Icon(
                            Icons.close,
                            color: theme.colorScheme.error,
                          ),
                          onPressed: () =>
                              appState.webShareRemoveFile(file.uuid),
                        ),
                      );
                    },
                  ),
              ],
            ),
          ),
        ],
      ),
      floatingActionButton: isRunning
          ? FloatingActionButton.extended(
              onPressed: () => appState.webShareAddFiles(),
              icon: const Icon(Icons.add),
              label: const Text('Add Files'),
            )
          : null,
    );
  }
}
