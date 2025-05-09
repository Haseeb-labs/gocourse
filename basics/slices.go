package main

import (
	"fmt"
	"slices"
)

func main() {

	// //slices are like dynamic arrays thats why they can declared without mentioning size initially

	// //declaration:- var sliceName[]elementType

	// var number []int  //this tells it contains series of numbers

	// var number1 [] int {1,2,3,4}

	// number2 := []{7,8,9}

	// slice := make([]int,5) // starting with size 5

	arr := [5]int{1, 2, 3, 4, 5}

	newSlice := arr[1:4]

	fmt.Println(newSlice) //4 is not included

	newSlice = append(newSlice,4,5,6,7,8)

	fmt.Println("new slice is:",newSlice)

	SliceCopy := make([]int,len(newSlice))

	copy(SliceCopy,newSlice)

	fmt.Println("copy of new slice:",SliceCopy)

	for i,v := range newSlice{
		fmt.Println(i,v)
	}

	if slices.Equal(newSlice,SliceCopy){
		fmt.Println("both slices are equal ")
	}

}
