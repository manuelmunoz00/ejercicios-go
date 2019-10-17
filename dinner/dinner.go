package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numeroHilos = 20

func main() {
	defer delicioso()
	invitados := [4]string{"Alice", "John", "Steven", "Arthur"}
	platillos := make(chan string, 3)
	for i := 0; i < len(invitados); i++ {
		go longanizas(invitados[i], platillos)
		go morcillas(invitados[i], platillos)
		go chunchules(invitados[i], platillos)
	}
	count := 0
	for elem := range platillos {
		count++
		fmt.Println(elem)
	}
	close(platillos)
}

func longanizas(persona string, p chan string) {
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", persona, " disfruta longanizas")
	p <- mensaje
}

func morcillas(persona string, p chan string) {
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", persona, " disfruta morcillas")
	p <- mensaje
}

func chunchules(persona string, p chan string) {
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", persona, " disfruta chunchules")
	p <- mensaje
}

func delicioso() {
	fmt.Println("Estuvo delicioso")
}
