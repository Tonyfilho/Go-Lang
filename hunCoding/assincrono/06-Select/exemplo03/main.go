package main

import (
	"context"
	"fmt"
	"log"

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
	fmt.Println("Entendendo o funcionamento do SELECT e FOR eterno com comunicação Não Bloqueante")
	fmt.Println()
	fmt.Println("A solução sem Default e por Context em um Case")
	/**Criaremos a var de contexto, que é um tipo de canal com a soma de todos timeout +1  */
	ctx, cancel:= context.WithTimeout(context.Background(), 7*time.Second)
	if cancel == nil {
		log.Fatal("Error no context",cancel)
	}

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
		case <-ctx.Done():
			fmt.Println()
			fmt.Println("Main Terminou........")
			return

		}

	}

}
