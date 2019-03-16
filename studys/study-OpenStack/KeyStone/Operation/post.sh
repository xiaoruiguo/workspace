#!/bin/bash
HOST="http://127.0.0.1:35357"
TOKEN=$3
# -v
OPTIONS="-v"

curl $OPTIONS -X POST -H "Content-Type: application/json" -H "X-Auth-Token: ${TOKEN}"  -d @$2 "$HOST/$1" | python -m json.tool
