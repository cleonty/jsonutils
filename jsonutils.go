// jsonutils project jsonutils.go
package jsonutils

import (
	"encoding/json"
	"io/ioutil"
)

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
