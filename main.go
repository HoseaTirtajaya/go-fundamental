package main

import "fmt"

func main() {
	// var name string // TO DECLARE A VARIABLE WITH EMPTY VALUE
	var name = "Hosea Tirtajaya" // TO DECLARE A RANDOM VALUE WITHIN VARIABLE
	var age = 21

	fmt.Println("Halo saya ", name, ". Saya berumur ", age)

	// fmt.Println("NUMBERS: INI SATU"[3])

	//DECLARING MULTIPLE VARIABLES
	var (
		firstName = "Hosea"
		lastName  = "Tirtajaya"
	)

	fmt.Println(firstName, lastName)

	//CONSTANT
	const alwaysTrue = true
	fmt.Println(alwaysTrue)

	//const can use const (
	// firstName = hosea
	// ) like var up there to declare multiple constants
}
