#!/usr/bin/env fish

# Main management script for the 'alat' project.

function help
    echo "Usage: ./manage.fish <command>"
    echo ""
    echo "Commands:"
    echo "  proto    - Generates Go and Dart code from .proto files."
    echo "  help     - Shows this help message."
end

function proto
    echo "Installing protobuf generators..."
    # Install the Go generator if not already installed
    # go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    # Install the Dart generator if not already installed
    #pub global activate protoc_plugin
    #TODO: Uncomment

    # Add Go and Dart plugins to the PATH for this session
    fish_add_path (go env GOPATH)/bin
    fish_add_path $HOME/.pub-cache/bin

    echo "Generating Go code..."
    protoc --proto_path=proto --go_out=pkg/core/protobuf --go_opt=paths=source_relative \
        proto/types.proto proto/service.proto

    # echo "Generating Dart code..."
    # # Create the output directory if it doesn't exist
    # mkdir -p apps/mobile/lib/src/proto
    # protoc --proto_path=proto --dart_out=apps/mobile/lib/src/proto \
    #     proto/types.proto proto/service.proto

    echo "Protobuf generation complete."
end

function dev-desktop
    cd apps/desktop
    wails dev -tags webkit2_41
end

# Command dispatcher
switch "$argv[1]"
    case proto
        proto
    case help
        help
    case dev-desktop
        dev-desktop
    case ""
        echo "Error: No command specified."
        help
        exit 1
    case '*'
        echo "Error: Unknown command '$argv[1]'"
        help
        exit 1
end
