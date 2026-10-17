package main

import (
	"fmt"
	"time"
)

/**Uma função simples que imprime uma mensagem e simula um atraso*/

func printMessage(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(2 * time.Second) // Simula um atraso de 2 segundos
	}
}

func main() {
  printMessage("Mensagem sem GoRoutines")

  fmt.Println("Fim do programa")
}
