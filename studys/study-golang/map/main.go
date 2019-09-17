//create: 2017/10/26 22:04:34 Change: 2019/03/29 16:36:23 lijiaocn@foxmail.com
package main

func main() {
	m := make(map[string]bool)
	v, ok := m[""]
	if ok {
		println("found")
	} else {
		println("not found: ", v)
	}
	if v != true {
		println("not true")
	}
	/*
		println(len(m))
		println("1234")
		m["a"] = 1
		m["b"] = 2
		m["c"] = 3
		m["a"] = 4
		println(m["a"])

		for _, v := range m {
			println(v)
		}
	*/
}
