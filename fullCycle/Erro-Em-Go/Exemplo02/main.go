package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Aula do Video04 Errors em Go Exemplo02 Criando seu proprio Error")
	fmt.Println()
	/*Aqui temos que criar as 2 variavels**/
	res, err := soma(1, 10)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Nosso Valor de Res: ", res)
	fmt.Println()
	fmt.Println("Causando um erro na Func Soma")
	fmt.Println()

	res, err = soma(20, 10)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Nosso Valor de Res será ZERO: ", res)
	fmt.Println()

}

/*Criaremos um função simples onde teremos o retorno de um dado e do erro caso haja**/
func soma(x, y int) (int, error) {
	total := x + y
	if total > 20 {
		// não interessa o valor quero o construtor do error
		return 0, errors.New("Error na Função Soma! Total maior que  20: ")
	}
	/*Agora tenho que Retornar SEMPRE 2 valore, neste caso o erro será NIL**/
	return total, nil
}
