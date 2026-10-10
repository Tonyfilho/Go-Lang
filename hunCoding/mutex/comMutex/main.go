package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Aula de Mutex onde temos uma variável compartilhada entre varias goroutines em for \n,	 usaremos o mutex neste exemplo.")
	fmt.Println()

	var counter uint64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 20; i++ {
		wg.Add(1)
		// criando uma goroutine para incrementar o contador
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {

				// então para cada incremento do contador, vamos travar o mutex para que apenas uma goroutine possa acessar a variável compartilhada por vez.
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
		wg.Wait()
		
	}
	fmt.Printf("Valor final do contador: %v\n", counter)
	fmt.Println()
	fmt.Println(" Agora todas as go routines incrementaram o contador com o mesmo valor.")
	fmt.Println(" Mas não aconteceu, pois não usamos o mutex para evitar condições de corrida.")
}
