#! /bin/sh
#
# clone.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#


git clone https://github.com/hyperledger/fabric-samples.git
git clone https://github.com/hyperledger/fabric.git

#下载安装脚本，需要翻墙
curl -sSL https://goo.gl/6wtTN5 >install.sh

#下载指定版本的fabric
./bash ./install.sh -s 1.0.6

git clone https://gerrit.hyperledger.org/r/fabric-sdk-node
git clone https://gerrit.hyperledger.org/r/fabric-docs
