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
	fmt.Println("Ingrese los dígitos de la tarjeta de crédito o débito")
	fmt.Scanf("%s", &tc)
	nuevaTarjeta := new(Tarjeta)
	nuevaTarjeta.numero = tc
	fmt.Println(validar(*nuevaTarjeta))
}

func validar(tc Tarjeta) bool {
	// se reemplazan guiones por vacios
	str := strings.Replace(tc.numero, "-", "", -1)
	// se invierte string de números dado que la validación es de izquierda a derecha
	strReversa := reversa(str)
	var eslice []int64
	for i, r := range strReversa {
		par := i + 1
		c := string(r)
		// se convierte a entero el string de números
		j, err := strconv.ParseInt(c, 10, 32)
		if err == nil {
			// por cada 2 dígitos se duplica el valor
			if par%2 == 0 {
				doble := j * 2
				// si el doble es mayor a 9 se resta al doble el valor de 9
				if doble > 9 {
					doble = doble - 9
				}
				// se agrega al slice el doble
				eslice = append(eslice, doble)
			} else {
				// se agrega el valor de los índices que no sean pares
				eslice = append(eslice, j)
			}
		}
	}
	// se recorre el slice para hacer la sumatoria según la evaluación de los dígitos realizado anteriormente
	var sumatoria int64
	for x := 0; x < len(eslice); x++ {
		sumatoria = sumatoria + eslice[x]
	}
	// si la sumatoria es divisible por 10 entonces el número es válido
	if sumatoria%10 == 0 {
		return true
	}
	return false
}

func reversa(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		b.WriteByte(str[i])
	}
	return b.String()
}
