---
layout: default
title: README
author: lijiaocn
createdate: 2017/08/22 14:25:08
changedate: 2017/08/22 14:38:05

---

* auto-gen TOC:
{:toc}

## 创建自签署的根证书

	cd ./RootCA
	./create_root_cert.sh

## 创建Server证书

	cd ./ServerCA
	./create_server_cert.sh
