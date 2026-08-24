package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("#### Aula13 Recursividade")

	fmt.Println(fatorial(4))
}

func fatorial(x int) int {
	// 1º quando for iqual a 1 eu termina com retorno de 1 
	if x == 1 {
		return x
	}
	
	// 2º quando for diferente retorno o Fatorial Inferior fatorial(x-1) VEZES o numero atual
	/* 4! = 4 * 3 * 2 * 1 (4 X fatorial de 3, 3 X fatorial de 2, 2 X fatorial de 1 e 1 X fatorial de 1.) **/
	return x * fatorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x // total = total * x
		x--
	}
	return total
}
