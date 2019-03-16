#!/bin/bash
openssl genrsa -aes256 -out ca1-key.pem 2048
openssl req -new -x509 -days 365 -key ca1-key.pem -sha256 -out ca1.pem

openssl genrsa -aes256 -out ca2-key.pem 2048
openssl req -new -x509 -days 365 -key ca2-key.pem -sha256 -out ca2.pem

openssl genrsa -aes256 -out ca3-key.pem 2048
openssl req -new -x509 -days 365 -key ca3-key.pem -sha256 -out ca3.pem

cat ca1.pem  >cas.pem
cat ca2.pem  >>cas.pem
cat ca3.pem  >>cas.pem

cat ca1-key.pem  >>cas.pem
cat ca2-key.pem  >>cas.pem
cat ca3-key.pem  >>cas.pem
