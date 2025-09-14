# GEMINI.md

This file provides a comprehensive overview of the `alat` project, its architecture, and development conventions. It is intended to be used as a context for AI-powered development assistants like Gemini.

## Project Overview

`alat` is a cross-platform application that enables seamless communication and service sharing between devices. It operates on a peer-to-peer (P2P) model, where each device acts as both a client and a server. The project is heavily inspired by KDE Connect but does not share any code with it.

## Architecture

`alat` has a modular, service-based architecture. The core logic is written in Go and is shared across all platform-specific applications.

### Core & FFI

*   **`pkg/core`**: A pure Go module containing the core P2P runtime, services, and business logic. It is completely platform-agnostic.
*   **`pkg/libalat`**: A Go module that acts as a C-style shared library wrapper around the `core` module. It uses the "Opaque Pointer" (or "Handle") pattern to expose a high-level, stateful API. All complex data is serialized to JSON for safe and flexible transport across the FFI boundary.
*   **`packages/dalat`**: A distributable Dart FFI plugin that consumes the `libalat` shared library. It provides a clean, idiomatic Dart API (`AlatInstance`) for Flutter applications to use, hiding all the FFI complexity.

### Platform-Specific Applications

*   **`apps/desktop`**: A Wails application (Go + SvelteKit) for desktop platforms. It interacts directly with the `pkg/core` module.
*   **`apps/mobile`**: A Flutter application for mobile platforms. It interacts with the Go core exclusively through the `dalat` plugin.

### Tech Stack

*   **Core Logic**: Go
*   **P2P Communication**: gRPC over a secure channel
*   **Desktop Frontend**: Wails + SvelteKit + TypeScript + Tailwind CSS
*   **Mobile Frontend**: Flutter + Dart (using the `provider` package for state management)
*   **FFI Bridge**: Go (`cgo`) -> C Shared Library -> Dart (`dart:ffi`)

## Project Structure

```
alat/
├── apps/
│   ├── desktop/       # Wails + SvelteKit desktop app
│   └── mobile/        # Flutter mobile app
├── packages/
│   └── dalat/         # Distributable Dart FFI plugin
├── pkg/
│   ├── core/          # Core P2P runtime and services
│   ├── libalat/       # Go FFI shared library wrapper
│   └── pbuf/          # gRPC and Protocol Buffer definitions
├── go.work
└── mng.fish
```

## Building and Running

The project uses `mng.fish` scripts to automate common development tasks.

### Prerequisites

*   Go 1.21+
*   Wails CLI
*   Flutter SDK
*   Fish shell

### Building `libalat` and `dalat`

To work on the mobile app, you must first build the Go library and generate the Dart bindings.

```bash
# From the `packages/dalat` directory:

# 1. Build the Go library and generate the C header
./mng.fish build libalat

# 2. Generate the Dart bindings from the header
dart run ffigen --config ffigen.yaml

# 3. Generate JSON serialization models
dart run build_runner build --delete-conflicting-outputs
```

### Running the Desktop App

```bash
# From the `apps/desktop` directory:
fish manage.fish dev
```

### Running the Mobile App

Once the `dalat` plugin is built, you can run the mobile app like any standard Flutter project.

```bash
# From the `apps/mobile` directory:
flutter run
```

## Development Conventions

### Git Workflow: Simplified Git Flow

This project uses a structured branching strategy for maintainability.

*   **`main`**: Always stable and represents the latest official release. Commits are only merged from `release/*` or `hotfix/*` branches.
*   **`dev`**: The primary development branch for the next release. All feature branches are merged into `dev`.
*   **`feature/<name>`**: Branched from `dev`. For developing new features. Merged back to `dev` via a Pull Request.
*   **`release/vX.X.X`**: Branched from `dev`. For release stabilization (bug fixes, version bumps). Merged into both `main` and `dev`.
*   **`hotfix/<name>`**: Branched from `main`. For critical production bug fixes. Merged into both `main` and `dev`.

### Backend Code

*   The backend is written in Go and follows standard Go conventions.
*   Core logic should be platform-agnostic and placed in `pkg/core`.
*   Platform-specific code or UI-related logic should be in the respective `apps/*` directory.
*   Libraries intended for cross-language use (like `libalat`) should not make assumptions about the caller's filesystem and should receive explicit paths from the caller.