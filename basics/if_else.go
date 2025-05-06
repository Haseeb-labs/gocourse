package main

import "fmt"

func main() {

	age := 25

	if age >= 18 {
		fmt.Println("you can watch")
	} else {
		fmt.Println("go and study")
	}

	temp := 25

	if temp >= 30 {
		fmt.Println("hot")
	} else {
		fmt.Println("not hot")
	}

	grade := 85

	if grade >= 90 {
		fmt.Println("A grade saruku")
	} else if grade >= 80 {
		fmt.Println("B grade saruku")
	} else if grade >= 70 {
		fmt.Println("C grade saruku")
	} else {
		fmt.Println("inkenduku cheppu")
	}

	num := 15
	if(num%3==0){
       if(num%2==0){
		  fmt.Println("divisible by 3 and 2")
	   } else{
		  fmt.Println("divisible by 3 but not 2")
	   }
	}else{
		fmt.Println("not divisible by 3")
	}
}
