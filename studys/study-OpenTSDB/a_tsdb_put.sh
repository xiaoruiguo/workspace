#!/bin/bash
set -x
timestamp=`date +"%s"`
curl \
-H "Content-Type: application/json" \
-X POST \
-d \
"{
    \"metric\": \"memory.usage\",
    \"timestamp\": ${timestamp},
    \"value\": 20,
    \"tags\": {
       \"resource_id\": \"2bd13e63-4bc9-4739-a28f-21b6ec166d72\"
    }
}"      \
127.0.0.1:4242/api/put?details | python -m json.tool
