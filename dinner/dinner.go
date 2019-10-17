package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// const numeroHilos = 20

func main() {
	fmt.Println("Welcome to the dinner")
	defer delicioso()
	morsels := []string{"Chorizos", "Longanizas", "Chunchules", "Morcillas", "Pamplonas"}
	invitados := make(chan string, 50)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		rand.Seed(time.Now().UnixNano())
		go alice(morsels[rand.Intn(len(morsels))], invitados, &wg)
		go arthur(morsels[rand.Intn(len(morsels))], invitados, &wg)
		go steven(morsels[rand.Intn(len(morsels))], invitados, &wg)
		go john(morsels[rand.Intn(len(morsels))], invitados, &wg)
	}

	// wait for all goroutines to end
	wg.Wait()
	// close the channel so that we not longer expect writes to it
	close(invitados)

	count := 0
	for elem := range invitados {
		count++
		fmt.Println(elem)
	}

}

func alice(plato string, p chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", "Alice disfruta ", plato)
	p <- mensaje
}

func arthur(plato string, p chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", "Arthur disfruta ", plato)
	p <- mensaje
}

func steven(plato string, p chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", "Steven disfruta ", plato)
	p <- mensaje
}

func john(plato string, p chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	min := 300
	max := 800
	tiempoProcesamiento := rand.Intn(max-min) + min
	time.Sleep(time.Duration(tiempoProcesamiento) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%s", "John disfruta ", plato)
	p <- mensaje
}

func delicioso() {
	fmt.Println("Estuvo delicioso")
}
