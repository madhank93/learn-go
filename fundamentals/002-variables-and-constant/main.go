package main

import "fmt"

func main() {

	// variables declared without initialization
	var age int
	var first_name string
	fmt.Println(age)
	fmt.Println(first_name)

	// variables declared with value initialized
	var x int = 5
	var rad = 10
	fmt.Println(x)
	fmt.Println(rad)

	// declaration of multiple variables
	// Ex 1
	var (
		a         int = 10
		b             = 15
		greetings     = "hello"
	)
	fmt.Println(a, b, greetings)

	// Ex 1
	var d, e int = 1, 2
	fmt.Println(d, e)

	// shorthand declaration of variables
	y := 7
	fmt.Println(y)

	// declaration of constants
	const pi float64 = 3.14272345

	// declaration of a string
	var name string = "Rob Pike"
	fmt.Println(name)

}
