import traceback
import zookeeper
import time
from os.path import basename
import os
import thread

class ZkJob(object):
	def __init__(self,hosts="127.0.0.1:2181",room="/ElectRoom",\
			chair="chair",debug=None,timeout=60000):
		self.hosts=hosts
		self.room=room
		self.chair=chair
		self.timeout=timeout
		self.zk=zookeeper.init(self.hosts,None,self.timeout)
		self.debug=debug
		self.chairNum=None
		self.lock=thread.allocate_lock()
		self.imleader=False

	def __del__(self):
		if self.zk != None:
			zookeeper.close(self.zk)

	def __createRoom(self):
		nodes = self.room.split('/')
		lastpath = ""
		for node in nodes:
			if len(node) <= 0:
				continue
			lastpath=lastpath+"/"+node
			if zookeeper.exists(self.zk, lastpath) == None:
				self.room = zookeeper.create(self.zk, lastpath,"",[{"perms":\
					zookeeper.PERM_CREATE|zookeeper.PERM_DELETE|\
					zookeeper.PERM_READ, "scheme":"world","id":"anyone"}], 0)

	def __enterRoom(self):
		self.chairNum = zookeeper.create(self.zk, self.room+"/"+self.chair,"",\
			[{"perms":zookeeper.PERM_READ,"scheme":"world","id":"anyone"}], \
			zookeeper.EPHEMERAL|zookeeper.SEQUENCE)
		self.chairNum = basename(self.chairNum)

	def __reConect(self):
		self.imleader = False
		zookeeper.close(self.zk)
		self.zk=zookeeper.init(self.hosts,None,self.timeout)
		stat = zookeeper.state(self.zk)
		while stat != zookeeper.CONNECTED_STATE :
			time.sleep(1)
			stat = zookeeper.state(self.zk)

	def isLeader(self):
		return self.imleader

	def setNotify(self, notifyFunc):
		self.notify = notifyFunc

	def elect(self):
		try:
			self.lock.acquire()
			if zookeeper.exists(self.zk, self.room) == None:
				self.__createRoom()
			if self.chairNum == None:
				self.__enterRoom()

			def chairsWatcher(zk,type,state,path):
				try:
					if self.debug != None:
						print("%s Func: %s" % \
							(self.__class__,self.elect.__name__))
						print("\tChildren changed:\n\t\ttype: %d \n\t\t"\
								"state:%s\n\t\tpath:%s" %(type, state, path))
					if type == zookeeper.SESSION_EVENT:
						self.chairNum = None
						self.elect()
					if type == zookeeper.CHILD_EVENT:
						self.elect()
				except (zookeeper.ConnectionLossException,\
						zookeeper.OperationTimeoutException),e:
						self.__reConect()
						self.lock.release()
						self.elect()

			chairs = zookeeper.get_children(self.zk, self.room, chairsWatcher)
			chairs.sort()
			if self.debug != None:
				print("%s Func: %s" % (self.__class__,\
					self.elect.__name__))
				print("\tChairs: %s" % chairs)
			if self.chairNum == chairs[0]:
				if self.debug != None:
					print("%s Func: %s" % \
						(self.__class__,self.elect.__name__))
					print("\tI am the Leader: %s" % self.chairNum)

				# Do Leader's Job at Here
				self.imleader = True
				self.notify()
				self.lock.release()
				return True
			else:
				if self.debug != None:
					print("\tI am not Leader: %s" % self.chairNum)
				self.imleader = False
				self.lock.release()
				return False
		except (zookeeper.ConnectionLossException,\
				zookeeper.OperationTimeoutException),e:
				self.__reConect()
				self.lock.release()
				self.elect()
	
	def close(self):
		if self.zk != None:
			zookeeper.close(self.zk)
			self.zk = None

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

    zk.close()
