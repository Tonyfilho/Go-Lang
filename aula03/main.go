package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {
	// 02 O tipo booleano só pode assumir os valores true ou false
	fmt.Println(x) // zero value
	x = true
	fmt.Println(x) // valor atribuido
	x = 10 == 100
	y := 10 == 100
	z := 10 > 100
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	a := 10
	b := 20
	c := 10.0
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(c == 10.0)

	// 03 Inteiros e floats podem ser comparados com operadores de comparação

	a1 := "e"
	b1 := "é"
	c1 := "香"
	fmt.Printf("%v, %v, %v\n", a1, b1, c)

	d := []byte(a1)
	e := []byte(b1)
	f := []byte(c1)

	fmt.Printf("%v, %v, %v", d, e, f)

	///descobrir o sistema operacional e arquitetura do processador
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	/// 04 Overflow
	///uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535
	var i uint16
	i = 65535      ///valor máximo
	fmt.Println(i) ///valor máximo
	i++
	fmt.Println(i) ///valor mínimo,
	/// fui colocado + 1, ao invez de dar overflow, crai uma nova contagem o valor foi para o mínimo "0"
	i++
	fmt.Println(i) ///valor mínimo ao invez de dar overflow, crai uma nova contagem o valor foi para  "1"

}
