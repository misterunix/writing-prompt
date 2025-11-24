#!/usr/bin/bash
archs=(amd64 arm64)

#for arch in ${archs[@]}
#do
#	env GOOS=linux GOARCH=${arch} go build -o bin/writing-prompt-linux-debug-${arch}
#done


for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -ldflags="-w -s" -o bin/writing-prompt-linux-${arch}
done

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -ldflags="-w -s" -o bin/writing-prompt-windows-${arch}.exe
done

