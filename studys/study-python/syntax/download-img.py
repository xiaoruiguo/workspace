#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

import urllib2

url = "http://wx.qlogo.cn/mmhead/Q3auHgzwzM58ua3QdhFf8k1h0VlcRQUDbcwxEHjmZ6dv5mYeXdNqjw/0"

#urllib.urlretrieve(url, "a.png")

#resp = urllib2.urlopen(url)
f = open("./b/a.png","wb")
#f.write(resp.read())
f.write(urllib2.urlopen(url).read())
f.close()
