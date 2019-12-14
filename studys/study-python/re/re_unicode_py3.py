#! /usr/bin/env python3
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

import re

if __name__ == "__main__":
    match = re.match('【(.*)】', '【我们】很好呢', re.UNICODE)
    if match:
       print(match.group(1))
    else:
       print('no match')
