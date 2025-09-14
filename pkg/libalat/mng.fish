function build -a out
    if test "$out" = ""
        echo "No output directory specified for build"
    end
    echo "Building libalat.so shared library and header file"
    go build -buildmode=c-shared -o $out/libalat.so
end

switch "$argv[1]"
    case build
        build "$argv[2]"
    case *
        echo "No command specified"
end
