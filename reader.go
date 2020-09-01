package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func fileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func OpenFile(filepath string) io.Reader {
	if !fileExists(filepath) {
		panic("File does not exist")
	}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(file)
}

func ReadHTTPSpec(stream io.Reader) *HTTPSpec {
	spec := new(HTTPSpec)

	buff := new(bytes.Buffer)
	buff.ReadFrom(stream)

	/* TODO: currently we read the whole file in memory.
	 * Any other option? */
	if err := yaml.Unmarshal(buff.Bytes(), spec); err != nil {
		panic(err)
	}
	return spec
}
