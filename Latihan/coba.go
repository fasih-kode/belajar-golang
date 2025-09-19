package main

import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
	return "Goodbye, Mars!"
}

func main() {
	HelloWorld()
	fmt.Println(HelloWorld())
}
