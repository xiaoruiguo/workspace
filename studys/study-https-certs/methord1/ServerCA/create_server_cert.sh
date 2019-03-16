#!/bin/bash

SERVER_DOMAIN="test.lijiaocn.com"
ROOT_KEY=../RootCA/output/root.key
ROOT_CRT=../RootCA/output/root.crt

WORKDIR=output
if [ ! -d $WORKDIR ];then
	mkdir -p $WORKDIR
fi
SERVER_KEY=$WORKDIR/server.key
SERVER_REQ=$WORKDIR/server.config
SERVER_CSR=$WORKDIR/server.csr
SERVER_CRT=$WORKDIR/server.crt

openssl genrsa -passout pass: -out ${SERVER_KEY} 1024

cat >${SERVER_REQ}<< EOF 
[ req ]
prompt                 = yes
#default_bits           = 1024
#default_keyfile        = $SERVER_KEY
distinguished_name     = req_distinguished_name
#attributes             = req_attributes
#x509_extensions        = v3_ca

#dirstring_type = nobmp

[ req_distinguished_name ]

countryName                    = Country Name (2 letter code)
countryName_default            = CN
countryName_min                = 2
countryName_max                = 2

localityName                   = Locality Name (eg, city)
localityName_default           = BeiJing

organizationalUnitName         = Organizational Unit Name (eg, section)
organizationalUnitName_default = lijiaocn

commonName                     = Common Name (eg, YOUR name)
commonName_default             = $SERVER_DOMAIN
commonName_max                 = 64

emailAddress                   = Email Address
emailAddress_default           = admin@lijiaocn.com
emailAddress_max               = 40

#[ req_attributes ]
#challengePassword              = A challenge password
#challengePassword_min          = 4
#challengePassword_max          = 20

#[ v3_ca ]
#subjectKeyIdentifier=hash
#authorityKeyIdentifier=keyid:always,issuer:always
#basicConstraints = CA:true
EOF

#生成要被签署的文件，会提示确认每个配置项
openssl req -new -key $SERVER_KEY -out $SERVER_CSR -config $SERVER_REQ

#签署，得到签署后的证书，如果rootkey有密码，会要求输入key的密码
openssl x509 -req -days 365 -in $SERVER_CSR  -CA $ROOT_CRT -CAkey $ROOT_KEY -CAcreateserial -out $SERVER_CRT
