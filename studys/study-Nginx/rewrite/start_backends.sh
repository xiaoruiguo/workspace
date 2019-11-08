#! /bin/sh
#
# start_echoserver.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

docker rm -f echoserver
docker run -idt --name echoserver -p 9090:8080 googlecontainer/echoserver:1.10

docker rm -f http-record
docker run -idt --name http-record -p 9091:8080 lijiaocn/http-record:0.0.1
