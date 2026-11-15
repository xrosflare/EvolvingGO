package main

import "fmt"

func main() {
	var name string
	var age int

	greeting := "Hello to the GO world"

	fmt.Println(greeting)
	fmt.Println("What's your name ?")
	fmt.Scan(&name)
	fmt.Println("Welcome", name, "!")
	fmt.Println("What's your age ", name, "?")
	fmt.Scan(&age)
	fmt.Println("So dear ", name, " you are ", age, " years old")
	fmt.Println("So there's ", ageCalc(age), " year left to you to reach the 100 years old birthday")

}

func ageCalc(quantity int) int {
	return 100 - quantity
}
