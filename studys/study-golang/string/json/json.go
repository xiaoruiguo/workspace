// Create: 2018/10/30 11:23:00 Change: 2018/10/30 11:23:00
// FileName: json.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"strings"
)

func main() {
	var a bool
	print(a)
	ss := []string{
		"a",
		"b",
	}
	print(strings.Join(ss, ","))
}
