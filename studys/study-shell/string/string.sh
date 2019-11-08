#! /bin/sh
#
# string.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#


str="abcddeddfghijka"

echo "第0个字符后: ${str:0}"
echo "第1个字符后: ${str:1}"
echo "第1个字符后的两个字符: ${str:1:2}"

echo "字符串长度：${#str}"

echo "head match #abc: ${str#abc}"
echo "head match ##abc: ${str##abc}"

echo "head match #abc*: ${str#abc*}"
echo "head match ##abc*: ${str##abc*}"

echo "tail match %ijk: ${str%ijk}"
echo "tail match %%ijk: ${str%%ijk}"

echo "tail match %*ijk: ${str%*ijk}"
echo "tail match %%*ijk: ${str%%*ijk}"

echo "replace /d/D（替换一个）: ${str/d/D}"
echo "replace //d/D（全部替换）: ${str//d/D}"
echo "replace /#a/A（从开始处匹配替换）: ${str/#a/A}"
echo "replace /%a/A（从结尾处匹配替换）: ${str/%a/A}"
