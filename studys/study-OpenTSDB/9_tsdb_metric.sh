#!/bin/bash
curl \
-H "Content-Type: application/json" \
-X POST \
-d \
'{
    "metric": [
        "memory.usage",
	"floatingip.incoming.packets"
    ]
}' \
127.0.0.1:4242/api/uid/assign | python -m json.tool
