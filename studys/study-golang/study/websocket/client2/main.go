//create: 2017/11/03 13:40:22 change: 2017/11/07 14:21:37 author:lijiao
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func main() {
	origin := "http://localhost/"
	//url := "ws://localhost:12345/echo"
	url := "ws://localhost:8080/v1/websocket/update"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	for true {
		if _, err := ws.Write([]byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")); err != nil {
			log.Fatal(err)
		}

		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
		time.Sleep(1 * time.Second)
	}
}
