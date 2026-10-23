package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Aula 08 Sync.ONCE.")
	fmt.Println()
	fmt.Println("Neste Exemplo  usaremos o Mutex.")
	fmt.Println()

	var wg sync.WaitGroup
	/**Caso queiramos Usar o Mutex sem ONCE*/
	var mu sync.Mutex
	/**E precisamos criar um Flag q troque o valor e proteja*/
	var passed bool


	load := func ()  {
		fmt.Println("Executando o Codigo do Init....")
	}
	
	
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			/**Antes do IF abriremos o Lock() do Mutex*/
			mu.Lock()
			if !passed {
				load()
				passed = true
				
			}
			/**Aqui no fim damos o UnLock do Mutex*/
			mu.Unlock()
			}()
		}
		
		wg.Wait()
		
		fmt.Println("Fim do Main().......")
		fmt.Println()
}