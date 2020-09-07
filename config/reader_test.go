package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_PrivateFileExists(t *testing.T) {
	t.Run("Test if fileExists returns false", func(t *testing.T) {
		if res := fileExists("non-existent-file"); res != false {
			t.Errorf("File exists should have returned true, got %v", res)
		}
	})
}

func Test_ReadHTTPSpec(t *testing.T) {
	get_request := &HTTPSpecMethodRequest{map[string]string{"name": "get-name-param"}, `{"msg": "Body of GET request"}`}
	get_response := &HTTPSpecMethodResponse{
		200, `{"msg": "Hello, world"}`, "application/json",
		map[string]string{"Content-Type": "application/json"},
	}
	post_request := &HTTPSpecMethodRequest{map[string]string{"name": "post-name-param"}, `{"msg": "request payload"}`}
	post_response := &HTTPSpecMethodResponse{
		201, `{"msg": "created"}`, "application/json",
		map[string]string{"Content-Type": "application/json"},
	}

	expected := &HTTPSpec{UrlPath{
		"/": {
			"get":  &HTTPSpecMethod{*get_request, *get_response, nil},
			"post": &HTTPSpecMethod{*post_request, *post_response, nil},
		},
	}}

	spec := ReadHTTPSpec(strings.NewReader(example))
	assert.Equal(t, spec, expected, "Spec should be equal to expected")
}
