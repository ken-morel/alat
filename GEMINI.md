# Gemini Project Context: Alat

## Project Overview

This repository contains the source code for "Alat", a cross-platform application designed to connect devices for file exchange, screen mirroring, clipboard syncing, and more.

The project is structured as a monorepo and utilizes a core-satellite architecture:
-   **`pkg/core`**: A central, headless daemon written in Go that contains all the core business logic, networking, and service discovery. This is the "brain" of the application.
-   **`apps/`**: This directory houses the platform-specific "heads" or user interfaces.
    -   **`apps/desktop`**: A desktop application built with Wails (Go + Svelte/TypeScript).
    -   **`apps/mobile`**: A mobile application for iOS and Android built with Flutter/Dart.
    -   **`apps/server`**: A Go-based web server that acts as a relay and discovery mechanism for devices not on the same local network.
-   **`proto/`**: Contains the Protocol Buffers (`.proto`) definitions that define the data structures and service contracts for all communication within the Alat ecosystem. This ensures type-safe and efficient data exchange between the Go core, the Flutter app, and other components.

The entire Go codebase is managed as a single unit via a `go.work` workspace file.

## Building and Running

The project uses a unified management script, `manage.fish`, to handle common development tasks.

### 1. Prerequisites

Before running any commands, you must have the Protocol Buffers compiler (`protoc`) installed.
-   **macOS:** `brew install protobuf`
-   **Linux:** `sudo apt-get install -y protobuf-compiler`
-   **Windows:** `choco install protoc`

### 2. Generating Protocol Code

To ensure all parts of the application can communicate, you must first generate the necessary Go and Dart code from the `.proto` files.

```fish
./manage.fish proto
```
This command will:
1.  Install the required `protoc` plugins for Go and Dart.
2.  Generate Go code in `pkg/core/proto/`.
3.  Generate Dart code in `apps/mobile/lib/src/proto/`.

### 3. Running the Applications

-   **Desktop App**:
    ```bash
    # Navigate to the desktop app directory
    cd apps/desktop

    # Run the Wails development server
    wails dev
    ```

-   **Mobile App**:
    ```bash
    # Navigate to the mobile app directory
    cd apps/mobile

    # Run the Flutter application
    flutter run
    ```

-   **Server**:
    ```bash
    # Navigate to the server app directory
    cd apps/server

    # Run the Go server
    go run .
    ```

## Development Conventions

-   **Monorepo**: All code for the project is located in this single repository.
-   **Protocol-First**: Any new feature or data structure change should begin with an update to the `.proto` files. After updating, run `./manage.fish proto` to propagate the changes across the codebase.
-   **Go Workspace**: The Go modules (`core`, `desktop`, `server`) are linked in the `go.work` file, allowing for easy cross-module development.
-   **Centralized Logic**: All core business logic should reside in the `pkg/core` module. The platform-specific apps in `apps/` should be as "thin" as possible, primarily handling UI and delegating logic to the core module.
