#!/bin/bash
output=$GOPATH/src/znr

client=$output/client
if [ ! -d  ${client} ];then
	mkdir -p ${client}
fi
./cli.sh generate -i ./template.yaml -l go -c ./config.json -o ${client} 
#./cli.sh generate -i ./swagger.json -l go -c ./config.json -o ${client} 

server=$output/server
if [ ! -d  ${server} ];then
	mkdir -p ${server}
fi
./cli.sh generate -i ./template.yaml -l go-server -c ./config.json -o ${server}
#./cli.sh generate -i ./swagger.json -l go-server -c ./config.json -o ${server}
