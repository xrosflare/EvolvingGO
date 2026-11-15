package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var operator string
	fmt.Println("Welcome to calculator, please type the first number")
	fmt.Scan(&num1)
	fmt.Println("please type the second number")
	fmt.Scan(&num2)
	fmt.Println("What operation do you want to do with theses numbers ? simbol or name are allowed")
	fmt.Scan(&operator)
	fmt.Println("The result of the operation", operator, "with", num1, "by", num2, "is", calculator(num1, num2, operator))
}
func greedings() {
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

func calculator(num1 float64, num2 float64, operator any) float64 {

	if operator == "+" || operator == "soma" {
		return num1 + num2
	}
	if operator == "-" || operator == "subtracao" {
		return num1 - num2
	}
	if operator == "*" || operator == "multiplicacao" {
		return num1 * num2
	}
	if operator == "/" || operator == "divisao" {
		return num1 / num2
	}
	return 0
}
