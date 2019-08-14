#!/bin/bash

protoc -I src/syncProto/ src/syncProto/statsSync.proto --go_out=plugins=grpc:src/syncProto