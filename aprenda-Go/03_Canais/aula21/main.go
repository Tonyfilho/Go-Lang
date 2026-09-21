package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula 21 - Canais e Goroutines")
	/**Canais é forma de comunicação entre goroutines sem ter problemas de concorrência*/
	fmt.Println()
	fmt.Println("Canais são o Jeito Certo® de fazer sincronização e código concorrente.")
	/***/
	fmt.Println("Eles nos permitem trasmitir valores entre goroutines.")
	fmt.Println()
	fmt.Println("Servem pra coordenar, sincronizar, orquestrar, e buffering.")
	/**Temos q criar uma canal, usando a palavra-chave make(chan type)*/
	/**Ao Criar um canal, ele é inicializado com um buffer de tamanho zero,
	o que significa que ele é um canal sem buffer, e isto tem que ser Usado, ou seja outro canal tem que fazer uso
	caso contrario teremos error.
	* Imagine um corredor de corrida, onde um corredor precisa passar a bastão para outro corredor.
	* O primeiro corredor não pode continuar correndo até que o segundo corredor esteja pronto para receber o bastão.
	* Da mesma forma, um canal sem buffer bloqueia a goroutine que está enviando o valor até que outra
	* goroutine esteja pronta para receber esse valor.
	*/
	fmt.Println("Temos q criar uma canal, usando a palavra-chave make(chan type)")
	fmt.Println()
	fmt.Println("Alem de usar o make, a tribuição é por meio do operador: <- \n, que é chamado de operador de envio e recebimento.")

	meuCanal := make(chan string) // Criando um canal do tipo string
	go func() {
		meuCanal <- "Olá, Sou um Canal e Preciso de uma Outro Canal para funcionar!" // Enviando valor para o canal
	}()

	fmt.Println(<-meuCanal) // Recebendo valor do canal
	fmt.Println("Temos 2 Goroutines, uma é a fmt.Println e outra é a anônima, que envia o valor para o canal.")
	fmt.Println()
	fmt.Println("Buffered Channels")
	fmt.Println("Com Buffer só precisamos de 1 Goroutine para enviar valores para o canal \n, e a outra goroutine pode receber os valores quando estiver pronta.")
	fmt.Println()
	
	/**Para usar um Canal, preciso de  2 goroutines*/
	
	/**Podemos criar canais com buffer, usando a função make com um segundo
	parâmetro que especifica o tamanho do buffer.
	* O Buffer permite que a goroutine que envia valores para o canal continue
	* executando mesmo que a goroutine que recebe os valores ainda não esteja
	* pronta para recebê-los. Em outras palavras so preciso de 1 goroutine para enviar valores
	* para o canal, e a outra goroutine pode receber os valores quando estiver pronta.
	*/
	
	meuCanalComBuffer := make(chan string, 2) // Criando um canal com buffer de tamanho 2
	
	meuCanalComBuffer <- "Olá, Sou um Canal com Buffer e Preciso de uma Outro Canal para funcionar!"
	fmt.Println("Recebendo valor do canal com buffer: ",<-meuCanalComBuffer) // Recebendo valor do canal com buffer
	fmt.Println()
	meuCanalComBuffer <- "Sou o Segundo Valor do Canal com Buffer!"
	fmt.Println("Recebendo valor do canal com buffer: ",<-meuCanalComBuffer) // Recebendo valor do canal com buffer
	fmt.Println()
	/**Temos um Buffer de Tamanho 2, e podemos enviar até 2 valores para ele */
	fmt.Println("Temos um Buffer de Tamanho 2, e podemos enviar até 2 valores para ele")
	fmt.Println()
}
