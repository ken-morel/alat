#!/usr/bin/fish

function build-libalat -a platform
    echo "Building libalat"
    cd ../../lib/libalat
    if test "$platform" = "android"
        ./mng.fish build-android ../../packages/dalat
    else
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
