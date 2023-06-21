package main

import "fmt"

func main() {

	// print line
	// fmt.Println("Hello World!")
	// // variable w/ no explicit type
	// name := "Ian"
	// fmt.Println(name)
	// fmt.Printf("Hello %v, how are you doing?", name)

	var name string
	var age int

	fmt.Println("Welcome to my Quiz game!")

	fmt.Printf("Name: ")
	fmt.Scan(&name)
	fmt.Printf("Age: ")
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Printf("Hello %v, how are you? You're %v years old! Welcome to the Game!!", name, age)
	} else {
		fmt.Printf("Sorry, %v years old are not allowed to play", float64(age))
	}

	var num int
	num = 1

	percent := float64(num) / float64(age) * 100
	fmt.Printf("Percentage is %v%%", percent)

}
