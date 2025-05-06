package main

import "fmt"

func main() {

	fruit := "apple"
    //in golang we dont need break after every case but 
	// we have "fallthrough" in case if we want to execute continuos cases
	//check at switch write variable there otherwise write at case
    switch fruit {
	case "apple":
		fmt.Println("its a red fruit")
	case "banana":
		fmt.Println("yellow one")
	default:
		fmt.Println("invalid one")
	}

	day := "Monday"

	switch day{
	case "Monday","Tuesday","Wednesday","Thursday","Friday":
		fmt.Println("its a week day , go to office")

	case "Saturday","Sunday":
		fmt.Println("NETFLIX")

	default:
		fmt.Println("Invalid day")
	}

	num := 15

	switch {
	case num < 10 :
		fmt.Println("less than  10")
	case num>=10 && num<=20:
		fmt.Println("between 10 and 20")
	default:
		fmt.Println("greater than or equal to 20")
	}
    
	number := 2
	switch{
	case number>=1:
		fmt.Println("greater than 1")
		fallthrough
	case number==2:
		fmt.Println("equal to two")
	default:
		fmt.Println("less than 1")
	}
    
	checktype(10)
	checktype(3.14)
    checktype("golang")	
}

func checktype(x interface{}){
	switch x.(type){
	case int:
		fmt.Println("its an integer type")
	case float64:
		fmt.Println("its a float type")
	case string:
		fmt.Println("its a string type")
	default:
		fmt.Println("its of unknown type")
	}
}