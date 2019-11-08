#! /bin/sh
#
# curl.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

cmd="curl 127.0.0.1:6699$*"
echo "\$ $cmd"
$cmd
