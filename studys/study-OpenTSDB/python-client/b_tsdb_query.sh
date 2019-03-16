#!/bin/bash
set -x
curl \
-H "Content-Type: application/json" \
-X GET \
'127.0.0.1:4242/api/query?start=5h-ago&m=sum:floatingip.incoming.packets\{resource_id=5ebe4635-9a4b-4870-8eb3-a192ce54b869\}&showQuery=true' | python -m json.tool

set -x
curl \
-H "Content-Type: application/json" \
-X GET \
'127.0.0.1:4242/api/query?start=5h-ago&m=sum:floatingip.incoming.packets&showQuery=true' | python -m json.tool

