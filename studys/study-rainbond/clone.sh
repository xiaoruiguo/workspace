#! /bin/sh
#
# clone.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

git clone https://github.com/goodrain/rainbond-ui.git
ln -s `pwd`/rainbond-ui $GOPATH/src/github.com/goodrain/rainbond-ui

virtualenv  python-env
source python-env/bin/activate

#	rainbond-ui的启动方式：
#	# python 2.7.8
#	pip install -r requirements.txt
#	
#	# console.py 
#	cp ./console.py /etc/goodrain/console.py
#	
#	# update console.py
#	# 配置数据库相关信息
#	# eg
#	DATABASES = {
#	    'default': {
#	        'ENGINE': 'django.db.backends.mysql',
#	        'NAME': 'console',
#	        'USER': 'write',
#	        'PASSWORD': '12345678',
#	        'HOST': '127.0.0.1',
#	        'PORT': 3306,
#	    },
#	}
#	
#	export REGION_TAG=cloudbang
#	
#	python manage.py migrate
#	python manage.py runserver

go get -d github.com/goodrain/rainbond
ln -s `pwd`/rainbond $GOPATH/src/github.com/goodrain/rainbond
#	cd $GOPATH/src/github.com/goodrain/rainbond
#	make build
