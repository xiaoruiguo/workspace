//create: 2015/09/16 11:54:45 change: 2015/09/16 11:56:13 author:lijiao
package main

import(
	"time"
	"fmt"
)


func main() {
loop:
	for true{
		fmt.Println("before break")
		break loop
		fmt.Println("after break")
		time.Sleep(1*time.Second)
	}
}
