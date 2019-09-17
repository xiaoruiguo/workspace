// Create: 2019/04/02 13:47:00 Change: 2019/04/02 15:49:15
// FileName: main.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"reflect"
)

type B struct {
	BNUM1 *int
}

type A struct {
	NUM1 *int
	NUM2 *int
	B
	Bptr *B
}

func TextCompare(a, b *A) {

	va := reflect.ValueOf(*a)

	for i := 0; i < va.NumField(); i++ {
		f := va.Field(i)
		fmt.Printf("%d kind is: %s  type: %s\n", i, f.Kind(), f.Type().Name())
	}
}

func main() {
	num1 := 1
	num2 := 2
	num3 := 3
	num4 := 4

	num11 := 1
	num22 := 2
	num33 := 3
	num44 := 4

	a := A{
		NUM1: &num1,
		NUM2: &num2,
		B: B{
			BNUM1: &num3,
		},
		Bptr: &B{
			BNUM1: &num4,
		},
	}

	b := A{
		NUM1: &num11,
		NUM2: &num22,
		B: B{
			BNUM1: &num33,
		},
		Bptr: &B{
			BNUM1: &num44,
		},
	}

	fmt.Println(a)
	fmt.Println(b)

	//	if a == b {
	//		fmt.Println("a=b")
	//	} else {
	//		fmt.Println("a!=b")
	//	}

	if reflect.DeepEqual(&a, &b) {
		fmt.Println("a=b")
	} else {
		fmt.Println("a!=b")
	}

	//	TextCompare(&a, &b)
}
