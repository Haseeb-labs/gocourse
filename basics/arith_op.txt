package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b int = 10, 4
	var result int

	result = a + b
	fmt.Println("Addition is : ", result)

	result = a - b
	fmt.Println("Subtraction is : ", result)

	result = a * b
	fmt.Println("Multiplication is : ", result)

	result = a / b
	fmt.Println("Divivsion is : ", result)
	//its same as c++ for int,int division
	//to get exact answer anyone of the two numbers to be a floating point number

	result = a % b
	fmt.Println("modulo is : ", result)

	//execution starts form right to left
	const PI float64 = 22 / 7.0
	fmt.Println("pi value is : ", PI)

	//overflow :- occurs when result of computational value exceeds the max capacity of given type
	//underflow :- smaller than minimum capacity

	//overflow with signed integers
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt += 1
	fmt.Println(maxInt)

	//overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt += 1
	fmt.Println(uMaxInt)

	//lets try underflow for floating
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
