package main

import "fmt"

//los middleware son como muñecas rusas que nos permiten encapsular
//la ejecucion de una funcion ya prexistente con otra funcion nueva

func main() {
	/*middware son interceptores que permiten ejecutar instruccines comnes a varias funciones.
	se utiliza a funciones con mismos parametros y devuelven lo mismos tipos de variables
	*/
	fmt.Println("inicio de programa")

	//
	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(4, 3)
	fmt.Println(result)

}
func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

//recibe una funcion con parametros y retorno
//f es el parametro de tipo funcion. a su vez se retorna una funcion
func operacionesMidd(f func(int, int) int) func(int, int) int {
	//retorna una funcion anonima
	return func(a, b int) int {
		//lo sig es lo que se  aplica previo a realizar todas las funciones
		fmt.Println("inicio de operacion")
		//retorno la funcion que pase por parametro.
		return f(a, b)
	}
}
