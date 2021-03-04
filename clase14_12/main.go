package main

import "fmt"

var array [7]int
var matriz [5][6]int
var mySlice []int // un arreglo pero sin definir su dimension

func main() {
	array[3] = 19
	fmt.Println(array)

	miVector := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(miVector)
	fmt.Println("contenido de mi vector")
	for i := 0; 1 < 10; i++ {
		fmt.Printf("mi vector [%d] = %d", i, miVector[i])
	}
	matriz[2][4] = 12
	fmt.Println(matriz)
	mySlice = []int{3, 6, 9}
	fmt.Println(mySlice)

	mySlice = append(mySlice, 8)
	fmt.Println(mySlice)

	//	variante2()
	//variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{2, 5, 6, 8, 9}
	porcion := elementos[2:5] //toma desde la pos 2 hasta la 4, ya que no es inclusiveal 5
	fmt.Println(elementos)
	fmt.Println(porcion)
	porcion = elementos[1:]
	fmt.Println(elementos)

}

func variante3() {
	elementos := make([]int, 6, 20) // creo un array. make tiene dos parametros. el 2° da la dimencion el 3°es para que me reserve por las dudas .
	fmt.Println(elementos)
	fmt.Printf("largo %d  Capacidad : %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 8; i++ {
		nums = append(nums, i)
		fmt.Println(nums) // me lo imprime con los corchetes.
	}
	fmt.Printf("largo:%d  Capacidad: %d", len(nums), cap(nums))
}
