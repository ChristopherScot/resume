#! /usr/bin/env bash

echo
echo "Building application..."
# GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o resume main.go
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go


echo
echo "Moving application to this directory and zipping for Lambda upload..."
mv bootstrap deploy
cd deploy
zip bootstrap.zip ./bootstrap
rm bootstrap
