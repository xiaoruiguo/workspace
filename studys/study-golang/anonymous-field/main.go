//create: 2017/11/15 17:02:58 change: 2017/11/15 19:39:04 author:lijiao
package main

import (
	"fmt"
)

type A struct {
	B
}

type B struct {
	b string
}

func (*B) sayHello() {
	fmt.Printf("Hi, i'm B, i'm say hello to you.\n")
}
func (b *B) SetB() {
	b.b = "i have been set"
}

func main() {
	a := A{
		B: B{
			b: "i'm  b in B",
		},
	}

	fmt.Printf("a.b=%s\n", a.b)
	a.sayHello()
	a.SetB()
	fmt.Printf("a.b=%s\n", a.b)
}
