package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Exemplo 02 - log.Println")
	fmt.Println("Package log implements a simple logging package ... \n writes to standard error and prints the date and time of each logged message ...")
	fmt.Println("log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.")
	/* A diferença é adição de data e horario no print**/
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
