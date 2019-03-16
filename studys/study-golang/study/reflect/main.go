//create: 2015/07/30 17:49:24 change: 2017/11/09 17:31:31 author:lijiao
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4

	//display Type and Value
	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x))

	//modify Value
	v := reflect.ValueOf(&x).Elem()
	fmt.Println("type: ", reflect.TypeOf(v))
	fmt.Println("value: ", reflect.ValueOf(v))
	fmt.Println("settable: ", v.CanSet())
	v.SetFloat(77)
	fmt.Println("x is:", x)

	//struct
	st := T{A: 1, B: "abc"}
	v = reflect.ValueOf(st)

	r := v.Interface().(T)
	fmt.Println("r.A is ", r.A)

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, t.Field(i).Name, f.Type(), f.Interface())
	}

	m := 10
	v = reflect.ValueOf(m)
	i := v.Interface()
	a := i.(string)
	fmt.Println(v, i, a)
}
