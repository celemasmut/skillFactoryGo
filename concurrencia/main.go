package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Celeste")
	fmt.Println("estoy aqui")
	fmt.Println("estoy aquix2")

	var x string
	fmt.Scanln(&x)
	miNombreLento("Luciano")

}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		//para que tarde un milisegundo por mil
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
