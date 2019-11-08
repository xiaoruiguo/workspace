#! /bin/sh
#
# signal.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

function debug {
   echo "DEBUG"
}
trap debug DEBUG

function err {
   echo "ERR"
}
trap err ERR

function cleanup {
   echo "EXIT"
}
trap cleanup EXIT

echo "exec command..."
ls /noexit
