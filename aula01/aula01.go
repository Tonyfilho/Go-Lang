package main

import "fmt"

var globalVariable string = "I am a local GLOBAL variable"
var globalIntVariable int = 42
func main() {
    name := "Go Developers"
    fmt.Println("Tony", name)

	// operador  curto de declaração de variável
	x := 10
	y := "Hello"
	fmt.Println(x, y)

	// vendo a typagem usando o operador %T
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("globalVariable is of type %T\n", globalVariable)
	fmt.Printf("globalIntVariable is of type %T\n", globalIntVariable)


}