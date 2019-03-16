#! /bin/sh
#
# create_api.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

#go get -u github.com/golang/dep/cmd/dep
#brew install dep
#brew upgrade dep
NAME="test.io/hello"

bee api $NAME
cd $GOPATH/src/$NAME

dep ensure -add  github.com/astaxie/beego
#bee generate view user
#bee generate model user -fields="name:string,age:int"
bee generate controller health
