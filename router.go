package main

import (
	"net/http"
)
type Router struct {
	name string
}


type routerHanler = func(ResponseWriter, *Request)

http.han
func (r *Router) addGetRouter(path string, handler routerHanler) {

}
