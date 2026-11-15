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
	fmt.Print("What's your age ", name, "?")
	fmt.Scan(&age)
	fmt.Print("So dear ", name, " you are ", age, " years old")

}
