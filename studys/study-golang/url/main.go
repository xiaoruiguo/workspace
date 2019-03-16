// Create: 2018/08/07 14:51:00 Change: 2018/08/07 14:51:00
// FileName: main.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"html"
	"net/url"
)

func main() {
	//url编码
	str := "http://aaawd.com?a=  中文- '  _. &b=我们    dfad"
	unstr := "%2f"
	fmt.Printf("url.QueryEscape:%s", url.QueryEscape(str))
	fmt.Println()
	s, _ := url.QueryUnescape(unstr)
	fmt.Printf("url.QueryUnescape:%s", s)
	fmt.Println()
	//字符转码
	hstr := "<"
	hunstr := "&lt"
	fmt.Printf("html.EscapeString:%s", html.EscapeString(hstr))
	fmt.Println()
	fmt.Printf("html.UnescapeString:%s", html.UnescapeString(hunstr))
	fmt.Println()
}
