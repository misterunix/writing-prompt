#!/usr/bin/bash
archs=(amd64 arm64)

#for arch in ${archs[@]}
#do
#	env GOOS=linux GOARCH=${arch} go build -o bin/writing-prompt-linux-debug-${arch}
#done


for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -ldflags="-w -s" -o bin/writing-prompt-linux-release-${arch}
done

