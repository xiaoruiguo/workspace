#! /bin/sh
#
# run.sh
# Copyright (C) 2020 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

if [ "$#" != "1" ];then
	echo "Usage: $0 dir"
	exit
fi

cp -rf $1/*  ./00-nginx/conf.d/

docker-compose stop && docker-compose up -d
