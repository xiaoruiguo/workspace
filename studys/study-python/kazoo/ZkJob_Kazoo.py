import traceback
import time
from os.path import basename
import os
import thread
from kazoo.client import KazooClient
from kazoo.client import KazooState
from kazoo.recipe.election import Election

class ZkJob(object):
    def __init__(self,hosts="127.0.0.1:2181",room="/CoxElectRoom34",\
            chair="chair",debug=None,timeout=60000):
        self.hosts=hosts
        self.room=room
        self.chair=chair
        self.timeout=timeout
        self.zk=KazooClient(hosts=hosts, timeout=timeout)
        self.debug=debug
        self.chairNum=None
        self.imleader=False
        self.zk.add_listener(self.__statListen)

    def __del__(self):
        if self.zk != None:
            self.zk.stop()
            self.zk.close()

    def __statListen(self,state):
        if state == KazooState.LOST:
            if self.debug != None:
                print 'Lost'
            self.imleader = False
        elif state == KazooState.SUSPENDED:
            if self.debug != None:
                print 'Suspended'
            self.imleader = False
        elif state == KazooState.CONNECTED:
            if self.debug != None:
                print 'Connected'

    def isLeader(self):
        return self.imleader

    def setNotify(self, notifyFunc):
        self.notify = notifyFunc

    def elect(self):
        self.zk.start()
        if not self.zk.exists(self.room):
            self.zk.ensure_path(self.room)
        
        self.chairNum = self.zk.create(path=self.room+'/'+self.chair, \
               ephemeral=True,sequence=True,makepath=True)
        self.chairNum = basename(self.chairNum)

        @self.zk.ChildrenWatch(self.room)
        def watch_room(chairs):
            chairs.sort()
            print('Chairs:%s' % chairs)
            if self.chairNum == chairs[0]:
                self.imleader = True
                self.notify()

    def close(self):
        if self.zk != None:
            self.zk.stop()
            self.zk.close()

if __name__ == "__main__":
    zk = ZkJob(debug=1)
    print("main: pid %d" % os.getpid())
    print("main: ThreadID %d" % thread.get_ident())

    def notify():
        print("notify: I'm Leader now ...")
        print("notify: pid %d" % os.getpid())
        print("notify: ThreadID %d" % thread.get_ident())

    zk.setNotify(notify)
    zk.elect()
    while True:
        if zk.isLeader():
            print "Leader is Working..."
            time.sleep(1)
            continue
        print "Not Leader, Waiting ..."
        time.sleep(1)

    time.sleep(1000)
    zk.Close()
