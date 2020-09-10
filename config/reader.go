package config

import (
	"bytes"
	"fmt"
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
		fmt.Printf("File '%s'does not exist\n", filepath)
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file '%s': %s \n", filepath, err)
		os.Exit(1)
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
		fmt.Printf("Error unmarshling file: %s \n", err)
		os.Exit(1)
	}
	return spec
}
