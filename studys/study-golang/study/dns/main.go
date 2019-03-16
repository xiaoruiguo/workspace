//create: 2017/11/09 19:42:59 change: 2017/11/09 19:47:51 author:lijiao

package main

import (
	"log"
	"net"
)

const (
	astrolabe = "registry.g.gov.cn"
)

func lookup() {
	log.Println(net.LookupHost(astrolabe))
}

func main() {
	lookup()
}
