#!/bin/bash

protoc --proto_path=pb --go_out=plugins=grpc:pb pb/logRPC.proto

