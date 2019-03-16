//create: 2015/05/08 14:03:36 change: 2015/05/08 14:05:58 author:lijiao
package main

import (
	"fmt"
)



type handle func(w int, b string)string

func main() {
	hmap := make(map[string] handle)
	fmt.Println(hmap["aa"])
}




