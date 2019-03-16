import traceback
import time
from os.path import basename
import os
import thread
import logging
import kazoo
from kazoo.client import KazooClient
from kazoo.client import KazooState
from kazoo.recipe.election import Election

logging.basicConfig()

class ZkJob(object):
    def __init__(self,hosts="127.0.0.1:2181",room="/CoxElectRoom9",\
            chair="chair",debug=None,timeout=60000):
        self.hosts=hosts
        self.room=room
        self.zk=KazooClient(hosts=hosts, timeout=timeout)
        self.debug=debug
        self.imleader=False
        self.notify=None
        self.reElect=True
        self.zk.add_listener(self.__statListen)
        self.e = Election(self.zk,self.room)

    def __statListen(self,state):
        if state == KazooState.LOST:
            if self.debug != None:
                print "Lost"
            self.imleader = False
        elif state == KazooState.SUSPENDED:
            if self.debug != None:
                print "Suspended"
            self.imleader = False
        elif state == KazooState.CONNECTED:
            if self.debug != None:
                print "Connected"
                self.e.cancel()
                self.e = Election(self.zk,self.room)
                self.__elect()
        else:
            print 'Something is Wrong!'

    def __del__(self):
        if self.zk != None:
            self.zk.stop()
            self.zk.close()

    def __elect(self):
        def setLeader():
            self.reElect = False
            self.imleader = True
            if self.notify != None:
                self.notify()
            self.e.run(setLeader)

    def elect(self):
        self.zk.start()

    def isLeader(self):
        return self.imleader

    def setNotify(self, notifyFunc):
        self.notify = notifyFunc
    
    def close(self):
        if self.zk != None:
            self.zk.stop()
            self.zk.close()

if __name__ == "__main__":
    zk = ZkJob(debug=1)
    print("pid %d" % os.getpid())
    print("ThreadID %d" % thread.get_ident())

    def notify():
        print("I'm Leader now ...")
        print("pid %d" % os.getpid())
        print("ThreadID %d" % thread.get_ident())

    zk.setNotify(notify)
    zk.elect()

    time.sleep(1)
    while True:
        print zk.isLeader()
        time.sleep(1)

    time.sleep(1000)
    xx.close()


