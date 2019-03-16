//create: 2015/08/21 16:26:48 change: 2017/11/10 16:20:42 author:lijiao
package main

import (
	"fmt"
	//	"time"
)

var i int

func geti() int {
	return i + 1
}

func main() {
	/*
		for{
			fmt.Println("1")
			time.Sleep(1*time.Second)
		}
	*/
	/*
		i = 1
		for j:=geti();;{
			fmt.Println(j)
			time.Sleep(1*time.Second)
		}
	*/

	var v []string
	for i := 0; i < 5; i++ {
		a := fmt.Sprintf("hello %d", i)
		v = append(v, a)
	}
	fmt.Println(v)

}
