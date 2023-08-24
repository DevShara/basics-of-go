package main

import "fmt"

func main() {
	fmt.Println("Start of the program.")

	name := "Alice"

	if name != "Bob" {
		defer panic("Unexpected name: " + name)
	}

	fmt.Println("End of the program.") // This line will not be executed
}
