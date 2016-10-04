package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cleonty/jsonutils"
)

func main() {
	args := os.Args
	filename := ""
	if len(args) > 1 {
		filename = args[1]
		content, err := jsonutils.MinifyFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

}
