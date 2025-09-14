#!/usr/bin/fish
function build -a out_dir
    if test -z "$out_dir"
        echo "Error: No output directory specified for build."
        return 1
    end

    echo "Building libalat.so shared library and header file in $out_dir"
    # cgo will create both libalat.so and libalat.h in the same directory
    go build -buildmode=c-shared -o "$out_dir/libalat.so" .
end

switch "$argv[1]"
    case build
        build "$argv[2]"
    case "*"
        echo "No command specified. Usage: ./mng.fish build <output_directory>"
end
