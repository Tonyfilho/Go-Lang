package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/**como não temos contexto real, vamos usar o contexto background */
	/**Podemo usar o background() e todo() pois são conext vazios */
	/*NO DeadLine() temos que passar uma data de vencimento*/
	/**demos um time.now e adicionamos 5 segundos*/
	/**Aqui não precisamos ter um timeout()  pois o proprio WithDeadline() lida com isso*/
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	/**Colocaremos um defer para ser o ultimo a ser executado*/
	defer cancel() /// não precisamos por cancel() é questão de deadlock

	/**Chamamos a função printUntilCanceled passando o contexto*/
	printUntilCanceled(ctx)






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
			fmt.Println("Printing deadline...", counter)
			time.Sleep(1 * time.Second)
		}
	}

}
