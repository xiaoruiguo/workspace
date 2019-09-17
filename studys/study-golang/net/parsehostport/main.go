// Create: 2019/03/21 17:33:00 Change: 2019/03/21 17:58:45
// FileName: main.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"net"
)

func main() {
	addr := "10.10.10.0:"
	host, port, err := net.SplitHostPort(addr)
	fmt.Println("host is:", host, "port: ", port, "err: ", err)
}
