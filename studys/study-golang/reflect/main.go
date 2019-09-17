// Create: 2019/06/06 17:42:00 Change: 2019/06/06 19:09:51
// FileName: main.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"github.com/lijiaocn/workspace/studys/study-golang/reflect/pkg"
	"reflect"
)

type ABC struct {
	Name string
}

func main() {
	abc := ABC{
		Name: "abc",
	}
	println(reflect.TypeOf(abc).Name())
	println(reflect.TypeOf(abc).PkgPath())

	abcd := pkg.ABCd{
		Name: "abcd",
	}
	println(reflect.TypeOf(abcd).Name())
	println(reflect.TypeOf(abcd).PkgPath())

	var i interface{}
	i = &abcd
	println(reflect.TypeOf(i).Name())
	println(reflect.TypeOf(i).PkgPath())

	println("ssss")
	println(reflect.TypeOf(ABC{}).Name())
}
