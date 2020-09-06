package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = `
paths:
  /:
    get:
      request:
        args: foobar=12
      response:
        status: 200
        payload: >-
          {"msg": "Hello, world"}
    post:
      request:
        args: foobar=12
        body: >-
          {"msg": "request payload"}
      response:
        status: 201
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
	get_request := &HTTPSpecMethodRequest{"foobar=12", ""}
	get_response := &HTTPSpecMethodResponse{200, `{"msg": "Hello, world"}`}
	post_request := &HTTPSpecMethodRequest{"foobar=12", `{"msg": "request payload"}`}
	post_response := &HTTPSpecMethodResponse{201, `{"msg": "created"}`}

	expected := &HTTPSpec{UrlPath{
		"/": {
			"get":  &HTTPSpecMethod{*get_request, *get_response, nil},
			"post": &HTTPSpecMethod{*post_request, *post_response, nil}}}}

	spec := ReadHTTPSpec(strings.NewReader(example))
	assert.Equal(t, spec, expected, "Spec should be equal to expected")
}
