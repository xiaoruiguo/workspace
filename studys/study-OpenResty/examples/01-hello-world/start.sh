#! /bin/sh
#
# start.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#


openresty -p `pwd` -c nginx.conf
