//create: 2015/05/06 14:13:27 change: 2017/11/09 19:09:13 author:lijiao
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	client := &http.Client{}
	res, err := client.Do(req)
	fmt.Println(res, err)
}
