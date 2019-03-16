//create: 2015/09/18 14:21:26 change: 2015/09/18 15:42:21 author:lijiao
package main

import(
	"fmt"
	"time"
)


func main() {
	strs := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbb", "cccccccccccc", "ddd","eee","ff"}

	for _,str := range strs{
		fmt.Printf("%s\r", str)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("aaaaaaaa\b\b\b\b\\n\f")


	
}



