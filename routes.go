package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func GenerateRoutes(spec *HTTPSpec, handler *httprouter.Router) {
	for _, methods := range spec.Paths {
		for name, method := range methods {
			method.SetHandler(handler, strings.ToUpper(name))
		}
	}
}

func GenerateHandler(method *HTTPSpecMethod) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for k, v := range method.Response.Header {
			w.Header().Add(k, v)
		}
		w.WriteHeader(method.Response.Status)
		w.Write([]byte(method.Response.Payload))
	}

}

func SetupRoutes(spec *HTTPSpec) *httprouter.Router {
	router := httprouter.New()
	GenerateRoutes(spec, router)
	for path, attrs := range spec.Paths {
		for _, m := range attrs {
			m.Invoke(router, path, GenerateHandler(m))
		}
	}
	return router
}
