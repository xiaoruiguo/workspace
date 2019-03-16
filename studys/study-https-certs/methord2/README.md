---
layout: default
title: README
author: lijiaocn
createdate: 2017/08/22 13:20:17
changedate: 2017/08/22 14:21:23

---

* auto-gen TOC:
{:toc}


[OpenSSL Certificate Authority][2]中有详细的说明。

## 制作自签署的根证书

	cd RootCA
	./create_root_cert.sh

## 制作用来签署服务证书的中间证书

	cd MediaCA
	./create_media_pair.sh

RootCA/output/index.txt中会记录已经签署的证书:

	V    270820052338Z        1000    unknown    /C=CN/ST=BeiJing/O=AAlijiaocn/OU=AAlijiaocn/CN=media.ca.lijiaocn/emailAddress=admin@lijiaocn.com

将MediaCA/output/certs/media.ca-chain.cert.pem，导入浏览器:

firefox:

	选项  -> 高级 -> 证书 -> 查看证书 -> 证书机构 -> 导入 -> 勾选信任项 -> 确认

## 制作服务证书

	cd ServerCA
	./create_server1.com.sh
	./create_server2.com.sh

MediaCA/output/index.txt中会记录已经签署的证书:

	v    180901052536z        1000    unknown    /c=cn/st=beijing/l=beijing/o=lijiaocn/ou=lijiaocn/cn=server1/emailaddress=admin@lijiaocn.com
	v    180901052548z        1001    unknown    /c=cn/st=beijing/l=beijing/o=lijiaocn/ou=lijiaocn/cn=server2.com/emailaddress=admin@lijiaocn.com

## 验证证书

	$openssl verify -CAfile ./MediaCA/output/certs/media.ca-chain.cert.pem ServerCA/output/certs/server1.com.cert.pem
	ServerCA/output/certs/server1.com.cert.pem: OK
	
	$ openssl verify -CAfile ./MediaCA/output/certs/media.ca-chain.cert.pem ServerCA/output/certs/server2.com.cert.pem
	ServerCA/output/certs/server2.com.cert.pem: OK

## 撤销证书, crl方式

在没有证书撤销之前:

	$openssl crl -in MediaCA/output/crl/media-ca.crl.pem -noout -text
	...
	No Revoked Certificates.
	    Signature Algorithm: sha256WithRSAEncryption
	...
	
	

撤销server2.com的证书:

	cd Revoke
	./revoke_server1.com_crl.sh

撤销后:

	$openssl crl -in MediaCA/output/crl/media-ca.crl.pem -noout -text
	...
	Revoked Certificates:
	    Serial Number: 1000
	        Revocation Date: Aug 22 06:13:36 2017 GMT
	...

同时，在MediaCA的index.txt记录中，撤销的证书的状态变为`R`:

	$cat MediaCA/output/index.txt
	R    180901061233Z    170822061336Z    1000    unknown    /C=CN/ST=BeiJing/L=BeiJing/O=lijiaocn/OU=lijiaocn/CN=server1.com/emailAddress=admin@lijiaocn.com
	V    180901061244Z        1001    unknown    /C=CN/ST=BeiJing/L=BeiJing/O=lijiaocn/OU=lijiaocn/CN=server2.com/emailAddress=admin@lijiaocn.com

## 撤销证书，OCSP方式

[https://jamielinux.com/docs/openssl-certificate-authority/online-certificate-status-protocol.html](https://jamielinux.com/docs/openssl-certificate-authority/online-certificate-status-protocol.html)

## 参考

1. [setting-openssl-create-certificates][1]
2. [OpenSSL Certificate Authority][2]

[1]: http://www.flatmtn.com/article/setting-openssl-create-certificates  "setting-openssl-create-certificates"
[2]: https://jamielinux.com/docs/openssl-certificate-authority/index.html  "OpenSSL Certificate Authority"
