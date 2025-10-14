#!/usr/bin/fish

function build-libalat
    echo "Building libalat"
    cd ../../lib/libalat
    ./mng.fish build ../../packages/dalat/src
    cd ../../packages/dalat/
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
    case gen-json
        gen-json
    case make
        build-libalat
        gen-ffi
        gen-json
    case "*"
        echo "No command specified, use ./mng.fish <command>"
end
