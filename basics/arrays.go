package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 10
	numbers[0] = 1
	fmt.Println(numbers)

	fruits := [4]string{"apple", "banana", "mango", "orange"}
	fmt.Println(fruits)
	fruits[0] = "grapes"
	fmt.Println(fruits)

	org_array := [5]int{1, 2, 3, 4, 5}
	copied_array := org_array
	copied_array[0] = 100
	fmt.Println("org_array:", org_array)
	fmt.Println("copied_array:", copied_array)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("element at index, ", i, ":", numbers[i])
	}

	for ind, value := range numbers {
		fmt.Printf("Index: %d , value: %d\n", ind, value)
	}

	for _, val := range numbers {
		fmt.Println("squares of numbers is:", val*val)
	}

	a, b := somefunction()
	fmt.Println(a)
	fmt.Println(b)

	c, _ := somefunction()
	fmt.Println(c)

	arr1 := [3]int{1,2,4}
	arr2 := [3]int{1,2,4}

	fmt.Println("equal or not:",arr1==arr2)

	var matrix [3][3]int = [3][3]int{
		{1,2,4},
		{1,3,9},
		{1,4,16},
	}
	fmt.Println(matrix)

	originalArray := [3]int{2,4,16}
	var copiedArray *[3]int
	copiedArray = &originalArray
	copiedArray[0] = 100
	fmt.Println("new original array:",originalArray)

}



func somefunction() (int, int) {
	return 1, 2
}
