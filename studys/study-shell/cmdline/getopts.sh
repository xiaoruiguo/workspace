#! /bin/sh
#
# getopts.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

while getopts "p:h:abc" option ;do 
	if [[ $option == "p" ]];then
		echo "listener port: $OPTARG"
	fi
	if [[ $option == "abc" ]];then
		echo "abc enabled"
	fi
	if [[ $option == "h" ]];then
		read name value <<< "${OPTARG//:/ }";
		headers="$headers\nproxy_set_header $name \"$value\";"
	fi
done

echo $headers
