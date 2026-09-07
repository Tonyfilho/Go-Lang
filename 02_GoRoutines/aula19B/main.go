package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Func01() {
	fmt.Println("Executando Func01")
	for i := 0; i < 5; i++ {
		fmt.Println("Func01:", i)
		time.Sleep(20) // Forçando um tempo de execução
	}

}

func Func02() {
	fmt.Println("Executando Func02")
	for i := 0; i < 5; i++ {
		fmt.Println("Func02:", i)
		time.Sleep(20) // Forçando um tempo de execução
	}
}

/**Criando um WaitGroup*/
var wg sync.WaitGroup

func main() {
	fmt.Println("Aula 19B - Concorrência – 2. Goroutines & WaitGroups")
	fmt.Println("Em outras palavras, a função main() É uma Go routine que terminou antes da goroutine go Func01() e foi descartada")
	/**Como ja vimos o main() terminou antes da goroutine go Func01() e foi descartada e precisamos usar o WaitGroup  para controlar a execução */
	/**Antes da goroutine, imprimindo os cores dos processadores*/
	
	Func01()
	Func02()
	fmt.Println()
	
	fmt.Println("Número de CPUs Antes:", runtime.NumCPU())
	fmt.Println("Número de Go Routines Antes:", runtime.NumGoroutine()) //Número de Go Routines Antes: 1
	/**Agora colocamos o Add(numero de goroutines)*/
	wg.Add(2) // Adicionando 2 goroutines q vai ter de esperar a conclusão

	go func() {
		Func01()
		wg.Done() // Indicando que a goroutine terminou
	}()

	go func() {
		Func02()
		wg.Done() // Indicando que a goroutine terminou
	}()
	fmt.Println("Número de CPUs Depois:", runtime.NumCPU())
	fmt.Println("Número de Go Routines Depois:", runtime.NumGoroutine())// Número de Go Routines Depois: 3
    /**A Saida são 3, pois o main() é uma goroutine e as outras duas são as funções executadas */
	wg.Wait() // Esperando a conclusão de todas as goroutines

	fmt.Println("As funções Func01() e Func02() foram executadas concorrentemente, ou seja, ao mesmo tempo.")
	fmt.Println("Adicionamos um time.Sleep(20) para dar  um tempo ao processando de fazer as goroutines de execução.")

	fmt.Println()
	fmt.Println("Func01: 0 Func01: 1 Func02: 1 Func02: 2 Func01: 2 Func01: 3 Func02: 3 Func02: 4 ...e continua até Func01: 9 Func02: 9")

	/** sainda no terminal, veja a itercão das funções Func01() e Func02() executando concorrentemente, ou seja, ao mesmo tempo.
	Func01: 0 	Func01: 1  	Func02: 1  	Func02: 2  	Func01: 2  	Func01: 3  	Func02: 3  	Func02: 4 	Func01: 4 	Func01: 5 	Func02: 5
	* Func02: 6 Func01: 6 	Func01: 7  	Func02: 7	Func02: 8 Func01: 8     Func01: 9 	Func02: 9*/
}
