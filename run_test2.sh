#/bin/bash

git clone git@github.com:VicGrygorchyk/testops_for_ta.git
cd testops_for_ta
go mod init testops_for_ta
go mod tidy
go test -v ./