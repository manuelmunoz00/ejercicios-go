package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("prueba.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Content file:", string(data))
}
