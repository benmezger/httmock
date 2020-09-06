package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
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
		params := r.URL.Query()
		for k, v := range method.Request.Params {
			if params.Get(k) != v {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

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
