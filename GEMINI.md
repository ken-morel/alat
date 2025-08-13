# Gemini Code Assistant Context

This document provides context for the Gemini Code Assistant to understand the `alat` project.

## Project Overview

`alat` is a cross-platform application designed to connect devices. It consists of a core Go daemon, a desktop application, a mobile application, and a server for remote discovery and relay.

The project is structured as a monorepo with the following components:

*   **`proto/`**: Contains the Protocol Buffer definitions that define the communication schema between the different parts of the application.
*   **`pkg/core/`**: The core Go logic shared across all platforms.
*   **`apps/`**: Contains the platform-specific applications:
    *   **`desktop/`**: A Wails application for Windows, macOS, and Linux. The frontend is built with Svelte.
    *   **`mobile/`**: A Flutter application for iOS and Android.
    *   **`server/`**: A Go server for remote discovery and relay.
*   **`go.work`**: The Go workspace file that manages the Go modules in the project.
*   **`manage.fish`**: The main project management script, used for tasks like generating code from the protocol buffers.

## Building and Running

The project is managed using the `manage.fish` script. The following commands are available:

*   **`./manage.fish help`**: Shows the help message with all available commands.
*   **`./manage.fish proto`**: Generates the Go and Dart code from the `.proto` files. This is a crucial step before building the applications.
*   **`./manage.fish desktop <command>`**: Manages the desktop application.
*   **`./manage.fish mobile <command>`**: Manages the mobile application.
*   **`./manage.fish server <command>`**: Manages the server application.

### Building and Running the Desktop App

To build and run the desktop app, you will need to have the Wails CLI installed. Then, you can use the following commands:

```bash
cd apps/desktop
wails dev
```

### Building and Running the Mobile App

To build and run the mobile app, you will need to have the Flutter SDK installed. Then, you can use the following commands:

```bash
cd apps/mobile
flutter run
```

### Building and Running the Server

To build and run the server, you can use the following command:

```bash
cd apps/server
go run .
```

## Development Conventions

*   **Protocol Buffers**: All communication between the different parts of the application is defined using Protocol Buffers. When making changes to the communication protocol, you must first update the `.proto` files and then run `./manage.fish proto` to regenerate the code.
*   **Code Style**: The project uses the standard Go and Dart code styles.
*   **Testing**: TODO: Add information about the testing strategy.
