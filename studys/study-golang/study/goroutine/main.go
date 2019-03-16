//create: 2015/09/17 17:13:57 change: 2016/04/10 16:47:22 author:lijiao
package main

import(
	"time"
)

func main() {
	
	i := 10

go func(){
	for true{
		println(i)
		time.Sleep(1 * time.Second)
	}
}()
	go gox()

	time.Sleep(5 * time.Second)
	i = 12
	time.Sleep(100 * time.Second)
}





