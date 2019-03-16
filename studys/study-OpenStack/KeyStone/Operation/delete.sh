#!/bin/bash

HOST="http://127.0.0.1:35357"
TOKEN=$2
curl -s -X DELETE -H "X-Auth-Token: ${TOKEN}" "$HOST/$1" | python -m json.tool
