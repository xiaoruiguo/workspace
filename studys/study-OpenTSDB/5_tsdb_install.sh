#!/bin/bash
curl -L https://github.com/OpenTSDB/opentsdb/releases/download/v2.2.0/opentsdb-2.2.0.noarch.rpm -o ./Files/opentsdb-2.2.0.noarch.rpm

sudo yum install ./Files/opentsdb-2.2.0.noarch.rpm
