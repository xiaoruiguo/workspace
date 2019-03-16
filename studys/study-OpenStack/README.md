---
layout: default
title: README

---

# README
创建时间: 2016/04/19 13:43:49  修改时间: 2016/04/19 15:59:14 作者:lijiao

----

## 摘要

Openstack组件学习。

## 安装依赖

	sudo apt-get install python-pip
	sudo apt-get install python-dev
	pip  install virtualenv

## 安装说明

以keystone为例:

	cd KeyStone
	./create-env.sh

在执行其它操作之前进入keystone的virtualenv:

	vagrant@~~/KeyStone$ source virtualenv/bin/activate
	(virtualenv) vagrant@~~/KeyStone$

安装:

	(virtualenv) vagrant@~~/Install$ cd Install
	(virtualenv) vagrant@~~/Install$ ./install.sh

安装成功以后，就可以使用keystone相关命令了:

	(virtualenv) vagrant@~~/Install$ keystone-
	keystone-all          keystone-manage       keystone-wsgi-admin   keystone-wsgi-public

退出virtualenv:

	(virtualenv) vagrant@~~/Install$ deactivate
	vagrant@~~/Install$

## 文献
