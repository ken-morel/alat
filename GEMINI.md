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
