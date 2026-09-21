package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula21B. Canais direcionais & Utilizando canais")
	fmt.Println()
	fmt.Println("## bidirecionais (send & receive)")
	fmt.Println()
	fmt.Println("## unidirecionais (send only) ou send chan←")
	fmt.Println()
	fmt.Println("## unidirecionais (receive only) ou receive <-chan")
	
	/*Criaremos um canal bidirecional e depois colocaremos comportamento unidirecional nele*/
	/**Podemos pegar um canal bidirecional  e manipularmos ele em canais unidirecionais*/
	canal := make(chan int)
	
	/**Criaremos 2 funções unidirecionais , uma send e uma receive*/
	/**a tribuição do canal será feita para a função send, que é do tipo unidirecional de envio**/
	/**OBS: Lembrando isto tem q rodar de maneira concorrente*/
	fmt.Println()
	fmt.Println("Porque por go send(canal), por ser um goroutine sem buffer, a função send vai esperar a função\n receive receber o valor do canal, caso contrário teremos um deadlock")
	fmt.Println()
	fmt.Println("Mas como o canal não tem buffer, o Go precisa encontrar alguém fazendo: receive(canal) para que a função send(canal) \n possa enviar o valor 42 para o canal, caso contrário teremos um deadlock")
	fmt.Println()
	/**Porque  uma go Func ? sem buffer, a função send vai esperar a função receive receber o valor do canal, 
	caso contrário teremos um deadlock*/
	go send(canal) 
	/**a tribuição do canal será feita para a função receive, que é do tipo unidirecional de recebimento**/
	receive(canal)

}

/**Observe a seta, depois do CHAN mostra o tipo de canal SEND */
/**Na função send, recebemos valores em um canal unidirecional de envio
e colocamos dentro desse canal c chan<- int */
func send(c chan<- int) {
	c <- 42
}

/**Observe a seta, depois do CHAN mostra o tipo de canal RECEIVE */
/**Na função receive, enviamos valores através de um canal unidirecional de recepção
* ou podemos disse que retiramos valores de um canal unidirecional de recepção c <-chan int
*/
func receive(c <-chan int) {
	fmt.Println()
	fmt.Println("Função receive:", <-c)
}
