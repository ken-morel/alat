# Alat

![Animated logo](./logo/animated/logo.svg)

Alat is a cross-platform application that allows you to seamlessly connect and share data between your devices. Whether you're using a desktop computer, a mobile phone, or a tablet, Alat helps you stay connected and productive.

## How it Works

Alat uses a combination of technologies to discover and communicate with devices on your local network. It consists of the following components:

- **A core daemon** that runs on each device and handles the discovery and communication.
- **A desktop application** for Windows, macOS, and Linux that allows you to manage your connected devices.
- **A mobile application** for iOS and Android that provides the same functionality on the go.
- **A server** that can be used for remote discovery and relay, allowing you to connect to your devices from anywhere in the world.

## Features

- **Automatic device discovery**: Alat automatically discovers other devices on your local network that are running the application.
- **Cross-platform**: Alat works on Windows, macOS, Linux, iOS, and Android.
- **Secure**: All communication between devices is encrypted.
- **Open source**: Alat is open source and available on GitHub.

## Getting Started

If you're a developer and want to contribute to the project, here's how to get started:

### 1. Install Protocol Buffers Compiler

You need the `protoc` compiler to generate code from the `.proto` files.

- **macOS:** `brew install protobuf`
- **Linux:** `sudo apt-get install -y protobuf-compiler`
- **Windows:** `choco install protoc` or download from the [official releases page](https://github.com/protocolbuffers/protobuf/releases).

### 2. Generate Protocol Code

Use the `manage.fish` script to generate the necessary Go and Dart code from the protocol definitions.

```fish
./manage.fish proto
```

This will:

1. Install the necessary Go and Dart plugins.
2. Generate Go code in `pkg/core/proto/`.
3. Generate Dart code in `apps/mobile/lib/src/proto/`.

You can see all available commands by running:

```fish
./manage.fish help
```

Now you are ready to start developing!

## Contributing

We welcome contributions from the community. If you'd like to contribute, please fork the repository and submit a pull request.

## License

Alat is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
