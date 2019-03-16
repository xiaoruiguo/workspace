#!/bin/bash
export JAVA_HOME=/usr
env COMPRESSION=NONE HBASE_HOME=./App/hbase-1.2.2 /usr/share/opentsdb/tools/create_table.sh
