#!/bin/bash
REPO="https://github.com/openstack/neutron.git"
TAG="8.0.0"
RUNPATH=`pwd`
LOCALEDIR=${RUNPATH}/../Source/neutron
CONF=${RUNPATH}/../Conf
###############################################################
if [ ! -e ./library.sh ];then
	wget https://raw.githubusercontent.com/lijiaocn/LinuxShell/master/library.sh
fi
source ./library.sh

func_git_check_tag $REPO master $TAG $LOCALEDIR

cd $LOCALEDIR; pip install -r requirements.txt && python setup.py install --record ${RUNPATH}/install.log; cd $RUNPATH

cd $LOCALEDIR
	for i in `ls etc/oslo-config-generator/`
	do
		echo "generate config:  $i"
		oslo-config-generator --config-file  etc/oslo-config-generator/${i}
	done
cd $RUNPATH

if  [ ! -d $CONF ];then
	mkdir -p $CONF
	cp -rf $LOCALEDIR/etc/*   $CONF/
	rm -rf $CONF/oslo-config-generator
fi
