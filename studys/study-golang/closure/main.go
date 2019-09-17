// Create: 2019/06/18 19:22:00 Change: 2019/06/19 11:22:59
// FileName: main.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
)

func closure() func() {
	i := 0
	return func() {
		i++
		fmt.Printf("i is %d\n", i)
	}
}

func call(f func()) {
	f()
}

func main() {
	f := closure()
	f()
	f()
	f()
	f()

	j := 0
	f2 := func() {
		j++
		fmt.Printf("j is %d\n", j)
	}
	call(f2)
	j += 100
	call(f2)
}
