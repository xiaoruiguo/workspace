//create: 2015/08/12 17:00:12 change: 2015/08/12 17:37:25 author:lijiao
package main

import(
	"fmt"
)

type A struct{
	Value int
}

type B struct{
	Value int
}

type I interface{
	Hello()
}

func (b B)Hello(){
	return
}

func main() {

	var v interface{} = 'a'     //v is dynamic value
	i,err := v.(int32)
	fmt.Println(i,err)

	var w I                    //w is interface type
	w = B{Value:33}
	x,err := w.(B)
	fmt.Println(x,err)

	a := A(w.(B))
	fmt.Println(a)

//	x,err = w.(A)
//	fmt.Println(x,err)
}


