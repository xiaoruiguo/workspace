//create: 2015/09/10 20:46:15 change: 2015/09/10 20:50:38 author:lijiao
package   main

import (
	"strings"
	"fmt"
)

func main() {
	var b []byte
	
	b = append(b, 'a', 'b','c','1','2')

	strb := string(b)

	fmt.Printf("%s\n",b)
	if strings.HasPrefix(strb, "abc"){
		fmt.Printf("have prefix abc\n")
	}
	if strings.HasPrefix(strb, "abcd"){
		fmt.Printf("don't have prefix abcd\n")
	}
}


