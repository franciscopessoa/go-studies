package main

import "fmt"

func main() {
	// creating vars
	var variable01 string = "var 01"
	variable02 := "var 02"

	fmt.Println(variable01)
	fmt.Println(variable02)

	var (
		variable03 string = "var 03"
		variable04 string = "var 04"
	)
	fmt.Println(variable03, variable04)

	variable05, variable06 := "var05", "var06"
	fmt.Println(variable05, variable06)

	// inverting values
	variable05, variable06 = variable06, variable05
	fmt.Println(variable05, variable06)

	// creating constants
	const constant01 string = "const 01"
	const constant02  = "const 02"
	fmt.Println(constant01, constant02)

}