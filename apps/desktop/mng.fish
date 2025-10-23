#!/usr/bin/env fish

function dev
    wails dev -tags webkit2_41 -v 2
end
function build-windows
    set -gx CGO_ENABLED 1
    set -gx GO111MODULE on
    wails build -v 2 -tags webkit2_41 -platform windows -nsis
end
function package-linux
    nfpm pkg --packager deb --target ./build/bin
end

function build-linux
    wails build -v 2 -tags webkit2_41
end

function build
    switch "$argv[1]"
        case windows
            build-windows
        case linux
            build-linux
        case ""
            echo "No build target specified, either `build linux` or `build windows`"
        case "*"
            echo "Invalid build target, either `build linux` or `build-windows`"
    end
end

function help
    echo "Usage: ./mng.fish <command>"
    echo ""
    echo "Commands:"
    echo "  dev              - Start the development server."
    echo "  build-windows    - Builds windows binaries."
    echo "  package          - Builds debian archive"
    echo "  build            - Builds linux binaries."
    echo "  help             - Shows this help message."
end

switch "$argv[1]"
    case help
        help
    case dev
        dev
    case build
        build "$argv[2]"
    case package
        package-linux
    case ""
        echo "Error: No command specified."
        help
        exit 1
    case '*'
        echo "Error: Unknown command '$argv[1]'"
        help
        exit 1
end
