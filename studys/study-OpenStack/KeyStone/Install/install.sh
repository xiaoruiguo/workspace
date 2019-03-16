#!/bin/bash
REPO="https://github.com/openstack/keystone"
TAG="8.0.0"
RUNPATH=`pwd`
LOCALEDIR=${RUNPATH}/../Source/keystone
###############################################################
if [ ! -e ./library.sh ];then
	wget https://raw.githubusercontent.com/lijiaocn/LinuxShell/master/library.sh
fi
source ./library.sh

func_git_check_tag $REPO master $TAG $LOCALEDIR

cd $LOCALEDIR; python setup.py install --record ${RUNPATH}/install.log; cd $RUNPATH
