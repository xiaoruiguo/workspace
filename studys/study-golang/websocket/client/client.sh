#! /bin/sh
#
# client.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#


#npm config set registry http://registry.npm.taobao.org/
#npm install -g wscat  --registry https://registry.npm.taobao.org 

wscat -c ws://127.0.0.1:8080/echo
