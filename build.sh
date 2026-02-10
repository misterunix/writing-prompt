#!/usr/bin/bash

rm bin/writing-prompt*

env GOOS=linux GOARCH=amd64 go  build -o bin/writing-prompt -ldflags="-w -s"
cp bin/writing-prompt ~/bin/writing-prompt


