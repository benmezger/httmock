package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type APIHandle func(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) error

func AppHandle(fn APIHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fn(w, r, p)
	}
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Here")
}
