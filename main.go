package main

import (
	"demo-app/greetings"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
