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

### Headless Server (`apps/server`)

This application is intended to run a headless `alat` node, likely for devices without a graphical interface. It is currently a placeholder and not the primary focus.