package main

import "fmt"

func main() {
	potencia(2)
}

func potencia(num int) {
	if num > 10000 { // salida corte.
		return
	}
	fmt.Println(num)
	potencia(num * 3) //llamado recursivo

}
