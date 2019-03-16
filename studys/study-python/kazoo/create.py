import logging
import time
import kazoo
import thread
from kazoo.client import KazooClient
from kazoo.client import KazooState
from kazoo.recipe.election import Election

username='test'
password='123'

room='/3kazoo'
node=room+'/nodebdb22'
node_acl=kazoo.security.make_digest_acl(username=username, password=password,\
        read=True,write=True,create=True,delete=True,all=True)

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

if not zk.exists(room):
    zk.ensure_path(room,acl=[node_acl])

xx=zk.create(path=node, value=b'data',acl=[node_acl],\
       ephemeral=True,sequence=True,makepath=True)
print xx

@zk.ChildrenWatch(room)
def watch_children(children):
    print("Children :%s" % children)

time.sleep(1000)
zk.stop()
zk.close()
