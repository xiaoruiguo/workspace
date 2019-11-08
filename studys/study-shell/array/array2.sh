#! /bin/sh
#
# array.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

array=("value0" "value1" "value2" "value3")

echo ${array[1]}

for i in "${array[@]}";do
	echo "value: $i"
done

# 打印数组中有值的 index
for i in ${!array[@]}
do
    echo "index: $i"
done

# 打印数组内元素个数
echo "array size: ${#array[@]}"
