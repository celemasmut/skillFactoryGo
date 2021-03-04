package main

import "fmt"

func main() {

	fmt.Println(leapYear(2020))
	fmt.Println(leapYear(2300))
	fmt.Println(leapYear(2000))
	fmt.Println(leapYear(1996))
	fmt.Println(leapYear(2400))
	fmt.Println(leapYear(2007))
	fmt.Println(leapYear(2008))
	fmt.Println(leapYear(1916))
	fmt.Println(leapYear(2040))

}

func leapYear(year int16) bool {
	isLeap := false
	if (year % 400) == 0 {
		isLeap = true
	} else if (year%4) == 0 && (year%100) != 0 {
		isLeap = true
	}
	return isLeap
}
