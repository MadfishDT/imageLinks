package main

import (
	"net/http"
)

//TestRouter this is test router
func TestRouter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("FFAWFEFWEFWEFWEFWEf"))
}
