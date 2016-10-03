package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	filename := ""
	if len(args) > 1 {
		filename = args[1]
	} else {
		return
	}
	inplace := true
	var content []byte
	in, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err = ioutil.ReadAll(in)
	if err != nil {
		fmt.Println(err)
		in.Close()
		return
	}
	in.Close()
	var f interface{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	var writer io.Writer
	if inplace {
		fp, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fp.Close()
		writer = fp
	} else {
		writer = os.Stdout
	}
	writer.Write(out)
}
