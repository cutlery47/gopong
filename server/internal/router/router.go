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
	for route, handler := range router.routes {
		if req.Method == route.method && req.RequestURI == route.uri {
			handler.ServeHTTP(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
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
