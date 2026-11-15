package main

import "fmt"

func main() {
	var name string

	greeting := "Hello GO world"

	fmt.Println(greeting)
	fmt.Print("What's your name ?")
	fmt.Scan(&name)
	fmt.Println("Welcome", name, "!")

}
