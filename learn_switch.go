package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	switch word {
	case "hello":
		fmt.Println("hi Yourself")
	case "goodbye":
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said")
	}
}
