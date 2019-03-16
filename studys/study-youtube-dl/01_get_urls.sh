#! /bin/sh
#
# 01_get_url.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

curl -H "User-Agent: Mozilla/5.0"  $1 >/tmp/1.html

for i in `cat /tmp/1.html |grep "var ytplayer" | tr '"' '\n' |sed 's|\\u0026|\'$'\n|g'  |grep "^url="`
do
	echo $i
done

