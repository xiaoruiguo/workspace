#!/bin/bash

VIRPATH=./virtualenv

if [ ! -d $VIRPATH ];then
	virtualenv $VIRPATH --no-site-packages
fi
#bash  --rcfile  ${VIRPATH}/bin/activate

