#! /bin/sh
#
# varname.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#


prefixV1="a"
prefixV2="b"
prefixV3="c"

echo ${!prefix*}
echo ${!prefix@}
echo ${!notdef:-notexist}
