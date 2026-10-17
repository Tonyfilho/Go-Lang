package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Aula de Mutex onde temos uma variável compartilhada entre varias goroutines em for \n,	 Não usaremos o mutex neste exemplo.")
	fmt.Println()

	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		// criando uma goroutine para incrementar o contador
		go func() {
			defer wg.Done()
			for j := 0; j < 5000; j++ {
				counter++
			}
		}()
		wg.Wait()

	}
	fmt.Printf("Valor final do contador: %v\n", counter)
	fmt.Println()
	fmt.Println(" Agora todas as go routines incrementaram o contador com valor diferente para cada compilação .")

}
