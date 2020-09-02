package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_HTTPSpecMethods(t *testing.T) {
	t.Run("Test if GetPaths returns a correct slice", func(t *testing.T) {
		spec := ReadHTTPSpec(strings.NewReader(example))
		if assert.NotNil(t, spec) {
			assert.Equal(t, spec.GetPaths(), append(make([]string, 1), "/"))
		}
	})

	t.Run("Test if GetPathHandlerByMethod returns nil ", func(t *testing.T) {
		spec := ReadHTTPSpec(strings.NewReader(example))
		if assert.NotNil(t, spec) {
			assert.Nil(t, spec.GetPathHandlerByMethod("/", "GET"))
		}
	})

	t.Run("Test if GetPathHandlerByMethod with invalid path", func(t *testing.T) {
		spec := ReadHTTPSpec(strings.NewReader(example))
		if assert.NotNil(t, spec) {
			assert.Nil(t, spec.GetPathHandlerByMethod("invalid-path/", "GET"))
		}
	})

	t.Run("Test if GetPathHandlerByMethod with invalid path method", func(t *testing.T) {
		spec := ReadHTTPSpec(strings.NewReader(example))
		if assert.NotNil(t, spec) {
			assert.Nil(t, spec.GetPathHandlerByMethod("/", "OPTION"))
		}
	})
}
