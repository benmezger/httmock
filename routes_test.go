package main

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kinbiko/jsonassert"
)

type valid struct{}

func (v *valid) Method() { return }

func Test_getTypeMethod(t *testing.T) {

	t.Run("getTypeMethod should return nil", func(t *testing.T) {
		type invalid struct{}
		assert.Equal(t, getTypeMethod(&invalid{}, "GET"), nil, "getTypeMethod should return nil")
	})

	t.Run("getTypeMethod should return func type", func(t *testing.T) {
		valid_t := &valid{}
		expected_name, err := reflect.TypeOf(valid_t).MethodByName("Method")
		if !err {
			t.Error(err)
		}

		assert.Equal(
			t,
			getTypeMethod(valid_t, "Method").(reflect.Method).Name,
			expected_name.Name,
			"getTypeMethod should return the ")

		assert.Equal(
			t,
			getTypeMethod(valid_t, "Method").(reflect.Method).Type,
			expected_name.Type,
			"getTypeMethod should return the ")
	})
}

func Test_GenerateRoutes(t *testing.T) {
	route := httprouter.New()
	spec := ReadHTTPSpec(strings.NewReader(example))
	GenerateRoutes(spec, route)
	assert.Equal(t, len(spec.Paths), 1, "There should be only one endpoint '/'")
	assert.Equal(t, len(spec.Paths["/"]), 2, "There should be one two methods 'post' and 'get'")
}

func Test_SetupRoutes(t *testing.T) {
	router := SetupRoutes(ReadHTTPSpec(strings.NewReader(example)))
	assert.NotNil(t, router)
}

func Test_GenerateHandler(t *testing.T) {
	spec := ReadHTTPSpec(strings.NewReader(example))
	route := SetupRoutes(spec)

	t.Run("Test '/' path without query params", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status 404")

		ja := jsonassert.New(t)
		ja.Assertf(
			fmt.Sprintf(`{"msg":  "Missing param '%s' with content '%s'"}`, "name", "get-name-param"),
			rr.Body.String())
	})

	t.Run("Test '/' path with query params", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		params := url.Values{}
		params.Add("name", "get-name-param")
		req.URL.RawQuery = params.Encode()

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200")
		ja := jsonassert.New(t)
		ja.Assertf(`{"msg": "Hello, world"}`, rr.Body.String())
	})

}
