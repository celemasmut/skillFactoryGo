package main

import "fmt"

var estado bool
var numero, numSecreto int

func main() {
	//estado = true
	/*if estado = true; estado == true {
		fmt.Println("estado es: ", estado)
	} else if miBool := true; miBool == false {
		fmt.Println(miBool)
	} else {
		fmt.Println("estado es:", miBool)
	}
	//numero = 4
	switch numero = 1; numero {
	case 1:
		fmt.Println("vale 1")
	case 2:
		fmt.Println("vale 2")
	default:
		fmt.Println("vale mas que 2")
	}*/

	//---> sentencia For
	/*	for i := 1; i <= 7; i++ {
			fmt.Println(i)
		}
	*/
	numSecreto = 9
	var primVez bool = true
	for {

		fmt.Println("ingrese un numero:")
		fmt.Scanln(&numero)
		if numero == numSecreto {
			fmt.Println("Adivinaste!!")
			if primVez {
				primVez = false
				continue
			} else {
				break
			}
		} else {
			fmt.Println("No adivinaste nada")
		}
	}
}
