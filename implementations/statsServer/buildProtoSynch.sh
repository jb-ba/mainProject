#!/bin/bash

protoc -I synchProto/ synchProto/statsSynch.proto --go_out=plugins=grpc:synchProto