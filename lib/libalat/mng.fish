#!/usr/bin/fish
function build -a out_dir
    if test -z "$out_dir"
        echo "Error: No output directory specified for build."
        return 1
    end

    echo "Building libalat.so shared library and header file in $out_dir"
    # cgo will create both libalat.so and libalat.h in the same directory
    go build -buildmode=c-shared -o "$out_dir/libalat.so" .
    cp ./*.h "$out_dir/"
end

function build-android -a out_dir
    if test -z "$out_dir"
        echo "Error: No output directory specified for build-android."
        return 1
    end

    set -l ndk_path $ANDROID_NDK_HOME
    if test -z "$ndk_path"
        echo "Error: ANDROID_NDK_HOME is not set."
        return 1
    end

    set -l toolchain "$ndk_path/toolchains/llvm/prebuilt/linux-x86_64"
    set -l api 24

    set -l targets armv7a-linux-androideabi aarch64-linux-android i686-linux-android x86_64-linux-android
    set -l go_archs arm arm64 386 amd64
    set -l jni_archs armeabi-v7a arm64-v8a x86 x86_64

    echo "BUilding libalat for $(count $targets) architectures on api $api with ndk at $ndk_path"

    for i in (seq (count $targets))
        set -l target $targets[$i]
        set -l go_arch $go_archs[$i]
        set -l jni_arch $jni_archs[$i]

        set -l out_path "$out_dir/android/src/main/jniLibs/$jni_arch"
        mkdir -p "$out_path"

        echo "Building libalat for $target \($go_arch, $jni_arch\) -> $out_path"

        set -x CC "$toolchain/bin/$target$api-clang"
        set -x CXX "$toolchain/bin/$target$api-clang++"
        set -x GOOS android
        set -x GOARCH $go_arch
        set -x CGO_ENABLED 1
        set -x CGO_CFLAGS -fPIC

        go build -buildmode=c-shared -o "$out_path/libalat.so" .; and echo "Done !"; and continue
        echo "Error building!"
    end
end

switch "$argv[1]"
    case build
        build "$argv[2]"
    case build-android
        build-android "$argv[2]"
    case "*"
        echo "No command specified. Usage: ./mng.fish build <output_directory>"
end
