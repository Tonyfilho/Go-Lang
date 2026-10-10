package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Exemplo 03  -log.Println-")
	fmt.Println("Quando usamos a a Criação de Files temos este tipo de tratamento, func log.SetOutput(w io.Writer)")
	fmt.Println("Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.")
	/**Nesta Opção Criamos um File e salvamos o error**/
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	/**A grande diferença é que agora todos os logs serão escritos no arquivo */
	/**func log.SetOutput(w io.Writer)*/
	/**Quando usamosa a Criação de Files temos este tipo de tratamento */
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.