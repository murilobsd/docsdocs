package util

import (
	"io/ioutil"
)

// SaveFile this function save content to file
func SaveFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

// ReadFile this function read content
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
