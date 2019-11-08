#! /bin/sh
#
# start_echoserver.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

docker rm -f echoserver
docker run -idt --name echoserver -p 9090:8080 -p 9443:8443 googlecontainer/echoserver:1.10
