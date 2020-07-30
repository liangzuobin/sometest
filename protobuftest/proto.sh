#!/bin/bash

protoc -I . ./demo/demo.proto --go_out=plugins=grpc:.