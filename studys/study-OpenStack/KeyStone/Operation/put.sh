#!/bin/bash

HOST="http://127.0.0.1:35357"
TOKEN=$2
# -s -v
OPTIONS="-v"
curl $OPTIONS -X PUT -H "Content-Type: application/json" -H "X-Auth-Token: ${TOKEN}" "$HOST/$1" | python -m json.tool
