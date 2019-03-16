---
layout: default
title: Readme
createdate: 2017/10/22 21:49:12
changedate: 2018/02/18 09:12:03

---

# Readme
创建时间: 2016/12/04 22:01:14  修改时间: 2016/12/05 01:03:47 作者:lijiao

----

## 摘要

Vim配置文件与常用插件。

## 安装

需要使用7.4以上的vim，在macOS上可以用下面的命令更新：

	brew install macvim    //YCM推荐使用macvim
	brew install vim --with-lua --with-override-system-vi   //或者用brew更新, brew info vim查看更多选项

依赖cmake:

	brew install cmake

用git复制到本地，执行安装脚本：

	git clone --recursive https://github.com/lijiaocn/vim-config.git
	cd vim-config/vim
	./install.sh

## YCM插件说明

语法级别的自动补全。

[https://github.com/Valloric/YouCompleteMe](https://github.com/Valloric/YouCompleteMe)

在macOS上安装：

	cd .vim/bundle
	git clone https://github.com/Valloric/YouCompleteMe.git

如果要支持C系列语言：

	cd ~/.vim/bundle/YouCompleteMe
	git submodule update --init --recursive
	./install.py --clang-completer

否则：

	cd ~/.vim/bundle/YouCompleteMe
	./install.py

如果需要支持其它语言，使用对应的参数：

	--go-completer   //go
	--js-completer   //js
	--rust-completer //rust
	--all            //所有支持的语言

## 文献
