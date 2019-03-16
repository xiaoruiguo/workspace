#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

import yaml
import os

data = {
        "title": "this is title",
        "list": ['item1','item2','item3'],
        "zh": "中文",
        "map": {
            "sub-map1": "sub-map1",
            "sub-map2": "sub-map2",
            }
        }

#print yaml.dump(data,default_flow_style=False,allow_unicode=True)
print yaml.dump(data,allow_unicode=True)
