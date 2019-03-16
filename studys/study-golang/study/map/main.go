//create: 2015/09/18 09:51:00 change: 2017/11/10 16:40:03 author:lijiao
package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	println(len(m))

	m["1"] = "a"
	m["2"] = "b"
	m["3"] = "c"
	m["4"] = "d"

	println(len(m))

	v, ok := m["5"]
	println("not exist ", v, ok)

	v = m["5"]
	println("not exist ", v)

	switch m["3"] {
	case "c":
		println("is 3")
	default:
		println("not known")
	}

	for i, j := range m {
		fmt.Println(i, j)
	}
}
