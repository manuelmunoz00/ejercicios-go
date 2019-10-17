package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numeroThreads = 25

func main() {
	defer salir()
	fmt.Println("Welcome to internet cafe")
	turista := make(chan string)
	// rand.Seed(time.Now().UnixNano())
	// min := 1
	// max := 25
	// computadora := make(chan int)
	// totalTuristas := make(chan int)
	// go obtenerStatus(turista)
	// turista <- 1
	// computadora <- 1
	// totalTuristas <- 25
	for i := 1; i <= numeroThreads; i++ {
		go online(i, turista)
		go waiting(i, turista)
		go done(i, turista)
	}

	count := 0
	for elem := range turista {
		count++
		fmt.Println(elem)
	}

}

func online(indice int, turista chan string) {
	var seconds int
	seconds = rand.Intn(100)
	time.Sleep(time.Duration(seconds) * time.Millisecond)
	mensaje := fmt.Sprintf("%s%d%s", "Turista ", indice, " online")
	turista <- mensaje
}

func waiting(indice int, turista chan string) {
	var seconds int
	seconds = rand.Intn(100)
	time.Sleep(time.Duration(seconds) * time.Second)
	mensaje := fmt.Sprintf("%s%d%s", "Turista ", indice, " en espera")
	turista <- mensaje
}

func done(indice int, turista chan string) {
	var seconds int
	seconds = rand.Intn(100)
	time.Sleep(time.Duration(seconds) * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 120
	tiempoConectado := rand.Intn(max-min) + min
	mensaje := fmt.Sprintf("%s%d%s%d", "Turista ", indice, " listo. tiempo conectado:", tiempoConectado)
	turista <- mensaje
}

func obtenerStatus(t chan int) {

	var nuevoTurista int
	turista := <-t
	fmt.Println("Turista", turista, "online")
	if turista < 25 {
		nuevoTurista = turista + 1
		fmt.Println("Turista", nuevoTurista, "en espera")
		go obtenerStatus(t)
		t <- nuevoTurista
	}

	time.Sleep(500 * time.Millisecond)

	if turista > 25 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 120
	tiempoConectado := rand.Intn(max-min) + min
	fmt.Println("Turista", turista, "está listo.Tiempo conectado", tiempoConectado, "minutos")
	t <- nuevoTurista

	// almacena el valor del canal t(variable int)
	/*
		turista := <-t
		if turista > 8 {
			fmt.Println("El turista", turista, "está esperando su turno")
		} else {
			fmt.Println("El turista ", turista, " está online")
			// declara y almacena el valor de turista y le suma 1
			newTuristaOnline := turista + 1
			// Se vuelve a llamar a obtenerStatus y se le asigna al canal t el valor de newTuristaOnline
			go obtenerStatus(t)
			t <- newTuristaOnline
			time.Sleep(5000 * time.Millisecond)
		}
	*/

	// turista := <-t
	// if turista > 25 {
	// 	return
	// }

	// if turista < 8 {

	// }

	// if turista > 8 {
	// 	// Turista que debe pasar a espera
	// 	newTurista := turista + 1
	// 	fmt.Println("Turista en espera", newTurista)
	// 	go obtenerStatus(t)
	// 	t <- newTurista

	// 	// Debe pasar a ejecución
	// } else {
	// 	fmt.Println("Turista online", turista)
	// 	// Turista conectado
	// 	turistaOnline := turista + 1
	// 	go obtenerStatus(t)
	// 	t <- turistaOnline

	// 	// Turista conectado que debe pasar a desconectado
	// 	rand.Seed(time.Now().UnixNano())
	// 	min := 1
	// 	max := 25
	// 	tiempoConectado := rand.Intn(max-min) + min
	// 	fmt.Println("Turista", turista, "está listo.Tiempo conectado", tiempoConectado, "minutos")
	// }

	// computadora := <-c
	// totalTuristas := <-tt
	// _, found := find(conectados, turista)

	// if totalTuristas < 26 {
	// 	if computadora < 9 && !found {
	// 		fmt.Println("el turista", turista, "está online")
	// 		conectados = append(conectados, turista)
	// 		rand.Seed(time.Now().UnixNano())
	// 		// tiempoConectado := calcularTiempoSesion()
	// 		tiempoSesion = append(tiempoSesion, calcularTiempoSesion())
	// 		min := 1
	// 		max := 25
	// 		nuevoTurista := rand.Intn(max-min) + min
	// 		nuevaComputadora := computadora + 1
	// 		quitarTurista := totalTuristas - 1
	// 		go obtenerStatus(t, c, tt)
	// 		t <- nuevoTurista
	// 		c <- nuevaComputadora
	// 		tt <- quitarTurista
	// 		time.Sleep(5000 * time.Millisecond)
	// 	} else {
	// 		fmt.Println("el turista", turista, "está a la espera")
	// 		conectados = append(conectados, turista)
	// 		rand.Seed(time.Now().UnixNano())
	// 		min := 1
	// 		max := 25
	// 		nuevoTurista := rand.Intn(max-min) + min
	// 		quitarTurista := totalTuristas - 1
	// 		fmt.Println("turista--", turista)
	// 		fmt.Println("quitarTurista--", quitarTurista)
	// 		go obtenerStatus(t, c, tt)
	// 		t <- nuevoTurista
	// 		tt <- quitarTurista
	// 	}
	// }

	/*rand.Seed(time.Now().UnixNano())
	min := 15
	max := 120
	minutosOnline := rand.Intn(max-min) + min
	fmt.Println("el turista", turista, "ha pasado", minutosOnline, "online")
	fmt.Println("el turista terminó la sesión")*/

}

func calcularTiempoSesion() int {
	// Devuelve el tiempo de sesión conectado de un turista
	rand.Seed(time.Now().UnixNano())
	min := 15
	max := 120
	return rand.Intn(max-min) + min
}

func find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func salir() {
	fmt.Println("The place is empty, let's close up and go to the beach!")
}
