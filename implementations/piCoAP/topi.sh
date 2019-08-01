#!/bin/bash
cd src/main/
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o main .
scp main piHome:main
rm main