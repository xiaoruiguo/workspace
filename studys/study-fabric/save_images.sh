#! /bin/sh
#
# save.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

mkdir hyperledger
docker images|grep hyper |grep -v latest |grep -v snapshot |awk '{print "docker save " $1 ":" $2 " | gzip >" $1".tar.gz"}'  |bash
