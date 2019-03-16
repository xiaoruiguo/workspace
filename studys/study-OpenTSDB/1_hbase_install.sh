#!/bin/bash
sudo yum install -y java-1.8.0-openjdk

curl http://apache.fayea.com/hbase/1.2.2/hbase-1.2.2-bin.tar.gz -o ./Files/hbase-1.2.2-bin.tar.gz

tar -C ./App/  -xvf ./Files/hbase-1.2.2-bin.tar.gz

cp ./hbase-site.xml ./App/hbase-1.2.2/conf/hbase-site.xml
