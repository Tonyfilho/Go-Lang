package main

import "fmt"

/**Neste Exemplo teremos um função ping q envia um canal e um dado(string)*/
	func ping(pingsChannelEnviaDados chan<- string, msg string)  {
		pingsChannelEnviaDados <- msg
	}

	/**Neste Exemplo teremos uma função que temos um canal que recebe um Ping e Passa para outro canal q Envia*/

	func pong(pingsChannelRecebeDados <-chan string, pongsChannelEnviaDados chan <- string)  {
		msg := <-pingsChannelRecebeDados //variavel recebendo de channel
		pongsChannelEnviaDados<- msg // variavel atribuindo ou Inserindo em Channel
		fmt.Println("Recebi o Ping e mando o Pong: ", msg)
		
	}


func main() {
/*Criamos um CANAL com o MAKE() e tipamos e colocamos o Bufferd 1**/
	pings := make(chan string ,1)
	pongs := make(chan string, 1)

	ping(pings, "Ping")
	pong(pings, pongs)
	/**Temos a entrada e saida temos que consumir este canal*/

	fmt.Println(<-pongs)

}