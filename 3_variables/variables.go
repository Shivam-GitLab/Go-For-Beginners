package main

import "fmt"

func main() {
	var name string = "golang"

	sm := name
	fmt.Print(name)
	fmt.Print(sm)

	// infer
	var names = "golang"
	var isAdult bool = true
	fmt.Print(names)
	fmt.Print(isAdult)

	// var age int = 30

	// shorthand syntax
	// name := "golang"

	// var name string
	// name = "golang"

	// var price float32 = 50.5
	// var price = 50.5
	// price := 50.5

	// fmt.Println(price)
}
