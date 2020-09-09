package http

import (
	"fmt"
	"github.com/benmezger/httmock/config"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienschmidt/httprouter"
	"github.com/kinbiko/jsonassert"
)

type valid struct{}

func (v *valid) Method() { return }

var example = `
paths:
  /:
    get:
      request:
        params:
          name: get-name-param
        body: >-
          {"msg": "Body of GET request"}
      response:
        status: 200
        mimetype: application/json
        header:
          Content-Type: application/json
        payload: >-
          {"msg": "Hello, world"}
    post:
      request:
        params:
          name: post-name-param
        body: >-
          {"msg": "request payload"}
      response:
        header:
          Content-Type: application/json
        status: 201
        mimetype: application/json
        payload: >-
          {"msg": "created"}
`

func Test_GenerateRoutes(t *testing.T) {
	route := httprouter.New()
	spec := config.ReadHTTPSpec(strings.NewReader(example))
	GenerateRoutes(spec, route)
	assert.Equal(t, len(spec.Paths), 1, "There should be only one endpoint '/'")
	assert.Equal(t, len(spec.Paths["/"]), 2, "There should be one two methods 'post' and 'get'")
}

func Test_SetupRoutes(t *testing.T) {
	router := SetupRoutes(config.ReadHTTPSpec(strings.NewReader(example)))
	assert.NotNil(t, router)
}

func Test_GenerateHandler(t *testing.T) {
	spec := config.ReadHTTPSpec(strings.NewReader(example))
	route := SetupRoutes(spec)

	t.Run("Test '/' path without query params", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", strings.NewReader(""))

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status 404")

		ja := jsonassert.New(t)
		ja.Assertf(
			fmt.Sprintf(`{"msg":  "Missing param '%s' with content '%s'"}`, "name", "get-name-param"),
			rr.Body.String())
	})

	t.Run("Test '/' path with query params with missing body", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", strings.NewReader(""))

		params := url.Values{}
		params.Add("name", "get-name-param")
		req.URL.RawQuery = params.Encode()

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status 401")
		assert.Equal(t, rr.Body.String(),
			"'' in request does not match expected '{\"msg\": \"Body of GET request\"}'\n",
			"Request body does not match")
	})

	t.Run("Test '/' path with query params and correct body", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", strings.NewReader("{\"msg\": \"Body of GET request\"}"))

		params := url.Values{}
		params.Add("name", "get-name-param")
		req.URL.RawQuery = params.Encode()

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200")
		assert.Equal(t, rr.Body.String(), "{\"msg\": \"Hello, world\"}")
	})
}
