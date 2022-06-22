#/bin/bash

git clone https://github.com/VicGrygorchyk/testops_for_ta.git
cd testops_for_ta
go mod init testops_for_ta
go mod tidy
go test -v ./