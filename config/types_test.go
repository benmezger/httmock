package config

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type valid struct{}

func (v *valid) Method() { return }

func Test_HTTPSpecMethods(t *testing.T) {
	t.Run("Test if GetPaths returns a correct slice", func(t *testing.T) {
		spec := ReadHTTPSpec(strings.NewReader(example))
		if assert.NotNil(t, spec) {
			assert.Equal(t, spec.GetPaths(), append(make([]string, 0), "/"))
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
