package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game")
	fmt.Printf("Enter your name: \n")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yep, u can play!")
	} else {
		fmt.Println("Sorry, u cannot play!")
		return
	}

	score := 0
	num_questions := 2
	fmt.Printf("What is better, the Macbook Pro or Macbook Air? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "Macbook Air" || answer+" "+answer2 == "macbook air" {
		fmt.Println("U are right!")
		score++
	} else {
		fmt.Println("U are not right!")
	}

	fmt.Printf("How many cores does the Ryzen 5 7500f have?")
	var cores uint
	fmt.Scan(&cores)
	if cores == 6 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Printf("You scored %v out of %v.", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
