//create: 2015/08/19 15:29:16 change: 2016/04/28 16:02:46 author:lijiao
package main

import (
	"fmt"
	"strings"
)

func main() {

	var a string = "hello"
	if "aaa" == "aaa" {
		fmt.Println("eqqa", a)
	}

	fmt.Println("bbb:" + a)
	strs := strings.Split("/a/b/c/d", "/")
	if strs[0] == "" {
		fmt.Println("empty")
	}

}
