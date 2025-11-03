# Alat: Your Devices, Unified.

<p align="center">
  <img src="./logo.png" alt="Alat Logo" width="200"/>
</p>

**Alat is a powerful, cross-platform tool that creates a seamless bridge between all your devices.** Inspired by the functionality of KDE Connect, Alat is built from the ground up to be a modern, fast, and intuitive way to make your digital life more connected.

Whether you're at your desktop or on the go with your phone, Alat allows your devices to talk to each other, share information, and work together as one.

## What Can Alat Do For You?

Alat is more than just a utility; it's a suite of services designed to streamline your workflow. As the project evolves, more services will be added, but here's what you can do right now:

*   **Send Files Instantly**: Quickly send files from your phone to your PC or from your laptop to your tablet. Alat handles the transfer seamlessly over your local network, no cloud storage required.

*   **Share with Anyone (WebShare)**: Need to share a file with someone who doesn't have Alat? The WebShare service starts a mini web server on your device. Just give someone the link and a passcode, and they can download files from—or upload files to—your device directly through their web browser. It's perfect for quick, secure sharing with friends or colleagues on the same network.

*   **Cross-Platform Native Experience**: Alat is designed to feel at home on every operating system. It provides a native desktop application for Windows, macOS, and Linux, and a fluid mobile app for Android and iOS.

*   **OS Integration**: On mobile, Alat integrates directly into the Android share sheet. This means you can share files from any app directly to Alat, choosing whether to send them to another device or add them to a WebShare session.

## How Does It Work?

Alat works its magic by creating a secure peer-to-peer (P2P) network between your devices. When you install Alat on your phone and your computer, they automatically discover each other on your local Wi-Fi network. After a simple one-time pairing process to ensure security, your devices are connected.

There's no central server and your data never leaves your local network, ensuring your privacy and security.

## The Evolution of Alat

Alat is a project in active development, constantly growing and improving. The goal is to build a comprehensive and indispensable tool for multi-device workflows. Here’s a glimpse of the journey so far:

*   **Foundation**: Established the core P2P communication layer using gRPC and a robust service-based architecture in Go.
*   **File Sharing**: Implemented the initial `FileSend` service, allowing direct device-to-device transfers.
*   **WebShare Service**: Introduced a powerful new way to share files with non-Alat devices via a web browser, complete with passcode protection and a modern UI.
*   **Mobile Integration**: Made Alat a native share target on Android, allowing any app to share files directly into the Alat ecosystem.
*   **UI/UX Refinements**: Continuously improving the user experience on both desktop and mobile, with a focus on intuitive design and native feel.

Future development will focus on adding more services (like clipboard sharing, notification sync, and remote control) and further refining the user experience across all platforms.

## Getting Started

To get started with Alat, you can build the applications from the source code.

### Prerequisites

*   Go 1.21+
*   Wails CLI
*   Flutter SDK
*   Fish shell
*   Protobuf compiler (`protoc`)

### Building the Apps

For detailed instructions on how to build the desktop and mobile applications, please refer to the [**Building from Source**](GEMINI.md#building-and-running) section in our technical documentation.

## Contributing

Alat is an open-source project, and contributions are always welcome! Whether it's by reporting a bug, suggesting a new feature, or writing code, you can help make Alat better. Please read the [development conventions](GEMINI.md#development-conventions) to get started.

## License

Alat is licensed under a custom license. Please see the [LICENSE](LICENSE) file for more details.