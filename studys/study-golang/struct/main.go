//create: 2017/11/17 15:47:46 Change: 2019/01/17 19:55:31 author:lijiao
package main

import (
	"fmt"
)

type A struct {
	a string
	m map[string]string
	s []string
	B
}

type B struct {
	b1 string
	b2 string
}

func display(b B) {
	fmt.Println(b.b1)
	fmt.Println(b.b2)
}

func main() {

	a := A{
		a: "1234",
		m: make(map[string]string, 5),
	}
	a.m["b"] = "134d"
	a.s = append(a.s, "slice1", "slice2")
	fmt.Println(a.m)
	fmt.Println(a.s)

	a.b1 = "b1"
	a.b2 = "b2"
	fmt.Println(a.B)

}
