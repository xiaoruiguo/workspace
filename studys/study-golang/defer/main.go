// Create: 2018/12/07 10:03:00 Change: 2018/12/07 10:03:00
// FileName: a.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

//change in defer
func det1() int {
	result := 2
	defer func() {
		println("det1")
		result = 1
	}()
	return result
}

//change pointer in defer
func det2() int {
	result := 2
	defer func(result *int) {
		println("det2")
		*result = 1
	}(&result)
	return result
}

//use name result and change in defer
func det3() (result int) {
	result = 2
	defer func() {
		println("det3")
		result = 1
	}()
	return result
}

//return is pointer
func det4() *int {
	result := 2
	defer func() {
		println("det4")
		result = 1
	}()
	return &result
}

func main() {
	a := det1()
	b := det2()
	c := det3()
	d := det4()
	println(a)
	println(b)
	println(c)
	println(*d)
}
