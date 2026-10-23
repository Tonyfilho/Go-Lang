package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Aula 08 Sync.ONCE.")
	fmt.Println()
	fmt.Println("Neste Exemplo não usaremos o Once.")
	fmt.Println()
	/**Começaremos criando WaitGroup para esperar a execução de todas as GoRoutine*/
	var wg sync.WaitGroup
	/*Criaremos uma função anonima e atribuiremos para varriavel Load**/
	load := func ()  {
		fmt.Println("Executando o Codigo do Init....")
	}
	
	/**Aqui temos um For 0 a 19 onde usaremos Add(1) para cada GoRoutine */
	/**dentro deste for teremos uma func Anonima q é uma GoRoutine , teremos o Defer p Done() */
	/**E teremos 20 GoRoutine executando a função LOAD()*/
	
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			load()
			}()
		}
		
		/**Aqui no final o Wait()**/
		wg.Wait()
		
		fmt.Println("Fim do Main().......")
		fmt.Println()
}