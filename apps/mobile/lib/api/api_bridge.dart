import 'dart:convert';
import 'dart:ffi';
import 'dart:io';
import 'package:ffi/ffi.dart';

// Define the signature of the Go function
typedef SearchDevicesFunc = Pointer<Utf8> Function();
// Define the Dart type for the function
typedef SearchDevices = Pointer<Utf8> Function();

class ApiBridge {
  late final DynamicLibrary _lib;
  late final Function searchDevices;

  ApiBridge() {
    _lib = Platform.isAndroid
        ? DynamicLibrary.open('libalat.so')
        : DynamicLibrary.open('build/libalat.so'); // For Linux testing

    final SearchDevices search = _lib
        .lookup<NativeFunction<SearchDevicesFunc>>('SearchDevices')
        .asFunction();

    searchDevices = () {
      final resultPtr = search();
      final result = resultPtr.toDartString();
      // Free the string allocated in Go
      // Note: Go's CString allocates memory that needs to be freed.
      // We will need to add a function to the Go bridge to free this memory.
      // For now, we will ignore this to get the functionality working.
      return jsonDecode(result);
    };
  }
}
