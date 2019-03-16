//create: 2018/01/02 12:26:00 change: 2018/01/02 12:29:23 lijiaocn@foxmail.com

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		// Why 8061 bytes? Because the response header on my computer
		// is 132 bytes, adding up the entire response to 8193 (1 byte
		// over 8kb)
		if _, err := w.Write(make([]byte, 19000)); err != nil {
			fmt.Println(err)
			return
		}
	})))
}
