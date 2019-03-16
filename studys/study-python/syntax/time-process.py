#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

from datetime import datetime

date = datetime.now()
print date
print type(date)
print date.strftime('%Y-%m-%d-%H-%M-%S')

import time
stamp = time.localtime()
print stamp
print type(stamp)

stamp = time.time()
print stamp
print type(stamp)
