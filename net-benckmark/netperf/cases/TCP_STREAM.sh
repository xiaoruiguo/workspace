#!/bin/bash
source ./config
TYPE=`basename $0 .sh`
set -x
netperf -${IPVERION} -H $HOST  -p $PORT  -t ${TYPE} $*  
