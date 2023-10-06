package main

import "fmt"

func main() {
	// Call the function
	fmt.Println(sayHello())

	// Call the function
	fmt.Println(sayGoodbye())
}

// Define the function
func sayHello() string {
	return "Hello World!"
}

func sayGoodbye() string {
	return "Goodbye World!"
}
