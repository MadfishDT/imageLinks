package main

import (
	"net/http"
)

func basicRouter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("FFAWFEFWEFWEFWEFWEf"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/aaaaa", basicRouter)
	http.ListenAndServe(":5000", nil)

}
