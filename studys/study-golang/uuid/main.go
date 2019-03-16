// Create: 2018/04/10 11:00:00 Change: 2018/04/10 11:03:50
// FileName: main.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	u2 := uuid.NewV4()
	s := fmt.Sprintf("%s", u2)
	println(s)
}
