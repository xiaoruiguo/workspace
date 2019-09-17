// Create: 2019/03/29 10:54:00 Change: 2019/03/29 10:56:31
// FileName: split.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"strings"
)

func main() {
	str := "a,b    ,c,d, e,f"
	array := strings.Split(str, ",")
	for i, v := range array {
		println(i, strings.TrimSpace(v))
	}
}
