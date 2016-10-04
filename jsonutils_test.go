// winutils_test.go
package jsonutils

import (
	"io/ioutil"
	"os"
	"testing"
)

var content string = `
{ "hello"
 :
 "world"}
`
var minifiedContent string = `{"hello":"world"}`
var prettyContent string = `{
  "hello": "world"
}`
var testFile = "test.json"

func TestPrettify(t *testing.T) {
	err := ioutil.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(testFile)
	pretty, err := PrettifyFile(testFile)
	if err != nil {
		t.Error(err)
		return
	}
	if pretty != prettyContent {
		t.Errorf("Expected \n%s\n got \n%s\n", prettyContent, pretty)
	}
}
func TestMinify(t *testing.T) {
	err := ioutil.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(testFile)
	minified, err := MinifyFile(testFile)
	if err != nil {
		t.Error(err)
		return
	}
	if minified != minifiedContent {
		t.Errorf("Expected %s got %s\n", minifiedContent, minified)
	}
}
