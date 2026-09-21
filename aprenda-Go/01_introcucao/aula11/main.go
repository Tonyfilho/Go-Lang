package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("#### Aula11 Funções Anônimas")

	x := 387

	/**Nas funções anonimas não temos nome e para fazer a invocação usamos (x), onde X é a variavel*/
	/**Aqui declara e executa ao mesmo tempo*/
	/**Muito usando com Go Rotinas, paralelismo*/
	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)

	fmt.Println()
	fmt.Println("#### Aula11A Func como expressão")

	/**Funções com Expressão**/

	xExpressao := 12

	yFuncExpressao := func(x int) int {
		//fmt.Println(xExpressao, "vezes 873648 é:")
		return x * 873648
	}

	fmt.Println(xExpressao, "vezes 873648 é:", yFuncExpressao(x))

	/**#### Aula11B Func Retorno Retornando uma função*/
	fmt.Println()
	fmt.Println("Aula11B Func Retorno Retornando uma função")
	
	/**Dividindo a função recursiva em 2 partes*/
	// 1º Criação da função Expressão, este retorno se tornou o valor da variavel xFuncRetorno
	xFuncRetorno := retornaumafuncao()
	
	// 2º Este retorno se  tornou a valor retornado da variavel xFuncRetorno
	yRecebeFunc := xFuncRetorno(3)
	
	// imprimindo o valor retornado
	fmt.Println("Retorno da Função", yRecebeFunc)
	
	// invocando a função retornaumafuncao()(4) diretamente sem ter que criar um função de expressão
	fmt.Println("Invocando a Função recursiva ", retornaumafuncao()(4))

	/**#### Aula11B Func Retorno Retornando uma função*/
	fmt.Println()
	fmt.Println("Aula11C Func CallBacks")
	
}

	/**#### Aula11B Func Retorno Retornando uma função*/

/**OBS Crio um funçao com a palavra Func e volto a usar a palavra Func()*/
/**Criamos uma função que retona um função, ao invez de por o retorno de um
primitivo colocamos a palavra FUNC()*/
func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
