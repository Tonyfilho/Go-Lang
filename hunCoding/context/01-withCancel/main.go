package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/**como não temos contexto real, vamos usar o contexto background */
	/**Podemo usar o background() e todo() pois são conext vazios */
	ctx, cancel := context.WithCancel(context.Background())

	/**Aqui no Main colocaremos a go routine */
	go printUntilCanceled(ctx)

	/**Simulando alguma operação que precise ser cancelada */
	/**dando 5 segundos antes de cancelar, pois ficamos ouvindo o sinal de DONE() */
	/**Ou seja printaremos 5 vezes e depois cancelaremos */
	time.Sleep(5 * time.Second)

	cancel()

	/**Aguardando a go routine terminar */
	time.Sleep(1 * time.Second)
}

/*Criamores um função q será um go routine no main**/

func printUntilCanceled(ctx context.Context) {
	counter := 0
	for {
		/*Do lembrando que o select é não bloqueante  e é similar a um switch case*/
		select {
		/**Aqui dentro ficaremos ouvindo o  sinal de DONE() assim que recebermos finaliza o  contexto*/
		/**Ou seja o Done() trabalha junto com o cancel()*/
		case <-ctx.Done():
			fmt.Println("Context canceled exiting...")
			return
			/**Poremos escutar o sinal de DONE() ficaremos printando dentro do default*/
		default:
			counter++
			fmt.Println("Printing...", counter)
			time.Sleep(1 * time.Second)
		}
	}

}
