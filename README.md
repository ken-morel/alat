# Alat Project

A cross-platform application for connecting devices.

## Project Structure

- `proto/`: Protocol Buffer definitions. This is the schema for all communication.
- `pkg/core/`: The core Go daemon logic, shared across platforms.
- `apps/`: Contains the platform-specific applications.
  - `desktop/`: Wails app for Windows, macOS, and Linux.
  - `mobile/`: Flutter app for iOS and Android.
  - `server/`: Go server for remote discovery and relay.
- `go.work`: Go workspace file for managing the Go modules.
- `manage.fish`: Main project management script.

## Getting Started

### 1. Install Protocol Buffers Compiler

You need the `protoc` compiler to generate code from the `.proto` files.

- **macOS:** `brew install protobuf`
- **Linux:** `sudo apt-get install -y protobuf-compiler`
- **Windows:** `choco install protoc` or download from the [official releases page](https://github.com/protocolbuffers/protobuf/releases).

### 2. Generate Protocol Code

Use the `manage.fish` script to generate the necessary Go and Dart code from the protocol definitions.

```fish
./manage.fish proto
```

This will:
1. Install the necessary Go and Dart plugins.
2. Generate Go code in `pkg/core/proto/`.
3. Generate Dart code in `apps/mobile/lib/src/proto/`.

You can see all available commands by running:
```fish
./manage.fish help
```

Now you are ready to start developing!
