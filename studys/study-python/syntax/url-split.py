#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

"""

"""
import urlparse

url="https://mmbiz.qpic.cn/mmbiz_jpg/NxcXQjmQe5ChtbJY88n04tCBkDeQiau2ZAqJ26JDd7T6qpnunoWezuUEhPrUEWVwFR4luMvxSKv44zK1tm4SvzA/640?wx_fmt=jpeg&test=text"
print url.split("/")[-2] + ".png"

print urlparse.parse_qs(urlparse.urlsplit(url).query)["wx_fmtx"][0]

