package main

import (
	"fmt"
	//"sync"
)

// - Chans par, ímpar, quit
// - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// - Func receive é um select entre os três canais, encerra no quit
func main() {
	fmt.Println("#### Aula21E   Canais – 4. Select # Exemplo 3:")
	fmt.Println()
	fmt.Println("Chans par, ímpar, quit")
	fmt.Println()
	fmt.Println("Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit")
	fmt.Println()
	fmt.Println("Func receive é um select entre os três canais, encerra no quit. Problema!")
	fmt.Println()

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)
	/**invocando a go routines 1x para n teremos q criar working group*/
	go mandaNumerosCanail(par, impar, quit)
	receiverCanal(par, impar, quit)

}
func mandaNumerosCanail(par, impar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
//	fmt.Println("close(quit) .....", quit)
	

}

func receiverCanal(par, impar chan int, quit chan bool) {
	/**Criaremos o For eterno e o Select*/
	for {
		select {
		case v := <-par:
			fmt.Println("O numero: , ", v, " é Par")
		case v := <-impar:
			fmt.Println("O numero: , ", v, " é Impar")
		case <-quit:
			fmt.Println("Quit...... Acabouuuuuuu, mas repita e veja o problema, o ZERO será impar e par, veremos a solução na proxima aula21F")
			return
		}
	}
}
