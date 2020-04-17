#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2020 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.

from prometheus_client import start_http_server, Gauge
import time

#if __name__ == '__main__':
#    # Start up the server to expose the metrics.
#    start_http_server(8000)
#    g = Gauge('workers_num', 'The Number of Active Workers')   # 对应指标名为 workers_num，实时值
#
#    # Generate some requests.
#    i = 0
#    while True:
#        g.set(i)      # 设置 workders 数量
#        i = i + 1
#        time.sleep(1)

import tornado.ioloop
import tornado.web

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, world")

def make_app():
    return tornado.web.Application([
        (r"/", MainHandler),
    ])


if __name__ == "__main__":
    g = Gauge('workers_num', 'The Number of Active Workers')   # 对应指标名为 workers_num，实时值
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
