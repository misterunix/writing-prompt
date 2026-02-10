#!/usr/bin/bash

rm bin/writing-prompt*

env GOOS=linux GOARCH=amd64 go  build -o bin/writing-prompt-linux-amd64 -ldflags="-w -s"
env GOOS=linux GOARCH=arm64 go  build -o bin/writing-prompt-linux-arm64 -ldflags="-w -s"
env GOOS=darwin GOARCH=amd64 go  build -o bin/writing-prompt-darwin-amd64 -ldflags="-w -s"
env GOOS=darwin GOARCH=arm64 go  build -o bin/writing-prompt-darwin-arm64 -ldflags="-w -s"
env GOOS=windows GOARCH=amd64 go  build -o bin/writing-prompt-windows-amd64.exe -ldflags="-w -s"

cp bin/writing-prompt-linux-amd64 ~/bin/writing-prompt-linux-amd64


