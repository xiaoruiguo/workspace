#!/bin/bash

HOST="http://127.0.0.1:35357"
TOKEN=$3
curl -s -X PATCH -H "Content-Type: application/json" -H "X-Auth-Token: ${TOKEN}"  -d @$2 "$HOST/$1" | python -m json.tool
