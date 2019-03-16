from paste import  httpserver
def app(env, start_response):
    start_response('200 OK',[('content-type','text/html')])
    return ['Hello World!']

if __name__ == '__main__':
    httpserver.serve(app,host='192.168.187.5',port='8080')
