package main

import (
	"reflect"
	"strings"
)

type HTTPSpecMethodResponse struct {
	Status   int
	Payload  string
	Mimetype string
	Header   map[string]string
}

type HTTPSpecMethodRequest struct {
	Params map[string]string
	Body   string
}

type HTTPSpecMethod struct {
	Request  HTTPSpecMethodRequest
	Response HTTPSpecMethodResponse
	Handler  interface{} `yaml:"-"`
}

type UrlSpec map[string]*HTTPSpecMethod
type UrlPath map[string]UrlSpec

type HTTPSpec struct {
	Paths UrlPath
}

func (s *HTTPSpec) GetPaths() []string {
	paths := make([]string, 0)
	for k := range s.Paths {
		paths = append(paths, k)
	}
	return paths
}

func (s *HTTPSpec) GetPathHandlerByMethod(path, method string) interface{} {
	val, err := s.Paths[path][strings.ToLower(method)]
	if !err {
		return nil
	}
	return val.Handler
}

func getTypeMethod(obj interface{}, name string) interface{} {
	val := reflect.TypeOf(obj)
	if method, exists := val.MethodByName(name); !exists {
		return nil
	} else {
		return method
	}
}

func (hs *HTTPSpecMethod) SetHandler(handler interface{}, method string) {
	hs.Handler = getTypeMethod(handler, strings.ToUpper(method))
}

func (hs *HTTPSpecMethod) Invoke(args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	hs.Handler.(reflect.Method).Func.Call(inputs)
}
