#!/bin/sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/unhidem_linux_x64 -ldflags="-s" .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/unhidem_windows_x64.exe -ldflags="-s" .
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/unhidem_windows_386.exe -ldflags="-s" .