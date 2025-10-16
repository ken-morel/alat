# Gemini Project: Alat Mobile Client

This document provides a comprehensive overview of the `alat` mobile application, its architecture, and development conventions, intended for use with AI development assistants like Gemini.

## Project Overview

This project is the mobile client for `alat`, a cross-platform application for seamless device communication and service sharing. It is built with Flutter and Dart, providing a native experience on both Android and iOS.

The mobile application interacts with the core `alat` logic, written in Go, through a Dart FFI (Foreign Function Interface) plugin named `dalat`.

## Architecture

### Tech Stack

*   **UI Framework**: Flutter
*   **Language**: Dart
*   **State Management**: `provider` package
*   **Navigation**: Custom routing implementation using `Navigator` and `onGenerateRoute`.
*   **Core Logic Bridge**: `dalat` (Dart FFI plugin)
*   **Localization**: `flutter_localizations` and `intl` package.

### Key Components

*   **`lib/main.dart`**: The entry point of the application. It initializes the necessary services (`NavigationService`, `NotificationService`) and the `AppState`.
*   **`lib/app.dart`**: Defines the main `AlatApplication` widget, which sets up the `MaterialApp`, themes, localization, and routing.
*   **`lib/state.dart`**: The main application state, managed using `ChangeNotifier` and `provider`. It handles the communication with the `dalat` plugin and manages the overall state of the application.
*   **`packages/dalat`**: A local package dependency that contains the Dart FFI bindings to the Go `libalat` shared library. This package is crucial for the mobile app to communicate with the core `alat` functionality.
*   **`lib/pages`**: Contains the different screens (pages) of the application, organized by feature (e.g., `dashboard`, `setup`).
*   **`lib/components`**: Contains reusable UI widgets used across different pages.
*   **`lib/services`**: Contains services like `NavigationService` for handling navigation and `NotificationService` for managing local notifications.

## Building and Running

The project includes a `mng.fish` script to simplify common development tasks.

### Prerequisites

*   Flutter SDK
*   Fish shell

### Running the App

To run the application in development mode, use the `mng.fish` script:

```bash
# To run on a selected device
./mng.fish dev

# To run on a linux device
./mng.fish dev linux
```

Alternatively, you can use the standard `flutter` commands:

```bash
# To run the app
flutter run
```

### Building the App

To build the application for a specific platform, you can use the `mng.fish` script:

```bash
# Example: Build for Android
./mng.fish build apk
```

Or use the standard `flutter` command:

```bash
flutter build <platform>
```

## Development Conventions

### State Management

The application uses the `provider` package for state management. The main application state is managed in the `AppState` class (`lib/state.dart`), which is a `ChangeNotifier`.

### Navigation

The application uses a custom routing implementation in `lib/app.dart` using the `onGenerateRoute` callback of the `MaterialApp`. A `NavigationService` (`lib/services/navigation_service.dart`) is used to manage the `Navigator`'s state.

### Localization

The application is localized using the `flutter_localizations` and `intl` packages. The localization files are located in `lib/l10n`.

### Code Style

The project follows the recommended lints from the `flutter_lints` package, as configured in `analysis_options.yaml`.
