#!/bin/bash

#CompileDaemon -command="./go-hexa-archi"

# build
go build -o ./tmp/main ./cmd/http/main.go

# run
./tmp/main
