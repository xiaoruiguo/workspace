//create: 2015/09/18 15:43:07 change: 2015/09/18 15:45:08 author:lijiao
package main

import(
	"github.com/lijiaocn/GoPkgs/terminal"
	"time"
)

func main() {
	
	strs := []string{"aaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbb", "cccccccccccc", "ddd","eee","ff"}

	for _,str := range strs{
		terminal.Reset()
		println(str)
		time.Sleep(1*time.Second)
	}
}
	
