#! /bin/sh
#
# start_nginx.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

if [ "$#" != "1" ];then
	echo "Usage: $0 dir"
	exit
fi

conf=echo.example.conf
rm -f /usr/local/etc/nginx/servers/$conf
ln -s `pwd`/$1/$conf /usr/local/etc/nginx/servers/$conf

brew services restart nginx
