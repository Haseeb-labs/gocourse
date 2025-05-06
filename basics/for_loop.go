package main

import "fmt"

func main() {

	var num int = 5

	//iterate over a range
	for i := 1; i <= num; i++ {
		fmt.Println("hello golang")
	}

	//iterate over a collection
	numbers := []int {1,2,3,4,5,6}
	for index,value := range numbers{
		fmt.Printf("Index : %d,Value : %d\n",index,value)
	}

	//break and continue statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd detected : ", i)
		if i == 5 {
			break
		}
	}

	//star pattern
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := range 10 {
		fmt.Println(10 - i)
	}

}
