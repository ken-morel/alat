#!/usr/bin/env fish

# Get the save path from the user
set file_path (zenity --file-selection --save --confirm-overwrite --title="Save Audio Recording")

# Check if a file path was provided
if test -z "$file_path"
    notify-send -t 2000 "Recording cancelled" "No file selected."
    exit 1
end

# Get the list of sources
set sources (pactl list sources | grep -E "Name:|Description:" | sed 's/^\s*Name: //g' | sed 's/^\s*Description: //g' | paste -d' ' - -)

# Check if any sources were found
if test -z "$sources"
    notify-send -t 2000 "Recording failed" "No audio sources found."
    exit 1
end

# Let the user choose a source with rofi
set chosen_source_line (echo -e "$sources" | rofi -dmenu -i -p "Select audio source")

# Check if a source was chosen
if test -z "$chosen_source_line"
    notify-send -t 2000 "Recording cancelled" "No audio source selected."
    exit 1
end

# Extract the source name from the chosen line
set source_name (echo "$chosen_source_line" | awk '{print $1}')

# Countdown from 3
for i in (seq 3 -1 1)
    notify-send -t 1000 "Recording starts in $i..."
    sleep 1
end

# Start recording
notify-send "Recording started!" "Saving to $file_path"
ffmpeg -f pulse -i "$source_name" -b:a 320k "$file_path"
