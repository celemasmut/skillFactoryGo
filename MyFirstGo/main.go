package main

import "fmt"

func main() {
	//useing var
	var name = /*"string" can be here but is not necessary because it's inferred with the ""*/ "Brsd"
	var age int32 = 37
	var isCool = true
	lastName := "Cunsolo" //another way to declarate a variable without the word var
	size := 1.2
	fmt.Println(name, lastName, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
