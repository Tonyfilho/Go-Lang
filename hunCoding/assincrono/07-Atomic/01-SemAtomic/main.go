package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Aula 07 Atomics.")
	fmt.Println()
	fmt.Println("Neste Exemplo não usaremos o Atomics.")
	fmt.Println()
	/**Este exemplo é muito precido com o do Mutex*/
	var counter uint64
	/*Com WaitGroup teremos certeza q todas as GoRoutines são executadas**/
	var wg sync.WaitGroup
	/**1 for fará 50 iterações*/
	for i := 0; i < 50; i++ {
		/**para cada iteração adicionaremos 1 add()*/
		wg.Add(1)
		/**E para cada iteração criaremos 1 GoRoutine*/
		go func ()  {
			/**Daremos em cada iteração o Defer DOne()*/
			defer wg.Done()
			/**E teremos outro For com 1000 iterações para cada interção do 1ºFor*/
			for c := 0; c < 1000; c++ {
				/**E aqui incremetaremos o contador*/
				counter++
			}
		}()
	}
	/**E por fim damos o Wait() para cada GoRoutine seja executada*/
	wg.Wait()

	fmt.Printf("Counter: %v\n ", counter)
	fmt.Println("Fim do Main().....")
  /**saidas sem o Atomics, podemos ver q o valor correto é 50000, mas há variações*/
	// Counter: 50000
	// Counter: 50000
	// Counter: 49442
	// Counter: 49994
 

}