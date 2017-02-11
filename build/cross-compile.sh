#!/bin/bash
set -e

gox -output "./pkg/{{.OS}}-{{.Arch}}/terraform-provider-sumologic" -os="linux darwin windows freebsd openbsd solaris" -arch="amd64 386 arm" -osarch="!darwin/arm !darwin/386"

for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    OSARCH=$(basename ${PLATFORM})
    echo "--> ${OSARCH}"

    pushd $PLATFORM >/dev/null 2>&1
    zip ../${OSARCH}.zip ./*
    popd >/dev/null 2>&1
    rm -rf $PLATFORM
done
