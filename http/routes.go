package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/benmezger/httmock/config"

	"github.com/julienschmidt/httprouter"
)

func GenerateRoutes(spec *config.HTTPSpec, handler *httprouter.Router) {
	for _, methods := range spec.Paths {
		for name, method := range methods {
			method.SetHandler(handler, strings.ToUpper(name))
		}
	}
}

func GenerateHandler(method *config.HTTPSpecMethod) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		params := r.URL.Query()
		for k, v := range method.Request.Params {
			if params.Get(k) != v {
				http.Error(w,
					fmt.Sprintf(`{"msg":  "Missing param '%s' with content '%s'"}`, k, v),
					http.StatusNotFound)
				return
			}
		}

		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != method.Request.Body {
			http.Error(w,
				fmt.Sprintf("'%s' in request does not match expected '%s'", string(body), method.Request.Body),
				http.StatusBadRequest)
			return
		}

		for k, v := range method.Response.Header {
			w.Header().Add(k, v)
		}

		w.WriteHeader(method.Response.Status)
		w.Write([]byte(method.Response.Payload))
	}
}

func SetupRoutes(spec *config.HTTPSpec) *httprouter.Router {
	router := httprouter.New()
	GenerateRoutes(spec, router)
	for path, attrs := range spec.Paths {
		for _, m := range attrs {
			m.Invoke(router, path, GenerateHandler(m))
		}
	}
	return router
}
