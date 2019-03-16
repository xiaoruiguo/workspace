---
layout: default
title: README

---

# README
创建时间: 2016/04/08 21:31:17  修改时间: 2016/04/08 22:10:19 作者:lijiao

----

## 摘要

使用netperf对网络性能进行测试。

## 准备

需要安装有gcc，对netperf进行编译。

	$./0_install.sh

## Server端

启动Server:

	$./1_start_server.sh

## 启动测试

进入cases目录，在config中完成配置后，执行对应的脚本即可。

例如要进行“TCP_RR”测试，只需要执行:

	$./TCP_RR.sh
	MIGRATED TCP REQUEST/RESPONSE TEST from 0.0.0.0 (0.0.0.0) port 0 AF_INET to 127.0.0.1 () port 0 AF_INET : first burst 0
	Local /Remote
	Socket Size   Request  Resp.   Elapsed  Trans.
	Send   Recv   Size     Size    Time     Rate
	bytes  Bytes  bytes    bytes   secs.    per sec

	16384  87380  1        1       10.00    56195.43
	16384  87380

# 文献
