package main

import (
	"fmt"
	"log"
	"os"
)





func main() {
	fmt.Println("Exemplo 04 - log.Fatalln")
	fmt.Println("the Fatal functions call os.Exit(1) after writing the log message ....")
	fmt.Println("Fatalln is equivalent to Println() followed by a call to os.Exit(1).")
	/**Aqui neste exemplo quando ocorre um erro, o programa é encerrado */
	/**Esta função com os.Exit() ja fica no inicio e finaliza o programa em caso de erro finaliza e  defer não roda*/
	/**os.Exit(), pára tudo seja go routine, canais tudo fecha sem choro, em caso Não erro retorna 0, em caso de erro retorna 1*/
	/**O Fatalln, é o erro onde temos que finalizar o programa , ou seja é um erro fatal */
	/**Para erros mais simples use log.Println() e trate o erro de forma apropriada */
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		/**Neste exemplo chmaremos o log.Fatalln  */
		/*Fatalln is equivalent to Println followed by a call to os.Exit(1).*/
		log.Fatalln(err)
		//		panic(err)
	}





}
func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

/*
... the Fatal functions call os.Exit(1) after writing the log message ...
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).