//create: 2017/11/15 10:42:12 change: 2017/11/15 10:44:24 author:lijiao
package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "abc"
	strp := &str

	fmt.Printf("str type: %s\n", reflect.TypeOf(str))
	fmt.Printf("strp type: %s\n", reflect.TypeOf(strp))
	fmt.Printf("*str type: %s\n", reflect.TypeOf(*strp))
}
