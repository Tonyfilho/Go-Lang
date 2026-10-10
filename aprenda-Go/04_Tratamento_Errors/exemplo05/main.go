package main

import (
	"fmt"
	"log"
	"os"
)




func main() {
	fmt.Println("Exemplo 05 - print e log")
	fmt.Println("Panicln is equivalent to Println() followed by a call to panic().")
	fmt.Println(" Fatalln is equivalent to Println() followed by a call to os.Exit(1).")
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)

		/**Temos neste ex: o Log.Panicln, temos exit.status 2*/
		/*Com Panic funções com defer rodam ou seja, o defer será executado diferente do os.Exit*/
		/**Caso enhamos 10 defer, todos serão executados e retornarão valores caso não haja panic neles*/
		/**Podemos controlar usando recover() */
		log.Panicln(err)
		//		panic(err)
	}
	



}