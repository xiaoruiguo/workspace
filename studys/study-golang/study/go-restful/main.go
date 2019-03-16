//create: 2016/05/11 16:45:03 change: 2016/05/11 18:00:04 author:lijiao
package main

import (
	"flag"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/lijiaocn/Study-Golang/study/go-restful/example"
	"net/http"
)

func main() {
	flag.Parse()
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	userHandler := example.NewExampleHandler("sqlite3", "sqlite://./sqlite.db")
	userHandler.Register(wsContainer)

	server := &http.Server{Addr: "0.0.0.0:9111", Handler: wsContainer}
	glog.Fatal(server.ListenAndServe())
}
