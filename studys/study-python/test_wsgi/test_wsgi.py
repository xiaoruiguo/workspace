import os
import eventlet
from eventlet import wsgi, listen
from paste import deploy

class UrlFilter(object):
    @classmethod
    def factory(cls, global_conf, **local_conf):
        def filter(app):
            return UrlFilter(app, global_conf)
        return filter
    def __init__(self, app, global_conf):
        self.app = app
        self.global_conf = global_conf

    def __call__(self, environ, start_response):
        #TODO: You can check something at here
        print 'Url Filter'
        print environ
        print type(environ)
        return self.app(environ,start_response)

class AuthFilter(object):
    @classmethod
    def factory(cls, global_conf, **local_conf):
        def filter(app):
            return AuthFilter(app, global_conf)
        return filter

    def __init__(self, app, global_conf):
        self.app = app
        self.global_conf = global_conf

    def __call__(self, environ, start_response):
        #TODO: You can check something at here
        print 'Auth Filter'
        return self.app(environ,start_response)

class WelcomeApp(object):
    @classmethod
    def factory(cls, global_conf, **local_conf):
        return cls()

    def __init__ (self, *args, **kwargs):
        pass

    def __call__(self, env, start_response):
        print 'WelcomeApp Filter'
        start_response('200 OK',())
        return ['welcome']

if __name__ == '__main__':
    _path = os.path.dirname(os.path.abspath(__file__))
    f = 'config:%s' % os.path.join(_path,'test_paste.ini')
    host = '0.0.0.0'

#    port = 8080
#    appname='test1'  # app name
#
#    _socket = listen((host,port))
#    app = deploy.loadapp(f, appname)
#
#    server = eventlet.spawn(wsgi.server, _socket, app)
#    server.wait()

    port2 = 8080
    appname2='admin'  # app name

    _socket2 = listen((host,port2))
    app2 = deploy.loadapp(f, appname2)

    server2 = eventlet.spawn(wsgi.server, _socket2, app2)
    server2.wait()

