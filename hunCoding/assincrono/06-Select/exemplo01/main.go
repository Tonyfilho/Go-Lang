package main

import (
	"fmt"
	"time"
)

/**Criaremos um metodo que pega o valor de Channel que recebe valor somente**/
/**Supondo q seja uma chamada em API externa, q demora justamente o time.Duration */
/**e logo apos isto temos o valores em nosso  channel*/
/**Supondo que temos 3 APIS e quem responder 1º fica com a conexão*/
func getValueFromChannelAPI(ch chan<- string, duration time.Duration) {

	/**Aqui é delay para atribuirmos valor ao channel*/
	time.Sleep(duration)
	/**aqui tribuiremos o valor ao Channel*/
	ch <- "Terminamos por aqui!"
}

/**Recebendo sinais de varios Channels e selecionando*/
func main() {
	fmt.Println("Aula06 Select em  GoRoutines")
	fmt.Println("Entendendo o funcionamento do SELECT")
	fmt.Println()
	/**1º passo criar um var local q é um channel*/
	/**2º Passo fazer chamada da suposta API 3x*/
	ch1 := make(chan string)
	go getValueFromChannelAPI(ch1, 2*time.Second)
	ch2 := make(chan string)
	go getValueFromChannelAPI(ch2, 2*time.Second)
	ch3 := make(chan string)
	go getValueFromChannelAPI(ch3, 2*time.Second)

	/**Não podemos forçar a execução chamando a variavel no print quero saber quem termonou 1º*/

	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)
	// fmt.Println(<-ch3)
	/*3º Como saber qual das 3 APIS retornam?  Ou melhor quem foi a 1ª ??**/
	/**Usaremos o SELECT a 1º GoRoutine q retornar ele termina*/
 
	select {
	case localReturno := <-ch1:
		fmt.Print("Canal 01 retornou: ", localReturno)
	case localReturno := <-ch2:
		fmt.Print("Canal 02 retornou: ", localReturno)
	case localReturno := <-ch3:
		fmt.Print("Canal 03 retornou: ", localReturno)
	}
	fmt.Println()
	fmt.Println("Main Terminou........")

}
