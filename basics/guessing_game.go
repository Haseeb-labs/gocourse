package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//generate a random number from 1 to 100
	target := random.Intn(100) + 1

	//welcome message
	fmt.Println("welcome to the guessing game")
	fmt.Println("I have chosen a number from 1 to 100")
	fmt.Println("try to guess it")

	var guess int
	for {
		fmt.Println("enter your guess : ")
		//if you dont use "&" it will create a copy of guess and stores the value in it
		fmt.Scan(&guess)

		if guess == target {
			fmt.Println("aha you found it!!!")
		} else if guess < target {
			fmt.Println("TRY higher")
		} else {
			fmt.Println("TRY lower")
		}
	}

}
