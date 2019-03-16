#!/bin/bash
source ./config
TYPE=`basename $0 .sh`
netperf -${IPVERION} -H $HOST  -p $PORT  -t ${TYPE} $*  
