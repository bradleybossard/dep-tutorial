package main

import "fmt"
import "io/ioutil"
import "github.com/pkg/errors"

func readFile() ([]byte, error) {
	dat, err := ioutil.ReadFile("./test.tmp")
	if err != nil {
		fmt.Println("Unwrapped = ", err)
		return dat, errors.Wrap(err, "read failed")
	}
	return dat, err
}

func main() {
	_, err := readFile()
	if err != nil {
		fmt.Println("Wraped = ", err)
	}
}
