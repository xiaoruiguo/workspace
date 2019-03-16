#! /bin/sh
#
# build.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

GOOS=linux GOARCH=amd64 CCGO_ENABLED=0 go build
