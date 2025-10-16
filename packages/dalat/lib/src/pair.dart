import 'package:dalat/dalat.dart';

class PairResponse {
  bool accepted;
  String reason;
  PairResponse({required this.accepted, required this.reason});
}

class PairRequest {
  String requestid;
  PairToken token;
  DeviceDetails device;
  PairRequest({
    required this.requestid,
    required this.token,
    required this.device,
  });
}
