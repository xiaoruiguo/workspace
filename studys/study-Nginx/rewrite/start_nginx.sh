#! /bin/sh
#
# start_nginx.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

conf="echo.example.conf"
rm -f /usr/local/etc/nginx/servers/$conf
ln -s `pwd`/$conf /usr/local/etc/nginx/servers/$conf

brew services restart nginx
