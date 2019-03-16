//create: 2016/05/05 15:02:23 change: 2016/05/11 18:07:35 author:lijiao
package example

import (
	"github.com/emicklei/go-restful"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

type ExampleHandler struct {
	engine *xorm.Engine
}

type ExampleResource struct {
	ID int64 `json:"id"`
}

func NewExampleHandler(driverName string, database string) ExampleHandler {
	engine, err := xorm.NewEngine(driverName, database)
	if err != nil {
		glog.Fatal("xorm new engine fail: %s", err)
	}
	return ExampleHandler{engine: engine}
}

//GET /example/resource/{id}
func (s ExampleHandler) getResource(request *restful.Request, response *restful.Response) {
	glog.V(10).Infof("%s %s --> %s\n", time.Now().Format(time.RFC3339), request.Request.RemoteAddr, request.Request.RequestURI)
	id, err := strconv.ParseInt(request.PathParameter("id"), 10, 64)
	if err != nil {
		glog.Errorf("Invalid id: %s", err)
		response.WriteErrorString(500, "Internal Error")
		return
	}
	res := ExampleResource{ID: id}
	response.WriteEntity(res)
}

func (s ExampleHandler) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/example").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(
		ws.
			GET("/resource/{id}").
			To(s.getResource).
			Doc("get example resource").
			Param(
			ws.
				PathParameter("id", "example resource's id").
				DataType("string")).
			Writes(ExampleResource{}))

	container.Add(ws)
}
