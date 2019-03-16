#!/bin/bash
set -x
curl \
-H "Content-Type: application/json" \
-X GET \
'127.0.0.1:4242/api/query?start=1m-ago&m=sum:memory.usage\{resource_id=2bd13e63-4bc9-4739-a28f-21b6ec166d72\}&showQuery=true' | python -m json.tool

set -x
curl \
-H "Content-Type: application/json" \
-X GET \
'127.0.0.1:4242/api/query?start=1m-ago&m=sum:2s-sum-nan:memory.usage\{resource_id=2bd13e63-4bc9-4739-a28f-21b6ec166d72\}&showQuery=true' | python -m json.tool
