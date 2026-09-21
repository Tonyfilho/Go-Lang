package main

import (
	"fmt"
)


func main() {
	fmt.Println("#### Aula21E   Canais – 4. Select # Exemplo 2:")
	fmt.Println()
	fmt.Println("Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit")
	fmt.Println()
	fmt.Println("Func 2 for infinito, select: case envia pra canal, case recebe de quit")
	fmt.Println()
	canal := make(chan int)
	quit := make(chan int)
	/**Usarei 1 go routine p não ter q usar working group*/
	go recebeQuit(canal, quit)
	enviaParaOCanal(canal, quit)
	fmt.Println()
	fmt.Println("Para perceber como a concorrência funciona \n, uma função dependendo da outra pra enviar valores, e a outra dependendo do valor quit pra encerrar")

}

/**Criando as 2 funções*/
func recebeQuit(localCanal chan int, localQuit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido Canal: ", <-localCanal)
	}
	/**enviar p localQuit*/
	localQuit <- 0
}

/**nesta função um for INFINITO */
func enviaParaOCanal(localcanal chan int, localQuit chan int) {
	/**Criamos uma var com valor a ser enviado*/
	qualquerCoisa := 1
	/**for infinito*/
	for {
		select {
		/** 1º case se tiver canal, mandaremos o valor de qualquerCoisa*/
		case localcanal <- qualquerCoisa:
			qualquerCoisa++
			/**2º case, caso recebe do quit, termino*/
		case <-localQuit:
			fmt.Println("Saiu... QUIT")
			return 
		}
	}

}
