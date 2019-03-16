#!/bin/bash

go tool 6g -S main.go  1>6g-S.out 2>&1
go build -gcflags -S main.go 1>gc-S.out 2>&1
