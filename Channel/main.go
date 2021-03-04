package main

import (
	"fmt"
	"time"
)

func main() {
	//chan es channel. time.duration es para calcular el tiempo de ejecucion
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("llegue hasta a aqui")
	//aca mg sera el resultado de escuchar canal1
	msg := <-canal1
	fmt.Println(msg)
}

func bucle(micanal chan time.Duration) {
	inicio := time.Now() //graba el tiempo del momento en que se entra a la rutina
	for i := 0; i < 100000000000; i++ {
	}
	final := time.Now()
	//no hay return pero con <- es como asignarle valor a la variable.
	//se realiza una resta para devolver la duracion
	//micanal esta escuchando lo que retorna final.Sub
	micanal <- final.Sub(inicio)
}
