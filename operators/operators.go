package main

import "fmt"

func main() {
	var num01 int16 = 10;
	var num02 int32 = 30;
	// cannot sum different types without converting
	sum := int(num01) + int(num02)
	fmt.Println(sum)

	// go does not have ternary operators, should use if/else
	// text := sum > 5 ? "text 01" : "text 02"
	var text string
	if sum > 5 {
		text = "text 01"
	} else {
		text = "text 02"
	}
	fmt.Println(text)
}