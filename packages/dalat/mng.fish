#!/usr/bin/fish

function build-libalat -a platform
    cd ../../lib/libalat
    if test "$platform" = android
        echo "Building libalat for android"
        ./mng.fish build-android ../../packages/dalat
    else
        echo "Building libalat for current system"
        ./mng.fish build ../../packages/dalat/src
    end
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
        build-libalat "$argv[2]"
    case gen
        gen-ffi
    case gen-json
        gen-json
    case make
        build-libalat "$argv[2]"
        gen-ffi
        gen-json
    case "*"
        echo "No command specified, use ./mng.fish <command>"
end
