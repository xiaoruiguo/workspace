//create: 2015/05/07 14:22:56 change: 2017/11/08 19:40:22 author:lijiao
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type (
	Inner struct {
		Mem string `json:"Mem"`
		Dat string `json:"Dat"`
	}

	Config struct {
		Log   string `json:"Log"`
		Addr  string `json:"Addr"`
		Port  int    `json:"Port"`
		Inner `json:",inline"`
	}
)

func test(a ...error) {
	fmt.Printf("%v\n", a)
	b, err := json.Marshal(a)
	fmt.Printf("%s  err: %v\n", string(b), err)
}

func main() {
	cf := Config{Log: "hello", Addr: "127.0.0.1", Port: 8080}
	cf.Mem = "memboer"
	cf.Dat = "Datassss"

	buf, err := json.Marshal(cf)

	if err != nil {
		fmt.Println(err)
		return
	}

	os.Stdout.Write(buf)

	slice := make([]string, 1, 5)
	for i := 0; i < 2; i++ {
		slice = append(slice, fmt.Sprintf("%d ", i))
	}
	buf, err = json.Marshal(slice)
	os.Stdout.Write(buf)

	fmt.Println("")

	test(errors.New("a"), errors.New("b"), errors.New("c"))
}
