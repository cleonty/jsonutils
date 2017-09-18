// Package jsonutils.
package jsonutils

import (
	"encoding/json"
	"io/ioutil"
)

// MinifyFile reads the file named filename,
// strips whitespace and newline characters,
// and returns its content as a string
func MinifyFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var f interface{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		return "", err
	}
	minifiedContent, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(minifiedContent), nil
}

// PrettifyFile reads the file named filename,
// indents with two spaces,
// and returns its content as a string
func PrettifyFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var f interface{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		return "", err
	}
	prettyContent, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyContent), nil
}
