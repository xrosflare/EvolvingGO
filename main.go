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
	fmt.Println(messageCalc(ageCalc(age)))

}

func ageCalc(quantity int) int {

	if quantity >= 100 {
		return 0
	}
	return 100 - quantity
}

func messageCalc(age int) string {
	if age >= 100 {
		return "Wow it's amazing , you're still young"
	}
	if age < 18 && age > 1 {
		return "keep growing"
	}
	return "enjoy the life"

}
