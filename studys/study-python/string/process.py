#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.

import re

if __name__ == "__main__":
    content="【这个#聚划算团购#宝贝不错:博朗剃须刀电动充电式3010s全身水洗刮胡刀 电动往复式男士胡须刀(分享自@手机淘宝android客户端)】https://m.tb.cn/h.ecizOob 点击链接，再选择浏览器咑閞；或復|制这段描述￥tPqJY1YnSUX￥后到👉淘♂寳♀👈"
    match = re.match('【(.*)】', content, re.UNICODE)
    if match:
        return match.group(1)
    return content
