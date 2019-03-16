#!/bin/bash

openssl ca -config ../MediaCA/openssl.cnf  -revoke ../ServerCA/output/certs/server1.com.cert.pem
openssl ca -config ../MediaCA/openssl.cnf -gencrl -out ../MediaCA/output/crl/media-ca.crl.pem
