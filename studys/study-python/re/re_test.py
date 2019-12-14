#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

import re

if __name__ == "__main__":
    match = re.match(ur'【(.*)】', '【我们】很好呢'.decode('utf-8'), re.UNICODE)
    if match:
       print match.group(1)
    else:
       print 'no match'
