//create: 2015/07/29 15:09:33 change: 2015/07/29 15:15:38 author:lijiao
package main
import (
	"fmt"
)

func xxx(args ...string){
	s := make([]string, 0)
	s = append(s, "1","2","3")
	s = append(s, args...)
	println(len(s))
	fmt.Println(s)
}

func main() {
	xxx("a","b","c")
}

