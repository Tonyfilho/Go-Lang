package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula21B. Assignment/conversion")
	fmt.Println()
	
	
	/*Criaremos um canal bidirecional e depois colocaremos comportamento unidirecional nele*/
	/**Podemos pegar um canal bidirecional  e manipularmos ele em canais unidirecionais*/
	canal := make(chan int)
	
	
	/**Porque  uma go Func ? sem buffer, a função send vai esperar a função receive receber o valor do canal, 
	caso contrário teremos um deadlock*/
	go send(canal) 
	/**a tribuição do canal será feita para a função receive, que é do tipo unidirecional de recebimento**/
	receive(canal)

}


func send(c chan<- int) {
	c <- 42
}


func receive(c <-chan int) {
	fmt.Println()
	fmt.Println("Função receive:", <-c)
}