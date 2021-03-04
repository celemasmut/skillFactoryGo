package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	archivo := "prueba.txt"         //archivo es la variable de mi archivo
	miArch, err := os.Open(archivo) //miArch es el puntero del archivo

	defer miArch.Close() //el defer se ejecuta antes de que termine el main o se corte el main

	if err != nil {
		fmt.Println("Error abriendo archivo")
		//os.Exit(1) // el numero es para enumerar los exit()
	}

	//PANIC
	ejemploPanic()
	//como el defer se ejecuta si o si cuando termina
	//se usa el Recover para volverr a tomar el control luego de un Panic
	ejemplo2Panic2()
}

//la funcion panic fuerza un error. muestra el error y aborta el programa
func ejemploPanic() {
	a := 3
	if a == 1 {
		panic("Se encontro el valor de 1")
	} else {
		fmt.Println("No se entro al Panic")
	}
}

func ejemplo2Panic2() {

	//abrir el archivo webserver.log para escritura
	arch, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//y cerrar cuando termine la funcion main
	defer arch.Close()
	//asociar el manejador del archivo al log
	log.SetOutput(arch)
	//agregar 10 lineas al archivo log
	for i := 0; i < 10; i++ {
		log.Printf("Error en linea %v", i) //esta salida es por el archivo y no por pantalla
	}

	//funcion anonima para el defer
	/*defer func() {
		recov := recover()
		if recov != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n%v", recov)
		}
	}()*/
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
