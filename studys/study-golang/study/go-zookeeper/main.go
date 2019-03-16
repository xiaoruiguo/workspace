//create: 2016/04/26 18:17:05 change: 2016/04/26 18:37:27 author:lijiao
package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

func zkReconnect(retry int, servers []string, timeout time.Duration) (zk.Conn, error) {
	for j := i; j > 0; j-- {
		conn, _, err := zk.Connect(cfs.ZKServers, cfs.ZKTimeout)
		if err != nil {
			log.Info("zk connect faild: %d, %v", i, err)
			return conn, nil
		}
	}

	log.Error("zk connect faild, have retry %d times", i)
	return nil, err
}

func main() {
	servers := make([]string, 1)
	servers = append(servers, "192.168.166.65:2181")
	path := "/test/update/version"

	conn, err := zkReconnect(5, servers, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	conn.Create(path, "test", zk.FlagSequence)

}
