#!/bin/bash
WORKPATH=`pwd`
OUTPUT=${WORKPATH}/out
INSTALL=/usr/local/

if [ ! -d ${OUTPUT} ];then
	mkdir ${OUTPUT}
fi

if [ ! -d ${INSTALL} ];then
	mkdir ${INSTALL}
fi

tar -C ${OUTPUT} -xvf ./netperf-2.7.0.tar.gz
cd ${OUTPUT}/netperf-2.7.0
./configure --prefix=${INSTALL}
sudo make install
