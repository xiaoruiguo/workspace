#!/bin/bash
HOST="http://127.0.0.1:35357"
TOKEN=$3
curl  -v -s -X GET  -H "X-Auth-Token: ${TOKEN}" -H "X-Subject-Token: $2" "$HOST/$1" | python -m json.tool
