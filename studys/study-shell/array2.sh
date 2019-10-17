#! /bin/sh
#
# array.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

array=("a"
"b"
"c"
)

for i in "${array[@]}";do
	echo $i
done
