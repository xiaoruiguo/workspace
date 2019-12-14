#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

import os
import urlparse

url = 'http://www.**.net/images/logo.gif?abc=1&cd=1'
filename = os.path.basename(url)
print(filename)
print(urlparse.urlsplit(url))
