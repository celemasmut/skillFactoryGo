package main

import "fmt"

func main() {

	fmt.Println(reverseString("Hello"))
	fmt.Println(reverseString("chocolate"))
	fmt.Println(reverseString("coffe"))
	fmt.Println(reverseString("Skill-Factory"))
	fmt.Println(reverseString("golang"))

}

func reverseString(origin string) string {
	copy := ""
	for i := range origin {
		copy += string(origin[len(origin)-1-i])
	}
	return copy
}
