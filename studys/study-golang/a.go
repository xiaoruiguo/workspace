// Create: 2019/04/01 11:19:00 Change: 2019/04/02 17:11:47
// FileName: a.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"strings"
)

func deleteEmptyPath(path string) string {
	for strings.Index(path, "//") != -1 {
		path = strings.Replace(path, "//", "/", -1)
	}
	return path
}

func main() {
	println(deleteEmptyPath("abcdef"))

	println(deleteEmptyPath("ab//cdef"))
	println(deleteEmptyPath("abcdef//////d"))
	println(deleteEmptyPath("/finup-store-web//health"))
}
