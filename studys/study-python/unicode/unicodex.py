#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

s1='中文'
print type(s1)

s2=u'你好'
print type(s2)

#s3=s1+s2

#s3=s1 + unicode(s2, 'utf-8')

s3=s2.encode('utf-8') 
print type(s3)
#s3=s1.encode('utf-8') + s2
print s3
