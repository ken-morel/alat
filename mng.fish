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
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

    protoc --proto_path=pkg/pbuf --go_out=pkg/pbuf --go_opt=paths=source_relative \
        --go-grpc_out=pkg/pbuf --go-grpc_opt=paths=source_relative \
        pkg/pbuf/*.proto
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
