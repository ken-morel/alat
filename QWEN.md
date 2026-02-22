# Alat Project Documentation

## Project Overview

Alat is a cross-platform application for device pairing and management, built with Rust. The project follows a modular architecture with three main components:

1. **app**: The user interface layer, built with Slint for cross-platform GUI development
2. **nlem**: The core logic layer, handling device management, discovery, security, and communication protocols
3. **platform**: Platform-specific bindings for Linux (currently only Linux is supported)

The application appears to focus on discovering, pairing with, and managing devices through a clean, modern UI.

## Technology Stack

- **Language**: Rust (edition 2021/2024)
- **UI Framework**: Slint (for cross-platform GUI)
- **Async Runtime**: Tokio (with multi-threaded runtime)
- **RPC/Protocol**: Tonic + Prost (gRPC with Protocol Buffers)
- **Data Serialization**: Serde + serde_json
- **Platform Integration**: Various crates for system integration (hostname, dirs, notify-rust, battery, sysinfo)

## Project Structure

```
alat/
├── Cargo.toml          # Workspace definition
├── bacon.toml          # Configuration for the bacon build tool
├── app/                # UI application
│   ├── src/main.rs     # Application entry point
│   ├── ui/             # Slint UI definitions
│   │   ├── app-window.slint  # Main window UI
│   │   └── models.slint      # Data models for UI
│   └── Cargo.toml      # App dependencies
├── nlem/               # Core logic
│   ├── src/lib.rs      # Core module definitions
│   ├── src/client/     # Client-side logic
│   ├── src/devicemanager/  # Device management
│   ├── src/discovery/  # Device discovery
│   ├── src/node/       # Node implementation
│   ├── src/platform/   # Platform abstraction
│   ├── src/proto/      # Protocol definitions
│   ├── src/security/   # Security functionality
│   ├── src/server/     # Server implementation
│   ├── src/service/    # Service management
│   ├── src/storage/    # Data storage
│   └── Cargo.toml
├── platform/           # Platform-specific code
│   ├── src/lib.rs      # Platform abstraction
│   ├── src/platform.rs # Linux implementation
│   ├── src/discovery/  # Discovery implementation
│   ├── src/storage/    # Storage implementation
│   ├── src/telemetry/  # Telemetry collection
│   └── Cargo.toml
├── logo/               # Logo assets
└── LICENSE             # Custom software license
```

## Building and Running

### Prerequisites
- Rust toolchain (latest stable)
- Cargo
- For UI development: Slint build tools

### Build Commands
The project uses `bacon` as a build tool (configured in `bacon.toml`). Common commands include:

- `bacon check` - Run cargo check (default job)
- `bacon clippy` - Run Clippy linter
- `bacon test` - Run tests
- `bacon run` - Run the application
- `bacon doc` - Generate documentation

Alternatively, you can use standard Cargo commands:
- `cargo check` - Check for compilation errors
- `cargo build` - Build the project
- `cargo run` - Run the application
- `cargo test` - Run tests

### Running the Application
```bash
# Using bacon (recommended)
bacon run

# Using cargo directly
cargo run --package app
```

## Development Conventions

### Code Style
- Follows Rust standard formatting (use `cargo fmt`)
- Uses `tokio::main` for async main functions
- Uses Arc/RwLock for shared state management
- Follows idiomatic Rust patterns for error handling

### Testing
- Tests are organized in each crate's `tests/` directory (not yet visible in current structure)
- Uses `cargo test` for running tests
- Supports nextest for more advanced testing scenarios

### UI Development
- Uses Slint for declarative UI definition
- UI components are defined in `.slint` files
- Data binding between Rust and Slint is handled through property declarations and callbacks

## Key Features

1. **Device Discovery**: Automatically discovers available devices on the network
2. **Device Pairing**: Secure pairing mechanism with user confirmation
3. **Cross-Platform UI**: Modern UI using Slint with Material Design components
4. **Telemetry Collection**: Collects system information for diagnostics
5. **Platform Abstraction**: Clean separation between core logic and platform-specific code

## License Information

The project uses a custom software license that permits educational use, personal modification, and contributions, but explicitly prohibits commercial use without explicit written permission.

For commercial licensing inquiries, contact: engon@engon.cm

## Contribution Guidelines

- All contributions become part of the original work under the same license
- Contributors retain credit for their specific contributions
- The project maintainer reserves the right to accept or reject contributions
- Please include proper attribution when sharing the software

## Architecture Overview

The application follows a layered architecture:
- **UI Layer (app)**: Handles user interaction with Slint
- **Core Logic Layer (nlem)**: Implements business logic, device management, security, and communication protocols
- **Platform Layer (platform)**: Provides OS-specific functionality (Linux currently supported)

The core components communicate through well-defined interfaces, with the `Node` struct acting as the central coordinator between the UI and backend services.