# Alat: Your Devices, Unified.

<p align="center">
  <img src="./logo.png" alt="Alat Logo" width="200"/>
</p>

Alat is a set of tools made to work for several platforms which provides
devices to share services together.

## Components

Alat constutes of a few parts including:

### Alat core

Core of the project, alat core provides most of alat services in a cross-platform
manner(best as possible). Alat core is written in :go: and contains logic and
procedures for:
- basic security using certificates and id's
- discovery using mdns, can be deactivated for platforms like `android` where
  that is not possible.
- alat common services, have their implementation, configuration, grpc servers
  and client there organised in a modularized manner.

The features of alat core are exposed through a `libalat` aimed at providing an easy way to integrate alat into other applications and is used for mobile application. It exposes a C-compatible API and language-specific bindings that make it easy to embed Alat functionality into desktop, mobile, and embedded applications. For mobile platforms we provide "dalat", an FFI wrapper used to integrate libalat into Android  builds when platform-specific code is required..

### Applications

![Alat desktop application screenshot](./media/alat-settings-device-shot.png)

A desktop application tested on windows and linux.
A mobile application is also on it's way but slowed down due to android limitations forcing me to crafttailored, platform-specific alternatives, the mobile application relies on `dalat`, providing ffi bindings for `libalat`.


## Services

What's makes alat alat, is the fact it provides several services on most if not all of it's supported platforms.

Roadmap:

### Implemented
- [x] **Device Information**: Share basic device details.
- [x] **File Sharing**
  - [x] Direct device-to-device transfer.
    - [ ] File verification.
    - [ ] File partitioning for continuous verification.
  - [x] Web sharing via a local server.

### Phase 1: Core Syncing Features
1. [ ] **Clipboard Synchronization**: Universal copy and paste.
2. [ ] **Notification Synchronization**: Mirror notifications across devices.
3. [ ] **Share Text/URL Snippet**: Send text directly to another device.

### Phase 2: Remote Control & Input
4. [ ] **Media Control**: Manage media playback remotely.
5. [ ] **Presentation Control**: Use your phone as a presentation remote.
6. [ ] **Remote Keyboard & Mouse**: Use your phone as a wireless keyboard/mouse.
7. [ ] **Run Commands**: Trigger predefined commands on your devices.

### Phase 3: Advanced Filesystem Operations
8. [ ] **File System Exploration**: Browse remote files.
9. [ ] **Folder Synchronization**: Keep folders in sync between devices.

### Phase 4: Future Enhancements
10. [ ] **Screen Casting**: Share your screen to another device.
11. [ ] **File Transfer Enhancements**
    - [ ] File verification with checksums.
    - [ ] Partitioning for large files.
