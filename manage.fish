#!/usr/bin/env fish

# Main management script for the 'alat' project.

function help
    echo "Usage: ./manage.fish <command>"
    echo ""
    echo "Commands:"
    echo "  proto    - Generates Go and Dart code from .proto files."
    echo "  help     - Shows this help message."
    echo "  desktop  - Manage apps/desktop."
    echo "  mobile   - Manage apps/mobile."
    echo "  server   - Manage apps/server."
end

function proto
    echo "Installing protobuf generators..."
    # Install the Go generator if not already installed
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    # Install the Dart generator if not already installed
    dart pub global activate protoc_plugin

    # Add Go and Dart plugins to the PATH for this session
    fish_add_path (go env GOPATH)/bin
    fish_add_path $HOME/.pub-cache/bin

    echo "Generating Go code..."
    if not test -e pkg/core/pbuf
        echo "Creating pkg/core/pbuf folder"
        mkdir pkg/core/pbuf
    end
    protoc --proto_path=proto --go_out=pkg/core/pbuf --go_opt=paths=source_relative \
        proto/types.proto proto/service.proto

    echo "Generating Dart code..."
    # Create the output directory if it doesn't exist
    mkdir -p apps/mobile/lib/src/api/pbuf
    protoc --proto_path=proto --dart_out=apps/mobile/lib/src/api/pbuf \
        proto/types.proto proto/service.proto

    echo "Protobuf generation complete."
end


# Command dispatcher
switch "$argv[1]"
    case proto
        proto
    case help
        help
    case desktop
        cd apps/desktop
        fish manage.fish $argv[2..]
    case mobile
        cd apps/mobile
        ./manage.fish $argv[2..]
    case server
        cd apps/server
        fish manage.fish $argv[2..]
    case ""
        echo "Error: No command specified."
        help
        exit 1
    case '*'
        echo "Error: Unknown command '$argv[1]'"
        help
        exit 1
end
