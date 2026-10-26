package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Aula 09 Sync.COND.")
	fmt.Println()
	fmt.Println("Neste Exemplo NÃO usaremos o Cond ")
	fmt.Println()
	
	var wg sync.WaitGroup
	/**Criaremos o nosso Once*/
	var once sync.Once
	
	load := func ()  {
		fmt.Println("Executando o Codigo do Init....")
	}
	

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			// load()
			/**E Aqui dentro simplimente chamaremos o Once.Do()*/
			/**Observe q ele rece uma func como paramentro*/
			once.Do(load)
			}()
		}
		
		/**Aqui no final o Wait()**/
		wg.Wait()
		
		fmt.Println("Fim do Main().......")
		fmt.Println()
}