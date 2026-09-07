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


/**Criando um loop no lugar da recursividade*/
func loops(x int) int {
	// 1º Cria a variavel que retornará, qualquer numero multiplicado por 1 será ele mesmo
	total := 1
	// 2º faço um FOR de uma variavel x que deve ser MAIOR que 1
	for x > 1 {
		// 3º faço o multiplicação dentro do for da variavel X e faço a atribuição para a varial Total
		total *= x // total = total * x
		// 4º tenho que fazer o decremento da variavel X
		x--
	}
	return total
}
