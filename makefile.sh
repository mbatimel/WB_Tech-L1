#!/bin/bash

for i in {00..26}; do
    directory="$i"
    mkdir -p "$directory"
    if [ -d "$directory" ]; then
        echo "Entering directory $directory"
        cd "$directory" || exit 1
        touch "main.go"
        echo "Created file in $directory"
        cd ..
    else
        echo "Directory $directory does not exist"
    fi
done
