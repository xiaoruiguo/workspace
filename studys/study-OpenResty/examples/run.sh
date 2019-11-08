#! /bin/sh
#
# start.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

if [[ $# < 1 ]];then
	echo "usage: $0 direcotry"
	exit 1
fi

./stop.sh

DIR=$1

if [[ ! -d $DIR/logs ]];then
	mkdir $DIR/logs
fi

pushd $DIR
	openresty -p `pwd` -c nginx.conf
popd
