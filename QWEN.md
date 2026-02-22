# Alat Project Documentation

## Project Overview

Alat is a Rust-based device unification project with the slogan **"That's one device that's all your devices."** It represents a complete reimagination of the original Alat project, moving from a complex multi-component architecture to a simpler, more robust design that prioritizes system integration over standalone applications.

The project aims to unify devices into a single computational fabric, where devices are not merely "connected" but truly integrated regardless of their type (as long as they support the Alat protocol).

## Architecture

### Three-Layer Structure

1. **`nlem` (Core)**: The "heart" of the system (named `nlem` instead of `core` to follow Rust crate naming conventions)
   - Implements the central `Node` coordinator
   - Provides `Services` (modular functionality units) and `Controllers` (components that consume services)
   - Handles device management, security, discovery, and communication protocols
   - Uses channel-based communication instead of shared mutable state

2. **`platform` (Platform Abstraction)**: OS-specific bindings
   - Currently implements Linux support
   - Provides platform-specific functionality (hostname, config directories, notifications, etc.)
   - Uses `cfg-if` for conditional compilation across platforms

3. **`app` (UI/Application)**: Cross-platform interface
   - Built with Slint for Material Design UI
   - Minimal interface focused on initial setup and monitoring
   - Designed to integrate with the operating system rather than be a standalone application

### Communication Model

- **Channel-based messaging**: Components communicate via dedicated channels instead of shared state
- **Node-centric**: The `Node` acts as the central coordinator that processes events and relays messages
- **Event-driven**: UI receives `NodeEvent` notifications to maintain state

### Service-Controller Pattern

- **Services**: Implement the `Service` trait and may provide:
  - Worker threads/tasks
  - gRPC servers
  - Client implementations
  - Storage partitions for cache/config/temp
- **Controllers**: Consume services to perform operations like:
  - Clipboard synchronization
  - File transfer
  - Notification forwarding
  - Media control

## Technology Stack

- **Language**: Rust (edition 2021/2024)
- **UI Framework**: Slint (for cross-platform GUI with Material Design)
- **Async Runtime**: Tokio (with multi-threaded runtime)
- **RPC/Protocol**: Tonic + Prost (gRPC with Protocol Buffers)
- **Data Serialization**: Serde + serde_json
- **Platform Integration**: Various crates for system integration (hostname, dirs, notify-rust, battery, sysinfo)

## Key Features

### Discovery Protocol
- UDP broadcast on port 4147 for device discovery
- Devices advertise every 5 seconds
- Devices considered lost after 15 seconds of no advertisement
- More reliable and simpler than mDNS while maintaining cross-platform compatibility

### Security Model
- Each device has a unique ID and certificate
- Pairing involves secure certificate exchange with user confirmation
- User approval via desktop notifications (Linux) ensures security

### Data Storage
- Configuration and paired device data stored in JSON format
- Location: `~/.config/cm.engon.alat/data.json` on Linux
- Structured for easy migration and backup

## Development Environment

### Build Tools
- **Bacon**: Configured in `bacon.toml` as the primary development tool
  - `bacon run`: Run the application
  - `bacon check`: Check for compilation errors
  - `bacon clippy`: Run Clippy linter
  - `bacon test`: Run tests
  - `bacon doc`: Generate documentation

### Cargo Workspace
- Defined in `Cargo.toml` with members: `app`, `nlem`, `platform`
- Shared dependencies defined in workspace section
- Uses resolver version 3

### Key Dependencies
- **app**: slint, platform, nlem, tokio
- **nlem**: prost, serde, tonic, rand, thiserror
- **platform**: cfg-if, hostname, dirs, notify-rust, sysinfo, battery

## Project Status

The project is in early development stages. As noted in the README, the developer acknowledges being "not that fast as a rust beginner" and that the project is "nothing but starting." Key areas that need development include:

- Implementation of core services (pairing, telemetry, etc.)
- Development of controllers (clipboard sync, file transfer, etc.)
- Expansion to additional platforms (macOS, Windows)
- Mobile client integration
- Advanced discovery options (Bluetooth, Wi-Fi Direct)

## Development Conventions

- **Code Style**: Follows Rust standard formatting (`cargo fmt`)
- **Error Handling**: Uses `thiserror` for custom error types
- **Async Patterns**: Uses `tokio::main` for async main functions
- **State Management**: Uses `Arc<RwLock<T>>` for shared state with channel-based communication
- **Testing**: Currently limited; the developer notes "You will notice I'm not the one to write tests, they're even worse than docs."

## Contact and Resources

- **Email**: engonken8@gmail.com or me@engon.cm
- **Alat Project Website**: https://alat.engon.cm
- **Developer Portfolio**: https://engon.cm
- **Original Project**: https://github.com/ken-morel/alat-old

## Important Notes

- The name "Alat" means "to link" in some languages, reflecting the application's purpose
- "nlem" means "heart" in the developer's naming convention
- The project follows a custom license that permits educational use and contributions but prohibits commercial use without explicit permission
- The developer humorously notes that "whatever way you try to pronounce 'alat' and especially 'nlem', your are surely not doing it right"

This documentation provides a comprehensive overview of the Alat project for future development and maintenance work.