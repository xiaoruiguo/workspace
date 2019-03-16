//create: 2015/09/21 11:27:07 change: 2015/09/21 17:58:53 author:lijiao
package main

import(
	"time"
	"fmt"
)

func main() {

	defer func(){
		println("exit")
	}()
	timer := time.NewTimer(1 * time.Second)

	go func(){
		for true{
			t := <-timer.C
			fmt.Printf("%v\n", t)
			timer.Reset(1 * time.Second)
		}
	}()

	time.Sleep(1000*time.Second)
}

