//create: 2015/04/28 11:51:41 change: 2015/05/07 16:35:33 author:lijiao
package main

import (
	"net/http"
	"log"
)

func main() {
	mux :=  http.NewServeMux()
	mux.Handle("/api/", Api{})
	server:=http.Server{Addr:":80",Handler:mux}
	log.Fatal(server.ListenAndServe())
}



