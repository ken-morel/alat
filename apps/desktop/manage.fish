#!/usr/bin/env fish

function dev
    wails dev -tags webkit2_41 -v 2
end
function build-windows
    wails build -v 2 -tags webkit2_41 -platform windows -nsis
end

function help
    echo "Usage: ./manage.fish <command>"
    echo ""
    echo "Commands:"
    echo "  dev              - Start the development server."
    echo "  build-windows    - Builds windows binaries."
    echo "  help             - Shows this help message."
end

switch "$argv[1]"
    case help
        help
    case dev
        dev
    case build-windows
        build-windows
    case ""
        echo "Error: No command specified."
        help
        exit 1
    case '*'
        echo "Error: Unknown command '$argv[1]'"
        help
        exit 1
end
