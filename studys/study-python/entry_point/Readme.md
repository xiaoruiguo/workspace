---
layout: default
title: Readme

---

# Readme
创建时间: 2015/12/15 11:07:25  修改时间: 2015/12/15 11:29:17 作者:lijiao

----

## 摘要

这里展示了如何使用entry_points。

## 说明

一共有三个entry_point:

	一个控制台命令: lijiao
	一个包含两个函数的Group:  lijiao.test
	包含的两个函数是: hello、bye

setup.py: 是直接用setuptools进行安装

setup_cfg.py: 使用了pbr, pbr要求项目的代码是用git管理的

call_entry_points.py: 展示了如何调用entry_point, 注意拷贝到项目外执行。

## 文献
