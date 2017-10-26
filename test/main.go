package main

import (
	"fmt"
	"io/ioutil"

	"ireul.com/binfs"
)

func main() {
	f, err := binfs.Open("/other/other2/hello.txt")
	if err != nil {
		panic(err)
	}
	s, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
