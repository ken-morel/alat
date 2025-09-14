#!/usr/bin/fish
function build-libalat
    cd ../../pkg/libalat
    # Call the libalat build script with the output DIRECTORY
    ./mng.fish build ../../packages/dalat/src
end

switch "$argv[1]"
    case build
        switch "$argv[2]"
            case libalat
                build-libalat
            case *
                echo "No build command specified, use ./mng.fish build <target>"
        end
    case *
        echo "No command specified, use ./mng.fish <command>"
end
