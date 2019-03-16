//create: 2015/09/17 16:09:22 change: 2015/09/17 16:25:25 author:lijiao
package main


func generator(strs []string) func()string{
	next := 0
	gen := func()string{
		if next == len(strs) - 1{
			str := strs[next]
			next = 0
			return str
		}
		str := strs[next]
		next++
		return str
	}
	return gen
}


func main() {
	strs := make([]string, 0, 0)
	strs = append(strs, "a", "b", "c", "d", "e", "f")
	gen := generator(strs)

	for i:=0; i <15; i++{
		println(gen())
	}

	strs[3]="dddddd"
	for i:=0; i <15; i++{
		println(gen())
	}

}



