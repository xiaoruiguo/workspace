#!/bin/bash
jar=`pwd`/swagger-codegen/modules/swagger-codegen-cli/target/swagger-codegen-cli.jar
java -jar $jar  $*
