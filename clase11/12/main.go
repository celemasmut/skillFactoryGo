package main

import "fmt"

var number1, number2 int

//variable de tipo funcion
var suma func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	/*fmt.Println("insert number1:")
	fmt.Scanln(&number1)
	fmt.Println("insert number2:")
	fmt.Scanln(&number2)
	fmt.Printf("Sum: %d", number1+number2)*/

	/*num, estado := dos(7)
	fmt.Println(uno(5))
	fmt.Println(num)
	fmt.Println(estado)*/

	fmt.Printf("sumo 6+3 = %d", suma(6, 3))
	//aca se redifine la variable con otra funcion anonima..
	suma = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("\nresto 6-3 = %d", suma(6, 3))

	//fmt.Println(calculo(34, 87, 77))

	//otra funcion anonima
	operaciones()

	//closures
	tabladel := 2
	miTabla := Tabla(tabladel) // miTabla es de tipo funcion
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}

}

//otra funcion anonima
func operaciones() {
	resultado := func() int {
		var a int = 34
		var b int = 27
		return a + b
	}
	fmt.Println(resultado())
}

//Tabla es una funcion de closure
func Tabla(valor int) func() int {
	numero := valor //la primera ejecucion va a ejecutar linea57 y 58
	secuencia := 0
	return func() int { //cuando se llama y se asigna a una variable, la variable solo toma la parte que se retorna.
		secuencia++
		return numero * secuencia
	}

}

//lo que esta fuera de los () va el tipo que se va a retonar
func uno(numero int) int {
	return numero * 2
}

//Cuando se devuelve mas de un parametro se ponen entre () devuelve entero y booleano
func dos(numero int) (int, bool) {
	if numero == 5 {
		return 5, true
	}
	return 10, false
}

//el ... indica que no se sabe cuantas variables por parametros van a pasar.
func calculo(numero ...int) int {
	total := 0
	//usaremos for con range
	//range devuelve dos valores   pos/valor. pero si no quiero usar uno de esos valores se le pone _ guion bajo.
	for pos, valor := range numero {
		total = +valor
		fmt.Printf("\nPos [%d]\n", pos)
	}
	return total
}
