#/bin/bash

git checkout add_bash
go mod tidy
go test -v ./
