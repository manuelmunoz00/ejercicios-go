package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cambiarString("Igual a tudes"))
}

func cambiarString(palabra string) string {
	return strings.Replace(palabra, "u", "o", 1)
}
