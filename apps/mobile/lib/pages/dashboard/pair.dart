import 'package:alat/components/founddevices.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:dalat/dalat.dart' as dalat;

class PairDevicePage extends DashboardBase {
  const PairDevicePage({super.key});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(title: Text("Connect a device"));
  }

  @override
  Widget buildContent(BuildContext context) {
    return Column(
      children: [
        Center(
          child: Text(
            "Found devices",
            style: Theme.of(context).textTheme.headlineLarge,
          ),
        ),
        SizedBox(
          height: 300,
          child: FoundDevicesList(
            onConnectionUserRequest: (device) {
              Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) {
                    return Scaffold(
                      appBar: AppBar(leading: BackButton()),
                      body: RequestingPairPage(device: device),
                    );
                  },
                ),
              );
            },
          ),
        ),
      ],
    );
  }
}

class RequestingPairPage extends StatefulWidget {
  final dalat.FoundDevice device;
  const RequestingPairPage({super.key, required this.device});
  @override
  createState() => _RequestingPairPageState();
}

class _RequestingPairPageState extends State<RequestingPairPage> {
  late final Future<dalat.RequestPairResponse> pairRequestPromise;
  @override
  void initState() {
    final appState = context.read<AppState>();
    pairRequestPromise = appState.node!.requestPair(widget.device.info.id);
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: pairRequestPromise,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return _buildResponseView(context, snapshot.data!);
        } else if (snapshot.hasError) {
          return Column(
            children: [
              Icon(Icons.error, size: 100),
              Text("Error trying to connect to device: ${snapshot.error}"),
            ],
          );
        } else {
          return Column(
            children: [
              CircularProgressIndicator(),
              Text("Waiting for a response from the device"),
            ],
          );
        }
      },
    );
  }

  Widget _buildResponseView(
    BuildContext context,
    dalat.RequestPairResponse response,
  ) {
    if (response.status != 0) {
      return Column(
        children: [
          Icon(Icons.sms_failed),
          Text(
            "Could not connect to device",
            style: Theme.of(context).textTheme.headlineMedium,
          ),
          Text(response.error),
        ],
      );
    } else if (response.accepted) {
      return Column(
        children: [
          Icon(Icons.check_circle_sharp),
          Text(
            "Device connected succesfully",
            style: Theme.of(context).textTheme.headlineMedium,
          ),
        ],
      );
    } else {
      return Column(
        children: [
          Icon(Icons.question_mark),
          Text(
            "Could not connect to device",
            style: Theme.of(context).textTheme.headlineMedium,
          ),
          Text(response.reason),
        ],
      );
    }
  }
}
