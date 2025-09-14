# GEMINI.md

This file provides a comprehensive overview of the `alat` project, its architecture, and development conventions. It is intended to be used as a context for AI-powered development assistants like Gemini.

## Project Overview

`alat` is a cross-platform application that enables seamless communication and service sharing between devices. It operates on a peer-to-peer (P2P) model, where each device acts as both a client and a server. The project is heavily inspired by KDE Connect but does not share any code with it.

The project is in its early stages of development, so breaking changes are expected.

### Architecture

`alat` has a service-based architecture. Each device (or "node") exposes a set of services that other paired devices can consume. The communication between devices is handled by gRPC over a secure channel.

A key feature of the architecture is the use of gRPC's streaming capabilities for high-performance, low-memory operations. For example, the `filesend` service uses a client-side stream to send a file in chunks, avoiding the need to load the entire file into memory. This makes it efficient for large files and low-power devices.

The core services include:

*   **sysinfo**: Provides system information and monitoring.
*   **rcfile**: Enables remote file operations.
*   **media**: Controls media playback.
*   **notifications**: Shares notifications across devices.
*   **clipboard**: Provides a shared clipboard.
*   **filesend**: Enables high-performance, stream-based file transfers.

### Tech Stack

*   **Backend**: Go with gRPC
*   **Desktop**: Wails + SvelteKit + TypeScript + Tailwind CSS
*   **Mobile**: Flutter + Dart with Go FFI bridge (planned)
*   **Protocol**: Protocol Buffers over gRPC

## Project Structure

The project is organized as a Go workspace with the following structure:

```
alat/
├── apps/
│   ├── desktop/       # Wails + SvelteKit desktop app
│   └── server/        # Headless server (planned)
├── pkg/
│   ├── core/          # Core P2P runtime, services, client/server
│   ├── dalat/         # Shared Go packages
│   └── pbuf/          # gRPC and Protocol Buffer definitions
├── go.work           # Go workspace configuration
└── mng.fish          # Development automation scripts
```

## Building and Running

The project uses a `mng.fish` script to automate common development tasks.

### Prerequisites

*   Go 1.21+
*   Wails CLI
*   Flutter SDK (for mobile development)
*   Fish shell

### Quick Start

The main `mng.fish` script is a command dispatcher that calls the `manage.fish` scripts in the subdirectories.

#### Desktop

To build and run the desktop application, use the following commands from the `apps/desktop` directory:

```bash
# Start the development server
fish manage.fish dev

# Build the application
fish manage.fish build

# Build the application for Windows
fish manage.fish build-windows
```

#### Protocol Buffers

To generate the Go and Dart code from the `.proto` files, use the following command from the root directory:

```bash
# Generate protocol buffer code
fish mng.fish proto
```

**Note**: The `dev` commands are targeted for devices running Ubuntu 25.04 on amd64.

## Development Conventions

### Backend

The backend is written in Go and follows standard Go conventions. The code is organized into packages, with each package having a specific responsibility.

**Service Structure**: Each core service (e.g., `sysinfo`, `filesend`) is self-contained within its own package under `pkg/core/service/`. For cleanliness and modularity, the gRPC server implementation for a service is co-located within its package (e.g., `pkg/core/service/filesend/server.go`).

### Frontend

The frontend of the desktop application is a SvelteKit project. It uses TypeScript for type safety and Tailwind CSS for styling. The code is organized into components, with each component having its own file.

### Commits

The project does not have a strict commit message format, but it is recommended to write clear and concise commit messages that explain the changes made.
