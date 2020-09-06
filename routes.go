package main

import (
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

func SetupRoutes(spec *HTTPSpec) *httprouter.Router {
	router := httprouter.New()
	GenerateRoutes(spec, router)
	for path, attrs := range spec.Paths {
		for _, m := range attrs {
			m.Invoke(router, path, Index)
		}
	}
	return router
}
