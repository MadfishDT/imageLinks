package main

import (
	"encoding/json"
	"net/http"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	mem := Member{"Alex", 10, true}
	jsonByte, _ := json.Marshal(mem)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonByte)
	})

	http.HandleFunc("/json/", TestRouter)
	http.ListenAndServe(":5000", nil)

}
