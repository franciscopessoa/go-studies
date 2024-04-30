package main

import "fmt"

// basic function
func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// returning more than one result
func calc(n1, n2 int8) (int8, int8) {
	sumResult := n1 + n2
	subResult := n1 - n2
	return sumResult, subResult

}

func main() {
	sumResult := sum(10, 20)
	fmt.Println(sumResult)

	// creating a function inside another function
	var f = func(text string) {
		fmt.Println(text)
	}
	f("something")

	sum01, sub01 := calc(20,10)
	fmt.Println(sum01, sub01)

	// ignoring some return value
	sum02, _ := calc(20,10)
	fmt.Println(sum02)

}