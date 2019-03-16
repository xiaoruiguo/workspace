// Create: 2018/08/31 15:24:00 Change: 2018/08/31 15:24:00
// FileName: main.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
)

func change(s *[]string) error {
	*s = append(*s, "change")
	return nil
}

func main() {
	var array []int
	fmt.Println("len is ", len(array))
	array = append(array, 1)
	fmt.Println("len is ", len(array), array)

	s := make([]string, 10)
	s = append(s, "a", "b")
	fmt.Printf("before: %v", s)
	change(&s)
	fmt.Printf("after: %v", s)
}
