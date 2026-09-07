package main

import (
	"fmt"
)

func main() {
	fmt.Println("Funcs  1. Funções")
	/**Invocando a função*/
	basica()

	fmt.Print()
	fmt.Println("Funcs 2 Função que aceita um argumento")

	argumento("manhã")
	argumento("tarde")
	argumento("jfdhjf")

	fmt.Print()
	fmt.Println("Funcs 3  Função com retorno.")

	valor := soma(10, 10)
	fmt.Println(valor)

	fmt.Print()
	fmt.Println("Funcs 4  FFunção com múltiplos retornos e parâmetro variádico.")

	/**Observe q estamos atribuindo todos os valores retornados de forma linear*/
	total, quantos, oi := somaMultRetornos(10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos, oi)

	total2, quantos, oi := somaMultRetornosComSaudacao("tarde", 10, 10, 1, 2, 3, 5)

	fmt.Println(total2, quantos, oi)

	fmt.Print()
	fmt.Println("Funcs 5  Função com Slice.")

	si := []int{10, 10, 1, 2, 3, 5}
	total3 := somaWithSlice(si...)
	fmt.Println(total3)

	/**Podemos passar em funções com argumentos variádicos valor vazio,
	ou seja, sem argumentos, ou
	podemos passar um slice com o operador ... (si...)*/
	total4 := somaWithSlice()
	fmt.Println("Variadico não é obrigatório passar valores neste caso retornará 0, pois é INT:", total4)


	/**Defer é uma forma de criar assincronismo*/
	fmt.Print()
	fmt.Println("Funcs 6  Função Defer, será criado assincronismo.")

	defer fmt.Println("1") // a ordem de execução do defer é LIFO, ou seja, o último a ser chamado será o primeiro a ser executado
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}

/**01 Funções Básicas*/
func basica() {
	fmt.Println("Oi, bom dia!")
}

/**02 Função que aceita um argumento*/

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

/**03 Função que retorna um valor, observem a TIPAGEM*/
func soma(x, y int) int {
	return x + y
}

/**Obs os multimos Argumentos usando o operador ... (x ...int) */
/**04 Função com múltiplos retornos e parâmetro variádico, Observem as TIPAGENS  (int, int, string)*/
/**È esquisito mas útil*/
func somaMultRetornos(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "Bom dia!"
}

/**04B Função de retorno múltiplo e mais de um parâmetro, obs: posso ter varios int recebidos, mas somente 1 string*/

func somaMultRetornosComSaudacao(s string, x ...int) (int, int, string) {
	oi := ""
	if s == "manhã" {
		oi = "Oi, bom dia!"
	} else if s == "tarde" {
		oi = "Oi, boa tarde!"
	} else {
		oi = "Oi, boa noite!"
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), oi
}

/**05 Função com Slice*/

func somaWithSlice(x ...int) int {

	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
