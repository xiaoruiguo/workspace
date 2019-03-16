//create: 2015/08/27 15:47:39 change: 2015/08/27 15:49:20 author:lijiao
package main
import(
	"fmt"
)

func main() {
	str := "aaaSAS"
	B := []byte(str)
	B[3] = 'a'
	fmt.Printf("%s\n",B[:])
}


