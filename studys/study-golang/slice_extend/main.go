// Create: 2018/04/19 14:09:00 Change: 2018/04/19 14:28:14
// FileName: main.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
)

func test(a, b string, args ...string) {
	fmt.Println(args)
}

func main() {
	//	var s []string
	s := make([]string, 0)
	s = append(s, "a")
	s = append(s, "b")
	test("1", "2", s...)
}
