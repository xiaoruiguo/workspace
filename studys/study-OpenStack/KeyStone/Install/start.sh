#!/bin/bash
RUNPATH=`pwd`
DATAPATH=${RUNPATH}/../Data
DBFILE=${DATAPATH}/keystone.db
CONFPATH=${RUNPATH}/../Conf
###############################################################
if [ ! -e ./library.sh ];then
	wget https://raw.githubusercontent.com/lijiaocn/LinuxShell/master/library.sh
fi
source ./library.sh

func_create_dirs $DATAPATH $CONFPATH

sed -i -e "s#{SQLITE_FILE_PATH}#${DBFILE}#" ${CONFPATH}/keystone.conf

if [ ! -e /vagrant/Data/keystone.db ]
then
	keystone-manage --config-dir $CONFPATH db_sync
fi

keystone-all --config-dir $CONFPATH  1>${DATAPATH}/keystone.log  2>&1  &
