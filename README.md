# Alat

**Alat: Your devices, unified.**

Alat is a cross-platform application that enables seamless communication and service sharing between your devices. It allows you to create a unified workspace where you can easily share files, receive notifications, and control your devices from a single place.

## Features

*   **Cross-Platform**: Alat is available for desktop (Windows, macOS, Linux) and mobile (Android, iOS) devices.
*   **P2P Communication**: Devices communicate directly with each other using a secure peer-to-peer connection.
*   **Service Sharing**: Share files, clipboard content, and more between your devices.
*   **Extensible**: The modular architecture allows for the addition of new services and features.

## Applications

### Mobile Application

The mobile application is built with Flutter and provides a user-friendly interface for managing your devices and services on the go. It interacts with the Alat core through the `dalat` Dart FFI plugin.

### Desktop Application

The desktop application is built with Wails, using Go for the backend and SvelteKit for the frontend. It provides a rich user experience with access to all of Alat's features.

## Alat Core

The core of Alat is a Go module that contains the P2P runtime, services, and business logic. It is completely platform-agnostic and can be embedded in any application.

### `libalat` Shared Library

`libalat` is a C-style shared library that wraps the Alat core, making it accessible from other languages. It uses an opaque pointer pattern to provide a high-level, stateful API.

### `dalat` Dart FFI Bindings

`dalat` is a Dart FFI plugin that consumes the `libalat` shared library. It provides a clean, idiomatic Dart API for Flutter applications to use, hiding all the FFI complexity.

## How It Works

Alat operates on a peer-to-peer model where each device on the network acts as both a client and a server. This enables direct communication and service sharing without relying on a central server.

### 1. Discovery

-   **mDNS**: When you start Alat on a device, it broadcasts its presence on the local network using multicast DNS (mDNS).
-   **Service Publishing**: Each Alat instance publishes a service with its device name, IP address, and port number.
-   **Listening**: Alat also listens for other devices publishing the same service, automatically discovering peers on the network.

### 2. Communication

-   **gRPC**: Once a device is discovered, a secure communication channel is established using gRPC, a high-performance RPC framework.
-   **Protocol Buffers**: The API for communication is defined using Protocol Buffers (`.proto` files). This ensures that data is serialized efficiently and that the API contract between devices is consistent.
-   **Pairing**: The first connection requires pairing to ensure that only trusted devices can connect to each other.

### 3. Services

-   **Modular Architecture**: Alat's functionality is built around a modular, service-based architecture. Each feature, like file sharing or system information, is implemented as a separate service.
-   **Service Registry**: The core of Alat includes a service registry that manages all available services.
-   **Extensibility**: This design makes it easy to add new features and services to Alat without modifying the core communication logic.

## Building from Source

To build Alat from source, you will need the following prerequisites:

*   Go 1.21+
*   Wails CLI
*   Flutter SDK
*   Fish shell
*   Protobuf compiler (`protoc`)

### Building Everything

The `mng.fish` script in the root of the project can be used to build everything.

```bash
# From the root directory:

# 1. Generate Go and Dart code from .proto files
./mng.fish proto

# 2. Build the `libalat` shared library and the `dalat` Dart package
./mng.fish dalat make

# 3. Build the desktop app
./mng.fish desktop build
```

### Building for Mobile

To work on the mobile app, you must first build the Go library and generate the Dart bindings.

```bash
# From the `packages/dalat` directory:

# 1. Build the Go library and generate the C header and Dart bindings
./mng.fish make
```

This command will:

1.  Build the `libalat` shared library.
2.  Generate the Dart FFI bindings from the C header.
3.  Generate the JSON serialization models.

Once the `dalat` plugin is built, you can run the mobile app like any standard Flutter project.

```bash
# From the `apps/mobile` directory:
flutter run
```

### Building for Desktop

```bash
# From the `apps/desktop` directory:
./mng.fish dev
```

## Contributing

Contributions are welcome! Please read the [development conventions](GEMINI.md#development-conventions) for more information.

## License

Alat is licensed under a custom license. Please see the [LICENSE](LICENSE) file for more details.
