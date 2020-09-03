package main

import (
	"reflect"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GenerateRoutes(spec *HTTPSpec, handler *httprouter.Router) {
	for path, attrs := range spec.Paths {
		for method := range attrs {
			spec.Paths[path][method].Handler = getTypeMethod(handler, strings.ToUpper(method))
		}
	}
}

func getTypeMethod(obj interface{}, name string) interface{} {
	val := reflect.TypeOf(obj)
	if method, exists := val.MethodByName(name); !exists {
		return nil
	} else {
		return method
	}
}

func invoke(method reflect.Value, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	method.Call(inputs)
}

func SetupRoutes(spec *HTTPSpec) *httprouter.Router {
	router := httprouter.New()
	GenerateRoutes(spec, router)
	for path, attrs := range spec.Paths {
		for _, m := range attrs {
			invoke(m.Handler.(reflect.Method).Func, router, path, Index)
		}
	}
	return router
}
