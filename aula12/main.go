package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("#### Aula12 Funções Callback")
	/*temos um função sendo invocada com 2 paramentros: soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...
	1º é passando uma outra função neste caso SOMA Obs: não invocamos
	2º é passado um array e temos o operador ... SliceOfInts ou Rest or Destruction itera os dados */
	t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)


	/*#### Aula12B Closure*/
	fmt.Println()
	fmt.Println("#### Aula12B Closure")

	/* Criamos um Expressão e a avarivel A vai se tonar na função 
	interna da callback return X, e X será somando */
	a := i()
	
	fmt.Println("a: ", a()) // retorno 1
	fmt.Println("a: ", a()) // retorno 2
	fmt.Println("a: ", a()) // retorno 3
	
	/* Criamos um Expressão e a avarivel B vai se tonar 
	na função interna da callback return X, e X será somando.
	Agora aqui é criado uma nova referencia de memoria e temos valores diferentes por causa do ESCOPO*/
	b := i()

	fmt.Println("b: ", b()) // retorno 1
	fmt.Println("b: ", b()) // retorno 2
	fmt.Println("b: ",b())  // retorno 3
	fmt.Println("b: ", b())  // retorno 4

}

/*
*Aula 12 Função soma que será usada dentro da função somentePares()
1º Ela recebe finito numeros de parametros do tipo int usando o operador SliceOfInt que prefiro chamar de Rest ou Destruction
*/
func soma(x ...int) int {
	// crio um var para receber a soma e retornará
	n := 0
	// criamos um for e usamos range de X
	for _, v := range x {
		n += v
	}
	return n
}

/** temos a função de callBack onde por paramentro receberemos OUTRA função */
/** 1º é passando uma outra função neste caso SOMA Obs: não invocamos
2º é passado um array e temos o operador ... SliceOfInts ou Rest or Destruction itera os dados
*/
func somentePares(f func(x ...int) int, y ...int) int {
	// criamos um array de INT
	var slice []int
	// criamos um for e o RAGE é a variavel q recebe os ints Y
	for _, v := range y {
		// usaremos o operador de RESTO % para separar o que é par ou impar
		if v%2 == 0 {
			// usamos o metodo Append  que recebe o um array e um valor e cria uma nova referencia na memoria com tamanho MAIOR
			slice = append(slice, v)
		}
	}
	// f func(x ...int) criamos uma nona variavel que RECEBE a FUNÇÂO, em nosso caso de SOMA, poderia ser outra função
	// y ...int  usaremos o operador ... sliceOfInt para iterar o array e retonamos o total
	total := f(slice...)
	return total
}

/**#### Aula12B Closure */

func i() func() int {
	// OU seja esta variavel X terá valores diferentes para cada Referencia de Memoria
	// E ainda somará o valores nas mesma Referencia de memoria.
	x := 0 // será usado uma variavel do escopo Externo, gerando copias diferentes para cada Expressão Criada 
	// aqui em baixo termos o CLOSURE onde retornamos o valor deste escopo interno e não extermo
	return func() int {
		///Ou seja para cada Invocação, será somado os valor de X
		x++
		return x
	}
}
