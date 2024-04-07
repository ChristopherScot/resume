#! /usr/bin/env bash
set -e

echo
echo "Building application..."
if [ "$1" = "amd64" ]; then
    echo "Building for Linux environment..."
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
elif [ "$1" = "arm64" ]; then
    echo "Building for ARM64 environment..."
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go
else
    echo "Building for local environment..."
    go build -o bootstrap main.go
fi
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go


echo
echo "Moving application to this directory and zipping for Lambda upload..."
mv bootstrap deploy
cd deploy
zip bootstrap.zip ./bootstrap
rm bootstrap
