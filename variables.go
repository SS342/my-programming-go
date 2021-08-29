package main

import "fmt"

var q string = "Hello World"

func main() {

	var x string = "Hello Word ! "
	fmt.Println(x)

	var y string
	y = "Hi!"
	fmt.Println(y)

	var z string

	z = "one"
	fmt.Println(z)

	z = "two"
	fmt.Println(z)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	fmt.Println(q)

	const BestLanguage = "Python"

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)
}
