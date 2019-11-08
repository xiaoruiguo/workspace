#! /bin/sh
#
# variable.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#


echo "未定义的变量，使用默认值：${notdef1:-word}  $notdef1"
echo "未定义的变量，赋值默认值：${notdef2:=word}  $notdef2"
echo "已定义的变量，使用替代数值：${notdef2:+valueinstead}"
echo "未定义的变量，打印错误信息：${notdef3:?notdef}"
