//create: 2015/09/17 19:38:58 change: 2015/09/18 09:41:04 author:lijiao
package main

import(
	"sync/atomic"
)

func main() {

	var n int32 = 0

	finished := make(chan interface{}, 10)

	for i:=0; i < 30; i++{
		go func(){
			println(atomic.AddInt32(&n, 1))
			finished <- 0
		}()
	}

	for i:=0; i < 30; i++{
		<-finished
	}
}
