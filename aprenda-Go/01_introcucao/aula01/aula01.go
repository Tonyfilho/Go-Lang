package main

import "fmt"

var globalVariable string = "I am a local GLOBAL variable"
var globalIntVariable int = 42

// vendo o zero value de uma variável não inicializada
var z int
var w string
var u bool
var v float64	

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

	fmt.Printf("Zero value of z: %v %T\n", z, z)
	fmt.Printf("Zero value of w: %v %T\n", w, w)
	fmt.Printf("Zero value of u: %v %T\n", u, u)
	fmt.Printf("Zero value of v: %v %T\n", v, v)

}
