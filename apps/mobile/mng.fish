function call_flutter
    flutter $argv[1..]
end
function flutter_dev
    switch "$argv[1]"
        case linux
            exec flutter run -d linux -v
        case "*"
            exec flutter run -d
    end
end

switch "$argv[1]"
    case flutter
        call_flutter $argv[2..]
    case dev
        flutter_dev $argv[2..]
    case build
        exec flutter build "$argv[2..]"
    case run
        exec flutter run "$argv[2..]"
    case arb-watch
        exec arb-util
    case "*"
        echo "No or invalid command, specify flutter|.."
end
