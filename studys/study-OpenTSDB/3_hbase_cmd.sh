#!/bin/bash
export JAVA_HOME=/usr
set -x
./App/hbase-1.2.2/bin/hbase $*
