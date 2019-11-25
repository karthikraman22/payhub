#!/bin/sh

os_archs=()

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
for goos in darwin linux 
do
    for goarch in amd64 
    do
        GOOS=${goos} GOARCH=${goarch} go build -o payhub-${goos}-${goarch}
        if [ $? -eq 0 ]
        then
            os_archs+=("${goos}/${goarch}")
        fi
    done
done

for os_arch in "${os_archs[@]}"
do
    echo ${os_arch}
done

