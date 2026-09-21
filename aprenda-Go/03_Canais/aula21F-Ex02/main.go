package main

import (
	"fmt"
	//"sync"
)

// - Chans par, ímpar, quit
// - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// - Func receive é um select entre os três canais, encerra no quit
func main() {
	fmt.Println("#### Aula21F   Canais – 5 ComaOK # Exemplo 1:")
	fmt.Println()
	fmt.Println("Como resolver os problema da aula21E Ex03")
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
	/**Temos q por o Canal quit antes do close()*/
	quit <- true
	close(par)
	close(impar)
	/*Aqui depois da Var quit, iremos por comaOk**/
	//v, ok := quit

}

func receiverCanal(par, impar chan int, quit chan bool) {
	/**Criaremos o For eterno e o Select*/
	for {
		select {
		case v := <-par:
			fmt.Println("O numero: , ", v, " é Par")
		case v := <-impar:
			fmt.Println("O numero: , ", v, " é Impar")
			/**Aqui no case do QUIT passaremos o ComaOK*/
		case v, ok := <-quit:
			if !ok {
				fmt.Println("ComaOk FALSE Erro ", v, ok)
			}
			fmt.Println("ComaOk TRUE FIM, este ZERO aqui é o valor VAZIO do Canal ", v, ok)
			return
			/*são por causa da leitura de um canal fechado. Na documentação do Go é dito que quando tentamos ler um canal
						*fechado ele responde com o valor zero do tipo do canal e um false no comma ok. Ali  no exemplo,
						*quando os canais 'par' e 'ímpar' são fechados, antes do true ser enviado para o 'quit',
						*a goroutine da função receive pode acabar realizando a leitura dos canais já fechados e recebendo um valor 0,
						*daí como não há a verificação do comma ok neles, esse 0 é tratado como um valor válido.
			*Acredito que uma possível solução seria enviar o true para o canal 'quit' e depois disso, fechar os
			canais 'par' e 'ímpar', assim não precisaria verificar o comma ok dos canais 'par' e 'ímpar', pois a
			*informação do canal canal 'quit' chegaria primeiro na função receive, e ela já pararia e não tentaria ler os demais canais.**/
		}
	}
}
