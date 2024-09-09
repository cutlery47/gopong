package router

import "net/http"

type Router struct {
	routes map[route]http.Handler
}

type route struct {
	method string
	uri    string
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req_route := route{
		method: req.Method,
		uri:    req.RequestURI,
	}

	req_handler := router.routes[req_route]
	if req_handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	req_handler.ServeHTTP(w, req)
}

func NewRouter(connectionHandler http.Handler) *Router {
	connectionRoute := route{
		method: "GET",
		uri:    "/",
	}

	routes := map[route]http.Handler{}

	routes[connectionRoute] = connectionHandler

	router := &Router{
		routes,
	}
	return router
}
