import logging
import time
import kazoo
import thread
from kazoo.client import KazooClient
from kazoo.client import KazooState
from kazoo.recipe.election import Election

#logging.basicConfig()

username='test'
password='123'

node_acl=kazoo.security.make_digest_acl(username=username, password=password,\
        read=True,write=True,create=True,delete=True,all=True)

def notify():
    while True:
        print 'i am leader now'
        time.sleep(1)

def election(zk,path):
    if not zk.exists(path):
        zk.create(path,makepath=True)

    elect = Election(zk,path)
    elect.run(notify)
    print elect.contenders()


def stat_listener(state):
    if state == KazooState.LOST:
        print 'Lost'
    elif state == KazooState.SUSPENDED:
        print 'Suspended'
    elif state == KazooState.CONNECTED:
        print 'Connected'
    else:
        print 'Other State'

zk = KazooClient(hosts='127.0.0.1:2181',\
        auth_data=[('digest',username+':'+password)])

zk.add_listener(stat_listener)
zk.start()

election(zk,"/x/election1")

time.sleep(1000)
zk.stop()
zk.close()
