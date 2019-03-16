#!/bin/bash
output=`pwd`/1_output
cmd="java -jar ../swagger-codegen/modules/swagger-codegen-cli/target/swagger-codegen-cli.jar "

client=$output/client
if [ ! -d  ${client} ];then
	mkdir -p ${client}
fi
$cmd generate -i ./swagger.yaml -l go -c ./config.json -o ${client}

server=$output/server
if [ ! -d  ${server} ];then
	mkdir -p ${server}
fi

$cmd generate -i ./swagger.yaml -l go-server -c ./config.json -o ${server}
