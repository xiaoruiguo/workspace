//create: 2015/10/14 16:35:29 change: 2015/10/14 16:37:24 author:lijiao
package main

import(
	_ "net/http/pprof"
	"net/http"
	"log"
)

func main() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}

