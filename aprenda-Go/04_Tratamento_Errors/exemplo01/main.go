package main

import (
	"fmt"
)





func main() {
	fmt.Println("#### Aula23 - Tratamento de Erros – 1. Entendendo erros")
	fmt.Println("Obs: A Funcão Scanln() retorna dois valores, o primeiro \n é a quantidade de valores lidos e o segundo é o erro, caso ocorra algum erro.")
	fmt.Println("Neste Exemplo usaremos o PANIC.")

	var anser1, anser2, anser3 string
	fmt.Print("Name: ")
	_ , err := fmt.Scanln(&anser1)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	fmt.Println()

	fmt.Print("Favorite Food: ")
	_ , err = fmt.Scanln(&anser2)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	fmt.Println()

	fmt.Print("Favorite Drink: ")
	_ , err = fmt.Scanln(&anser3)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	fmt.Println()

	fmt.Println("As Respostas: ", anser1, anser2, anser3)

}