#!/bin/sh
func_create_dirs(){
	for i in $*
	do
		if [ ! -d $i ];then
			mkdir -p $i
		fi
	done
}

#制作自签署的根证书
WORKDIR=`pwd`/output
CERTS=$WORKDIR/certs
CRL=$WORKDIR/crl
CSR=$WORKDIR/csr
NEWCERTS=$WORKDIR/newcerts
PRIVATE=$WORKDIR/private
INDEXFILE=$WORKDIR/index.txt
SERIALFILE=$WORKDIR/serial
CRLNUMBER=$WORKDIR/crlnumber

func_create_dirs $CERTS $CRL $CSR $NEWCERTS $PRIVATE
touch $INDEXFILE
echo 1000 > $SERIALFILE
echo 1000 > $CRLNUMBER

KEY=$PRIVATE/media.ca.key.pem
CERT=$CERTS/media.ca.cert.pem
CSRPEM=$CSR/media.ca.csr.pem
CACHAIN=$CERTS/media.ca-chain.cert.pem
CRLPEM=$CRL/media-ca.crl.pem
CONF=./openssl.cnf
CACONF=../RootCA/openssl.cnf
CACERT=../RootCA/output/certs/ca.cert.pem

cat >$CONF <<EOF
# OpenSSL root CA configuration file.
# Copy to /root/ca/openssl.cnf.

[ ca ]
# man ca
default_ca = CA_default

[ CA_default ]
# Directory and file locations.
dir               = $WORKDIR
certs             = \$dir/certs
crl_dir           = \$dir/crl
new_certs_dir     = \$dir/newcerts
database          = \$dir/index.txt
serial            = \$dir/serial
RANDFILE          = \$dir/private/.rand

# The root key and root certificate.
private_key       = \$dir/private/media.ca.key.pem
certificate       = \$dir/certs/media.ca.cert.pem

# For certificate revocation lists.
crlnumber         = \$dir/crlnumber
crl               = \$dir/crl/media-ca.crl.pem
crl_extensions    = crl_ext
default_crl_days  = 30

# SHA-1 is deprecated, so use SHA-2 instead.
default_md        = sha256

name_opt          = ca_default
cert_opt          = ca_default
default_days      = 375
preserve          = no
policy            = policy_loose

[ policy_strict ]
# The root CA should only sign intermediate certificates that match.
# See the POLICY FORMAT section of man ca.
countryName             = match
stateOrProvinceName     = match
organizationName        = match
organizationalUnitName  = optional
commonName              = supplied
emailAddress            = optional

[ policy_loose ]
# Allow the intermediate CA to sign a more diverse range of certificates.
# See the POLICY FORMAT section of the ca man page.
countryName             = optional
stateOrProvinceName     = optional
localityName            = optional
organizationName        = optional
organizationalUnitName  = optional
commonName              = supplied
emailAddress            = optional

[ req ]
# Options for the req tool (man req).
default_bits        = 2048
distinguished_name  = req_distinguished_name
string_mask         = utf8only

# SHA-1 is deprecated, so use SHA-2 instead.
default_md          = sha256

# Extension to add when the -x509 option is used.
x509_extensions     = v3_ca

[ req_distinguished_name ]
# See <https://en.wikipedia.org/wiki/Certificate_signing_request>.
countryName                     = Country Name (2 letter code)
countryName_default             = CN

stateOrProvinceName             = State or Province Name
stateOrProvinceName_default     = BeiJing

localityName                    = Locality Name
localityName_default            = BeiJing

0.organizationName              = Organization Name
0.organizationName_default      = AAlijiaocn

organizationalUnitName          = Organizational Unit Name
organizationalUnitName_default  = AAlijiaocn

commonName                      = Common Name
commonName_default              = media.ca.lijiaocn
commonName_max                  = 64

emailAddress                    = Email Address
emailAddress_default            = admin@lijiaocn.com
emailAddress_max                = 40

[ v3_ca ]
# Extensions for a typical CA (man x509v3_config).
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always,issuer
basicConstraints = critical, CA:true
keyUsage = critical, digitalSignature, cRLSign, keyCertSign

[ v3_intermediate_ca ]
# Extensions for a typical intermediate CA (man x509v3_config).
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always,issuer
basicConstraints = critical, CA:true, pathlen:0
keyUsage = critical, digitalSignature, cRLSign, keyCertSign

[ usr_cert ]
# Extensions for client certificates (man x509v3_config).
basicConstraints = CA:FALSE
nsCertType = client, email
nsComment = "OpenSSL Generated Client Certificate"
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer
keyUsage = critical, nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, emailProtection

[ server_cert ]
# Extensions for server certificates (man x509v3_config).
basicConstraints = CA:FALSE
nsCertType = server
nsComment = "OpenSSL Generated Server Certificate"
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer:always
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth

[ crl_ext ]
# Extension for CRLs (man x509v3_config).
authorityKeyIdentifier=keyid:always

[ ocsp ]
# Extension for OCSP signing certificates (man ocsp).
basicConstraints = CA:FALSE
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer
keyUsage = critical, digitalSignature
extendedKeyUsage = critical, OCSPSigning
EOF

openssl genrsa -passout pass: -out $KEY 4096
openssl req -config $CONF -new -sha256 -key $KEY -out $CSRPEM
openssl ca -config $CACONF -extensions v3_intermediate_ca  -days 3650 -notext -md sha256 -in $CSRPEM -out $CERT
cat $CERT $CACERT>$CACHAIN
openssl ca -config $CONF -gencrl -out $CRLPEM
