package main

import (
	"strings"
	"testing"
)

func Test_PrivateFileExists(t *testing.T) {
	t.Run("Test if fileExists returns false", func(t *testing.T) {
		if res := fileExists("non-existent-file"); res != false {
			t.Errorf("File exists should have returned true, got %v", res)
		}
	})
}

func Test_ReadHTTPSpec(t *testing.T) {
	t.Run("Test if ReadHTTPSpec reads the specification", func(t *testing.T) {
		spec := ReadHTTPSpec(strings.NewReader("version: 1"))
		if spec == nil {
			t.Errorf("Spec is nil, got %v", spec)
		}
		if spec.Version != 1 {
			t.Errorf("Spec.Version is not 1, got %v", spec)
		}
	})
}
