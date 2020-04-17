#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2020 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.

from prometheus_client import Gauge
from tornado.web import Application, RequestHandler
from tornado.ioloop import IOLoop
from tornado_prometheus import PrometheusMixIn, MetricsHandler

class SampleApp(PrometheusMixIn, Application):
    pass

if __name__ == '__main__':
  app =  SampleApp([
      (r"/metrics", MetricsHandler),
  ])

  # 定义一个指标
  g = Gauge('workers_num', 'The Number of Active Workers')   # 对应指标名为 workers_num，实时值
  # 在 workker 数量发生变化时，更新指标值
  g.set(5)

  app.listen(8888)
  IOLoop.current().start()
