package main

import (
	"fmt"
)

// Cuadrado type Struct
type Cuadrado struct {
	ancho int
	alto  int
}

func main() {
	// Struct cuadrado que recbia ancho y alto para calcular area y perímetro
	var ancho, alto int
	fmt.Println("Ingrese el ancho")
	fmt.Scanf("%d", &ancho)
	fmt.Println("Ingresse el alto")
	fmt.Scanf("%d", &alto)

	nuevoCuadrado := new(Cuadrado)
	nuevoCuadrado.alto = alto
	nuevoCuadrado.ancho = ancho
	fmt.Println(obtenerAreaCuadrado(*nuevoCuadrado))
	fmt.Println(obtenerPerimetroCuadrado(*nuevoCuadrado))

}

func obtenerAreaCuadrado(c Cuadrado) int {
	return c.ancho * c.alto
}

func obtenerPerimetroCuadrado(c Cuadrado) int {
	return c.ancho*2 + c.alto*2
}
