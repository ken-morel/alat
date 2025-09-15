#!/usr/bin/fish
function build-libalat
    echo "Building libalat"
    cd ../../pkg/libalat
    # Call the libalat build script with the output DIRECTORY
    ./mng.fish build ../../packages/dalat/src
end
function gen-ffi
    echo "Generating ffi bindings for dalat"
    dart run ffigen --config ffigen.yaml
end
function gen-json
    echo "Generating json serializable"
    dart run build_runner build
end

switch "$argv[1]"
    case build
        build-libalat
    case gen
        gen-ffi
    case make
        build-libalat
        cd ../../packages/dalat/
        gen-ffi
        gen-json
    case "*"
        echo "No command specified, use ./mng.fish <command>"
end
