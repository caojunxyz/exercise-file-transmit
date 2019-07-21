#!/bin/sh
protoc -I . deliver.proto --go_out=plugins=grpc:.