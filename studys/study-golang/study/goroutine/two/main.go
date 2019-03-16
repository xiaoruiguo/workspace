//create: 2015/09/21 16:53:39 change: 2015/09/21 17:00:12 author:lijiao
package main

import(
	"time"
	"fmt"
)
func main() {

	for i := 0; i < 5; i++{
		go func(){
			t := time.Now()
			fmt.Printf("%v\n", t)
		}()
	}

	time.Sleep(10*time.Second)
}
	

