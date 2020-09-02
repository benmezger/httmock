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
