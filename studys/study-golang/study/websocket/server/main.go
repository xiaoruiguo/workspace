//create: 2017/11/03 14:38:41 change: 2017/11/03 14:51:07 author:lijiao
package main

import (
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServer: " + err.Error())
	}
}
