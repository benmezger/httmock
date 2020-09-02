package main

import (
	"net/http"
)

type HTTPSpecMethodResponse struct {
	Status  int
	Payload string
}

type HTTPSpecMethodRequest struct {
	Args string
	Body string
}

type HTTPSpecMethod struct {
	Request  HTTPSpecMethodRequest
	Response HTTPSpecMethodResponse
	Handler  *http.Handler `yaml:"-"`
}

type HTTPSpec struct {
	Paths map[string]map[string]HTTPSpecMethod
}

func (s *HTTPSpec) GetPaths() []string {
	paths := make([]string, len(s.Paths))
	for k := range s.Paths {
		paths = append(paths, k)
	}
	return paths
}

func (s *HTTPSpec) GetPathHandlerByMethod(path, method string) interface{} {
	val, err := s.Paths[path][method]
	if err {
		return nil
	}
	return val.Handler

}
