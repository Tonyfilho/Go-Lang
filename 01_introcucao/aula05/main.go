package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula 05 - Fluxos de Controle,Loops e Outros!")

	/** 01 For simples*/
	fmt.Println("\n 01 For e if juntos")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/** 02 For e if juntos*/
	fmt.Println("\n 02 For e if juntos")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é ímpar")
		}

	}

	/** 02A for encadeado, de agoras ate 5 horas*/
	fmt.Println("\n 02A for encadeado, de agoras ate 5 horas")
	for horas := 0; horas < 5; horas++ {
		fmt.Printf("Horas: %d\n", horas)
		/**este é o loop de minutos vai iterar 60 vezes para um valor do loop de horas*/
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Printf("horas: %d, minutos: %d\n", horas, minutos)
		}
		fmt.Println("\n Fim do loop de minutos")
	}

	/** 03 For no lugar do While, uso da palavra-chave break*/
	fmt.Println("\n 03 For uso  break")
	x := 0
	for x < 10 {
		fmt.Println("Em Go o For desta forma substitui While , pois não existe While: X < 10: ", x)
		x++
	}

	/**For no lugar do DoWhile um For infinito iqualmente os servidores HTTP fazem com DoWhile*/
	for {
		if x > 10 {
			fmt.Println("Loop infinito, usamos  break para sair do loop")
			break
		}
		fmt.Println("Loop infinito, para sair use Ctrl + C, ou use break para sair do loop")
		x++

	}

	/** 04 For uso do Break e Continue*/
	fmt.Println("\n 04 For uso  Continue")

	x1 := 0

	for x1 <= 20 {

		x1++

		if x1%2 != 0 {
			fmt.Println("O Break é usado para sair do loop, o Continue é usado para pular a iteração atual e ir para a próxima iteração do loop")
			continue
		}

		fmt.Println("Imprimindo os números pares:", x1)

	}

	/**05 Desafio supresa*/
	// fmt.Println("\n 05 Desafio supresa")

	// for x := 33; x <= 122; x++ {
	// 	fmt.Printf("%d - %v\n", x, string(x))
	// }

	fmt.Println("\n 06 Declaração IF")

	/** 06 If Basico, observe a delaração da variavel onde pode ficar, entre o if e o bloco de código */
	fmt.Println("\n 06 Declaração IF Básico")
	if x2 := 10; !(x2 > 100) {
		fmt.Println("Hello, playground")
	}

	/** 06 If & Else*/
	fmt.Println("\n 06 Declaração IF & Else")

	if x3 := 500; x > 100 {
		fmt.Println("Tony é maior que cem")
	} else if x3 < 10 {
		fmt.Println("Tony é menor que déis")
	} else {
		fmt.Println("Tony não é menor que déis nem maior que cem")
	}

	/**07 Declaração Switch*/
	fmt.Println("\n 07 Declaração Switch")

	switch x4 := 10; x4 {
	case 0:
		fmt.Println("x4 é zero")
	case 1:
		fmt.Println("x4 é um")
	default:
		fmt.Println("x4 não é zero nem um")
	}

	/**declaração switch com fallthrough*/
	/**07 Declaração Switch com fallthrough ou precipitação*/
	fmt.Println("\n 07 Declaração Switch com fallthrough, faz com que o próximo case seja executado mesmo que não seja verdadeiro")

	switch x5 := 10; x5 {
	case 0:
		fmt.Println("x5 é zero fallthrough")
		fallthrough
	case 1:
		fmt.Println("x5 é um fallthrough")
		fallthrough
	default:
		fmt.Println("x5 não é zero nem um, ele pula para o próximo case mesmo que não seja verdadeiro, isso é o fallthrough")
	}

	/**Filtrando não valores e sim TIPOS*/
	fmt.Println("\n 08 Filtrando não valores e sim TIPOS")

	switch v := interface{}(10); v.(type) {
	case int:
		fmt.Println("v é um int")
	case string:
		fmt.Println("v é um string")
	case bool:
		fmt.Println("v é um bool")
	case float64:
		fmt.Println("v é um float64")
	default:
		fmt.Println("v é de outro tipo")
	}

}
