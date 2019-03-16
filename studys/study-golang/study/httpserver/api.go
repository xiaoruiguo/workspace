//create: 2015/04/28 13:32:57 change: 2015/09/21 10:26:01 author:lijiao
package main

import (
	"net/http"
)

type Api struct{
}

func (a Api)ServeHTTP(w http.ResponseWriter, r *http.Request){
	/*
	cookie := Cookie
	http.SetCookie(w,)
	*/
	w.Write([]byte("Hello\n"))

}

