package main

import "fmt"

func main() {

	fmt.Println("Int:")
	fmt.Println("1 + 1 = ", 1+1)

	fmt.Println(" ")

	fmt.Println("Float:")
	fmt.Println("1 + 1 = ", 1.0+1.0)

	fmt.Println(" ")

	fmt.Println("Hello World!")

	fmt.Println("Len:")
	fmt.Println(len("Hello Word"))

	fmt.Println(" ")

	fmt.Println("Hello word"[0])
	fmt.Println("Hello" + " Word" + "!")

	fmt.Println(" ")

	fmt.Println("Logic:")

	fmt.Println(" ")

	fmt.Println(true && true)
	fmt.Println(true && false)

	fmt.Println(true || true)
	fmt.Println(true || false)

	fmt.Println(!true)

}
