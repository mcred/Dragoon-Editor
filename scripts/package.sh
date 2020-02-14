#!/usr/bin/env bash

set -ex

export PATH=$PATH:$(go env GOPATH)/bin
fyne package -executable ./build/LODeditor -os darwin -icon ./assets/icon.png