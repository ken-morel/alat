function call_flutter
    flutter $argv[1..]
end
function flutter_dev
    flutter run
end

switch "$argv[1]"
    case flutter
        call_flutter $argv[2..]
    case dev
        flutter_dev
    case "*"
        echo "No or invalid command, specify flutter|.."
end
