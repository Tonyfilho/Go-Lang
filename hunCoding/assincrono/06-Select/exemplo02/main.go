package main

import (
	"fmt"
	"time"
)

/**Criaremos um metodo que pega o valor de Channel que recebe valor somente**/
/**Supondo q seja uma chamada em API externa, q demora justamente o time.Duration */

func getValueFromChannelAPI(ch chan<- string, duration time.Duration) {

	time.Sleep(duration)

	ch <- "Terminamos por aqui!"
}

/**Recebendo sinais de varios Channels e selecionando*/
func main() {
	fmt.Println("Aula06 Select em  GoRoutines")
	fmt.Println("Entendendo o funcionamento do SELECT e FOR eterno")
	fmt.Println()

	ch1 := make(chan string)
	go getValueFromChannelAPI(ch1, 2*time.Second)
	ch2 := make(chan string)
	go getValueFromChannelAPI(ch2, 2*time.Second)
	ch3 := make(chan string)
	go getValueFromChannelAPI(ch3, 2*time.Second)

	/**Usaremos o SELECT a 1º GoRoutine q retornar ele termina*/
	/**Quando adicionamos o FOR eterno temos todos os retornos, mas somente o For teremos fatal.Error*/
	for {
		select {
		case localReturno := <-ch1:
			fmt.Println()
			fmt.Print("Canal 01 retornou: ", localReturno)
		case localReturno := <-ch2:
			fmt.Println()
			fmt.Print("Canal 02 retornou: ", localReturno)
		case localReturno := <-ch3:
			fmt.Println()
			fmt.Print("Canal 03 retornou: ", localReturno)
		}
	}
	
	fmt.Println("Main Terminou........")
}
