//create: 2015/10/02 15:10:28 change: 2015/10/02 15:14:12 author:lijiao
package main

import(
	"net/url"
	"log"
	"fmt"
)

func main() {
	uri,err := url.Parse("?query=true&abc=123")
	if err != nil{
		log.Fatal(err)
	}
	values := uri.Query()
	fmt.Printf("%v\n",values)
}



