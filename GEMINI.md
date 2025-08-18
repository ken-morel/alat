# Project: Alat

Alat is a set of tools created by your user `ama`, which aims at connecting and exchanging services
seamlessly between all kinds of devices. The project is a monorepo featuring mainly a go workspace
(see [./go.work]) linking the several contained go applicaitons.

Go is the main language for backend and application programming. Alat currently seperates into few parts:

- `./proto/`: Contains gRPC and protocol buffer definitions used by alat between communicating devices.
- `./pkg`: Stores packages used by other applicaitons.
  - `./pkg/core/`: Stores alat core package holding services client and server definitions, p2p server-client runtime.
    Services are abit more detailed in `./pkg/core/services.md`.
- `./logo`: Stores multiple sized logo definitions of two kinds:
  - `./logo/animated/`: Stores the animated svg logo(`./logo/animated/logo.svg`)
  - `./logo/static/`: Stores static svg and png logos

## Application Architecture

Alat operates on a peer-to-peer (P2P) model. Each application instance, whether on desktop or mobile, acts as both a client and a server.

### Desktop App (`apps/desktop`)

The desktop application, built with Wails and Svelte, is the primary reference implementation. It is a full `alat` node:

-   **P2P Server:** On startup, it initializes and runs the core P2P server defined in `pkg/core/server`. This exposes the device's services (`sysinfo`, `rcfile`, etc.) to other `alat` nodes on the local network.
-   **Client:** It includes client-side logic (`pkg/core/client`) to discover other devices, manage pairing, and consume services from paired devices.
-   **Configuration:** It manages all device and service configurations, storing them in a local configuration directory (`~/.config/alat`).

### Mobile App (`apps/mobile`)

The mobile application is built with Flutter. It is currently under development and will implement the same client and server P2P functionality as the desktop app.

-   **UI:** The app uses the `provider` package for state management and is set up for internationalization with support for English and French.
-   **Theming:** It features a dark theme consistent with the desktop application.
-   **Go Bridge:** A placeholder Go package (`pkg/mobile_bridge`) has been created to facilitate communication between the Dart frontend and the Go `alat` core via FFI.

### Headless Server (`apps/server`)

This application is intended to run a headless `alat` node, likely for devices without a graphical interface. It is currently a placeholder and not the primary focus.

## Development Scripts (`manage.fish`)

The root `manage.fish` script is the main entry point for common development tasks. It can be used to orchestrate builds, generate code, and run the different applications.

-   `./manage.fish proto`: Compiles protobuf files (`.proto`) into Go and Dart code.
-   `./manage.fish desktop <command>`: Delegates to the desktop app's specific `manage.fish` script.
-   `./manage.fish mobile <command>`: Delegates to the mobile app's specific `manage.fish` script. For example, `./manage.fish mobile dev` will run the Flutter app in development mode.
-   `./manage.fish server <command>`: Delegates to the server app's `manage.fish` script.