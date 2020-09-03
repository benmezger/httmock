package main

import (
	"reflect"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GenerateRoutes(spec *HTTPSpec, handler *httprouter.Router) map[string][]interface{} {
	func_handlers := make(map[string][]interface{})

	for endpoint, attrs := range spec.Paths {
		for method := range attrs {
			spec.Paths[endpoint][method].Handler = getTypeMethod(handler, strings.ToUpper(method))
			func_handlers[endpoint] = append(
				func_handlers[endpoint],
				getTypeMethod(
					handler,
					strings.ToUpper(method)))
		}
	}
	return func_handlers

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
	r := GenerateRoutes(spec, router)
	for endpoint, funcs := range r {
		for _, m := range funcs {
			method := m.(reflect.Method)
			invoke(method.Func, router, endpoint, Index)
		}
	}

	return router
}
