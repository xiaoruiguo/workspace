//create: 2015/08/24 16:54:00 change: 2016/05/11 13:41:00 author:lijiao
package main

import (
	"flag"
	"fmt"
)

func main() {
	//	var list []string
	f1 := flag.String("f1", "fffff11111", "flag1")
	f2 := flag.String("f2", "fffff11111", "flag1")
	//	f3 := flag.Var(&list, "list", "lists")

	flag.Parse()
	fmt.Println(f1)
	fmt.Println(f2)
	//	fmt.Println(f3)
}
