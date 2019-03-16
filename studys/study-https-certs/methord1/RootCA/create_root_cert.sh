#!/bin/sh

#制作自签署的根证书
WORKDIR=output
if [ ! -d $WORKDIR ];then
	mkdir -p $WORKDIR
fi
ROOT_KEY=$WORKDIR/root.key        #即是用来签署证书的key，也是证书的key
ROOT_REQ=$WORKDIR/root.config     #签署请求
ROOT_CSR=$WORKDIR/root.csr        #根据签署请求生成的等待被签署的文件
ROOT_CRT=$WORKDIR/root.crt        #签署后的证书

#生成ROOT_KEY，生成的时候，会要求设置密码，以后用这个key进行签署的时候需要输入设置的密码
#`-passout pass:`表示密码为空，冒号后面是密码
openssl genrsa -passout pass: -out ${ROOT_KEY} 1024

#准备签署请求
cat >${ROOT_REQ}<< EOF 
[ req ]
prompt                 = yes
default_bits           = 2048
default_keyfile        = $ROOT_KEY
distinguished_name     = req_distinguished_name
attributes             = req_attributes
x509_extensions        = v3_ca

dirstring_type = nobmp

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
commonName_default             = lijiaocn
commonName_max                 = 64

emailAddress                   = Email Address
emailAddress_default           = admin@lijiaocn.com
emailAddress_max               = 40

[ req_attributes ]
challengePassword              = A challenge password
challengePassword_min          = 4
challengePassword_max          = 20

[ v3_ca ]
subjectKeyIdentifier=hash
authorityKeyIdentifier=keyid:always,issuer:always
basicConstraints = CA:true
EOF

#生成要被签署的文件，会提示确认每个配置项
openssl req -new -key $ROOT_KEY -out $ROOT_CSR -config $ROOT_REQ

#签署，得到签署后的证书，会要求输入key的密码
openssl x509 -req -days 365 -in $ROOT_CSR -signkey $ROOT_KEY -out $ROOT_CRT

#查看签署后证书
openssl x509 -in $ROOT_CRT -noout -text
