<h1 style="text-align: center; font-size: xxx-large;">
  <img src="./logo/animated/logo.svg" alt="Animated logo" style="display: inline;height: 1em;position: relative;top: 5px; right: -10px;" />
  lat
</h1>

[![wakatime](https://wakatime.com/badge/github/ken-morel/alat.svg)](https://wakatime.com/badge/github/ken-morel/alat)


Alat is a set of services and applications aimed at providing a seamless interface common to all devices to connect them together. Its work is greatly inspired by [KDE Connect](https://kdeconnect.kde.org/), though they don't have any direct relationship.

For it to work, Alat exposes features to paired devices through services, which you can configure and control.

**⚠️ Early Development Notice**: Alat has barely started its development, and very breaking changes are expected.

## Architecture Overview

Alat operates on a **peer-to-peer (P2P) model** where each device acts as both a client and server, enabling seamless bidirectional communication and service sharing.

```
┌─────────────────┐    gRPC/Protocol Buffers    ┌─────────────────┐
│   Desktop App   │◄──────────────────────────►│   Mobile App    │
│                 │                              │                 │
│ ┌─────────────┐ │                              │ ┌─────────────┐ │
│ │P2P Server   │ │          Discovery           │ │P2P Server   │ │
│ │- sysinfo    │ │◄────────(mDNS/Local)───────►│ │- sysinfo    │ │
│ │- rcfile     │ │                              │ │- rcfile     │ │
│ │- media      │ │                              │ │- media      │ │
│ └─────────────┘ │                              │ └─────────────┘ │
│ ┌─────────────┐ │                              │ ┌─────────────┐ │
│ │P2P Client   │ │         Service Calls        │ │P2P Client   │ │
│ │- discovery  │ │◄──────────────────────────►│ │- discovery  │ │
│ │- pairing    │ │                              │ │- pairing    │ │
│ └─────────────┘ │                              │ └─────────────┘ │
└─────────────────┘                              └─────────────────┘
         ▲                                                ▲
         │                                                │
         ▼                                                ▼
┌─────────────────┐                              ┌─────────────────┐
│ Headless Server │                              │  Other Devices  │
│                 │                              │   (TV, IoT...)  │
└─────────────────┘                              └─────────────────┘
```

### Service-Based Architecture

Each Alat node exposes **services** that other paired devices can consume:

- **sysinfo**: System information and monitoring
- **rcfile**: Remote file operations
- **media**: Media playback control
- **notifications**: Cross-device notifications
- **clipboard**: Shared clipboard functionality
- _(More services planned...)_

## Project Structure

```
alat/
├── proto/              # gRPC and Protocol Buffer definitions
├── pkg/                # Shared Go packages
│   ├── core/          # Core P2P runtime, services, client/server
│   └── mobile_bridge/ # Go-Dart FFI bridge for mobile
├── apps/
│   ├── desktop/       # Wails + SvelteKit desktop app
│   ├── mobile/        # Flutter mobile app
│   └── server/        # Headless server (planned)
├── logo/
│   ├── animated/      # Animated SVG logos
│   └── static/        # Static SVG and PNG logos
├── go.work           # Go workspace configuration
└── manage.fish       # Development automation scripts
```

## Features (Planned)

- 🔗 **Seamless P2P Connection**: Automatic device discovery and pairing
- 🔒 **Secure Communication**: Encrypted gRPC channels
- 🎛️ **Service Framework**: Modular, extensible service system
- 🖥️ **Cross-Platform**: Desktop (Linux/Windows/macOS) and Mobile (Android/iOS)
- ⚙️ **Configurable**: Per-device and per-service configuration
- 🔄 **Real-time Sync**: Live service state synchronization

## Building

The app should be pretty easy to build. Make sure you have the Go programming language installed, then use it to install the Wails CLI (check https://wails.io), then just build as a normal Wails application or using the `manage.fish` scripts.

### Prerequisites

- [Go](https://golang.org/) 1.21+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)
- For mobile: [Flutter SDK](https://flutter.dev/docs/get-started/install)

### Quick Start

```bash
# Build and run desktop app
./manage.fish desktop build
./manage.fish desktop dev

# Build mobile app (development)
./manage.fish mobile dev

# Generate protocol buffer code
./manage.fish proto
```

**Notice**: The dev commands are targeted for devices running Ubuntu 25.04 on amd64.

## Development Scripts

The root `manage.fish` script orchestrates common development tasks:

- `./manage.fish proto` - Compile protobuf files to Go and Dart
- `./manage.fish desktop <command>` - Desktop app operations
- `./manage.fish mobile <command>` - Mobile app operations  
- `./manage.fish server <command>` - Headless server operations

## Contributing

Alat is structured as a Go workspace with a core package and multiple applications. A more detailed explanation of how the project works internally can be found in [./GEMINI.md](./GEMINI.md).

### Current Status

- ✅ Project structure and architecture defined
- ✅ gRPC protocol definitions
- 🚧 Desktop app (Wails + SvelteKit + TypeScript + Sass)
- 🚧 Mobile app (Flutter with Go FFI bridge)
- 📋 Headless server (planned)
- 📋 Core services implementation

### Tech Stack

- **Backend**: Go with gRPC
- **Desktop**: Wails + SvelteKit + TypeScript + Sass
- **Mobile**: Flutter + Dart with Go FFI bridge
- **Protocol**: Protocol Buffers over gRPC
- **Architecture**: Peer-to-peer with service-oriented design

---

**Thanks for your interest in Alat!** 🚀