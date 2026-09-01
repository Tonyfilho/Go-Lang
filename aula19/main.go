package main

import (
	"fmt"
)

func Func01() {
	fmt.Println("Executando Func01")
	for i := 0; i < 5; i++ {
		fmt.Println("Func01:", i)
	}

}

func Func02() {
	fmt.Println("Executando Func02")
	for i := 0; i < 5; i++ {
		fmt.Println("Func02:", i)
	}
}

func main() {
	fmt.Println("Aula 19 Concorrência VS Paralelismo")
	fmt.Println("Concorrência é a capacidade de executar múltiplas tarefas em paralelo,\n mas não necessariamente em simultaneidade.")
	fmt.Println("Paralelismo é a capacidade de executar múltiplas tarefas em simultaneidade.")
	/**Concorrência é a capacidade de executar múltiplas tarefas em paralelo, mas não necessariamente em simultaneidade. */
	/**Paralelismo é a capacidade de executar múltiplas tarefas em simultaneidade. */
	fmt.Println()
	fmt.Println("Exemplo de Concorrência:")
	fmt.Println("Imagine 2 setas, uma a direita e outra a esquerda,\n elas podem se mover em paralelo, mas não necessariamente em simultaneidade.\n Ou seja, elas podem se mover em paralelo, mas não necessariamente ao mesmo tempo.")
	fmt.Println()
	fmt.Println("Exemplo de Paralelismo:")
	fmt.Println("Imagine 2 linhas de produção, uma produzindo carros e outra produzindo motos,\n elas podem produzir em simultaneidade, ou seja, ao mesmo tempo.\n Assim que termina 1 carro, termina 1 moto, e assim por diante.")
	fmt.Println()
	fmt.Println("Criaremos Funções Concorrentes o GO irá executar em paralelo, caso precise.\n Isto depende do número de núcleos do processador, se tiver 1 núcleo, ele irá executar em concorrência\n, caso tenha 2 núcleos, ele irá executar em paralelismo.")
	fmt.Println()
	fmt.Println("O Runtime do GO irá gerenciar a execução das funções concorrentes, caso tenha 1 núcleo, ele irá executar em concorrência\n, caso tenha 2 núcleos, ele irá executar em paralelismo automaticamente.")
	fmt.Println()
	fmt.Println()
	fmt.Println("Aula19 Concorrencia com Goroutines & WaitGroups")
	fmt.Println("O código abaixo é linear. Como fazer as duas funções rodarem concorrentemente?")
	/**Teremos 2 funções que são executadas linearmente Func01() e Func02()*/
	Func01()
	Func02()
	fmt.Println("As funções Func01() e Func02() foram executadas linearmente, ou seja, uma após a outra.")
	fmt.Println()
	/**É aqui que usamos as Goroutines */
	/**Goroutines são funções que podem ser executadas concorrentemente , parecidas com as threads do Java */
	fmt.Println("No próximo exemplo, vamos executar as funções Func01() e Func02() concorrentemente\n, ou seja, ao mesmo tempo.")
	fmt.Println("Fazemos isso utilizando Goroutines, que são funções que podem ser executadas concorrentemente.")
	/**hread (em português: fio de execução[1] ou encadeamento de execução) é uma forma como um processo/tarefa
	de um programa de computador é divido em duas ou mais tarefas que podem ser executadas concorrentemente ("simultâneo"). */
	fmt.Println()
	/**Na Pratica, temos q usar o goFunc*/
	/*Na pratica basta eu por a palavra-chave go antes da chamada da função */
	fmt.Println("Na pratica basta eu por a palavra-chave go antes da chamada \n da função Ex: go Func01() e go Func02()	")
	fmt.Println()
	/**Mas depois que for executado, precisamos esperar a conclusão das goroutines ,
	não conseguimos controlar a ordem de execução e nem mudar o estado delas, pense
	 em um foguete q depois de lançado para espaço já não pode ser controlado */
	/**Não vai aparecer o print da goroutine , pois a função main() terminou antes da goroutine go Func01()*/
	go Func01()
	Func02()
	/* Em outras palavras, a função main() terminou antes da goroutine go Func01() e foi descartada */
	fmt.Println("Em outras palavras, a função main() terminou antes da goroutine go Func01() e foi descartada.")
	fmt.Println()
	fmt.Println("Para controlar a execução das goroutines, podemos usar o WaitGroup do pacote sync")
	fmt.Println()
	fmt.Println("O WaitGroup é uma estrutura que permite esperar a conclusão de um conjunto de goroutines.")
	fmt.Println()
	fmt.Println("Ele possui 3 métodos: Add(), Done() e Wait()")
	fmt.Println()
	fmt.Println("Add() adiciona uma goroutine ao WaitGroup, Done() indica que a goroutine terminou e Wait() espera a conclusão de todas as goroutines.")
	fmt.Println()
	


}
