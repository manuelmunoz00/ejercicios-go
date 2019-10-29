package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Tarjeta Estructura
type Tarjeta struct {
	numero string
}

func main() {
	var tc string
	fmt.Println("Welcome")
	fmt.Println("Ingrese los dígitos de la tarjeta de crédito o débito")
	fmt.Scanf("%s", &tc)
	nuevaTarjeta := new(Tarjeta)
	nuevaTarjeta.numero = tc
	fmt.Println(validar(*nuevaTarjeta))
}

func validar(tc Tarjeta) string {
	str := strings.Replace(tc.numero, "-", "", -1)
	strReversa := reversa(str)
	var eslice []int64
	for i, r := range strReversa {
		par := i + 1
		c := string(r)
		j, err := strconv.ParseInt(c, 10, 32)
		if err == nil {
			if par%2 == 0 {
				doble := j * 2
				if doble > 9 {
					doble = doble - 9
				}
				eslice = append(eslice, doble)
			} else {
				eslice = append(eslice, j)
			}
		}
	}
	var sumatoria int64
	for x := 0; x < len(eslice); x++ {
		sumatoria = sumatoria + eslice[x]
	}
	fmt.Println("sumatoria--", sumatoria)
	if sumatoria%10 == 0 {
		return "tc valida"
	}
	return "tc invalida"
}

func reversa(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		b.WriteByte(str[i])
	}
	return b.String()
}
