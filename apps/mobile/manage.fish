#!/usr/bin/env fish

function help
    echo "Usage: ./manage.fish <command>"
    echo ""
    echo "Commands:"
    echo "  l10n     - Generates Flutter localization files."
    echo "  build_bridge - Compiles the Go FFI bridge."
    echo "  dev      - Builds the bridge and runs the Flutter app."
    echo "  build    - Builds the Flutter app for the Linux desktop."
    echo "  help     - Shows this help message."
end

function l10n
    echo "Generating Flutter localization files..."
    flutter gen-l10n
end

function build_bridge
    echo "Building Go bridge for FFI..."
    go build -buildmode=c-shared -o build/libalat.so ../../pkg/mobile_bridge/bridge.go
end

function dev
    echo "Building bridge..."
    build_bridge
    echo "Running Flutter app for Linux in development mode..."
    flutter run -d linux
end

function build
    echo "Building Flutter app for Linux..."
    flutter build linux
end

switch "$argv[1]"
    case l10n
        l10n
    case dev
        dev
    case build
        build
    case help
        help
    case ""
        echo "Error: No command specified."
        help
        exit 1
    case '*'
        echo "Error: Unknown command '$argv[1]'"
        help
        exit 1
end
