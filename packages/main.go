package main

import (
	"fmt"
	"module/package02"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	package02.Write()

	err := checkmail.ValidateFormat("franciscoaap7@gmail.com")
	fmt.Println(err);
}