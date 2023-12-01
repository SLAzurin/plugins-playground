#!/bin/bash
set -e
cd plugins || exit 1
mkdir -p bin

plugins=()
while IFS='' read -r -d $'\0'; do
    plugins+=("$REPLY")
done < <(find . -type d -not -name 'bin' -not -name '.' -print0)
pids=()
for p in "${plugins[@]}"; do
    if [ "$p" = './bin/' ]; then continue; fi
    {
        cd "$p"
        go build -buildmode=plugin -o "../bin/$(basename "$p")" ./*.go
    } &
    pids+=("$!")
done

for pid in "${pids[@]}"; do
    wait "$pid"
done