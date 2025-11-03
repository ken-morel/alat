import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ShareHandlerPage extends StatelessWidget {
  static const String routeName = '/share-handler';

  const ShareHandlerPage({super.key});

  @override
  Widget build(BuildContext context) {
    final appState = context.watch<AppState>();
    final sharedFiles = appState.sharedFiles.value;

    return Scaffold(
      appBar: AppBar(
        title: const Text('Share Files'),
      ),
      body: sharedFiles.isEmpty
          ? const Center(child: Text('No files to share.'))
          : Column(
              children: [
                Expanded(
                  child: ListView.builder(
                    itemCount: sharedFiles.length,
                    itemBuilder: (context, index) {
                      final filePath = sharedFiles[index];
                      final fileName = filePath.split('/').last;
                      return ListTile(
                        title: Text(fileName),
                        subtitle: Text(filePath),
                      );
                    },
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    children: [
                      ElevatedButton(
                        onPressed: () {
                          // TODO: Implement navigation to device selection for sending files
                          // For now, just clear shared files
                          appState.sharedFiles.value = [];
                          Navigator.pop(context);
                        },
                        child: const Text('Share to Device'),
                      ),
                      const SizedBox(height: 16),
                      ElevatedButton(
                        onPressed: () async {
                          await appState.webShareAddFiles();
                          appState.sharedFiles.value = [];
                          Navigator.pop(context);
                        },
                        child: const Text('Add to WebShare'),
                      ),
                    ],
                  ),
                ),
              ],
            ),
    );
  }
}
