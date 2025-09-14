#!/usr/bin/fish
function build-libalat
    cd ../../pkg/libalat
    # Call the libalat build script with the output DIRECTORY
    ./mng.fish build ../../packages/dalat/src
end
function build-dalat
    dart run ffigen --config ffigen.yaml
end

switch "$argv[1]"
    case build
        switch "$argv[2]"
            case libalat
                build-libalat
            case dalat
                build-dalat
            case "*"
                echo "Invalid build command specified, use specify target: dalat, libalat, or leave empty"
        end
    case "*"
        echo "No command specified, use ./mng.fish <command>"
end
