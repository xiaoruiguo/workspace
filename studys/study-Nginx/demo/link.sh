#! /bin/sh
#
# baidu_link.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

conf="echo.example.conf"
ln -s `pwd`/$conf /usr/local/etc/nginx/servers/$conf
