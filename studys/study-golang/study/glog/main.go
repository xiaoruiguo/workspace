//create: 2016/05/11 13:48:28 change: 2016/05/11 15:24:34 author:lijiao
package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("hello")
	glog.Flush()
	//glog.Fatal("XXX")
	glog.V(3).Info("XXX")
}
