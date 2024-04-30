package main

import (
	"errors"
	"fmt"
	"math/big"
)

func main() {
	
	// int values
	var (
		num8 int8 = 22
		num16 int16 = 22222
		num32 int32 = 222222222
		num64 int64 = 2222222222222222222
	)
	fmt.Println(num8, num16, num32, num64)

	// unsigned and positive numbers
	var positiveNumber uint8 = 1
	var negativeNumber int8 = -1
	fmt.Println(positiveNumber, negativeNumber)

	// int default by system
	var anyNumber int = 1000000000000000000
	fmt.Println(anyNumber)

	// big int
	var bigInt big.Int
	bigInt.Exp(big.NewInt(10), big.NewInt(99), nil)

	newBigIntFromInt := big.NewInt(int64(10000))

	fmt.Println(newBigIntFromInt)

	// To convert a big integer to uint64, run this code:
	var smallNum, _ = new(big.Int).SetString("2188824200011112223", 10)
	num := smallNum.Uint64()
	fmt.Println(num)

	// Converting from a string
	var bigNum, ok = new(big.Int).SetString("218882428714186575617", 0)
	fmt.Println(bigNum, ok)

	// Aliases
	var numInt32 rune = 123456
	var numUint8 byte = 123
	fmt.Println(numInt32, numUint8)
	
	// float values
	var numFloat32 float32 = 123.45
	var numFloat64 float64 = 123456789.45
	numFloat := 123456789.45
	fmt.Println(numFloat32, numFloat64, numFloat)

	// string values
	var str string = "Text"
	str02 := "Text02"
	fmt.Println(str, str02)

	// asci numbers -> when using single quotes
	char := 'B' // result: 66
	fmt.Println(char)

	// initial values
	var someText string
	var someNumber int8
	fmt.Println(someText, someNumber)

	// boolean values
	var booleanValue bool
	fmt.Println(booleanValue)

	// error value
	var err error
	fmt.Println(err)
	err = errors.New("Internal error")
	fmt.Println(err)
}