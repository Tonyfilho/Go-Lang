package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Aula 12 Mutex - Desafio 5 Filósofos")
	fmt.Println()

	/**7º começaremos fazendo wait groups*/
	var wg sync.WaitGroup

	garfos := make([]*sync.Mutex, numFilosofos) //criando um slice de ponteiros para mutex, ou seja, cada posição do slice é um garfo
	filosofos := make([]*Filosofo, numFilosofos) //criando um slice de ponteiros para filósofos, ou seja, cada posição do slice é um filósofo

	// inicializando os garfos
	// para cada garfo criaremos um mutex, ou seja, cada garfo é um mutex
	for i := range garfos {
		garfos[i] = &sync.Mutex{} //criando um mutex para cada garfo
	}

	// inicializando os filósofos
	// para cada filósofo criaremos um objeto Filosofo, ou seja, cada filósofo é um objeto Filosofo
	for i := 0; i < numFilosofos; i++ {
		filosofos[i] = &Filosofo{
			id:        i + 1, //id do filósofo, começando do 1 e não do filosofo 0
			refeicoes: 0,     //quantidade de refeições do filósofo, começando do 0
			garfoEsq:  garfos[i], //garfo da esquerda do filósofo, ou seja, o garfo da posição i
			garfoDir:  garfos[(i+1)%numFilosofos], //garfo da direita do filósofo, ou seja, o garfo da posição i+1, mas como é circular, usamos o operador % para pegar o resto da divisão por numFilosofos
		}
	}

  /**8º iniciando as goroutines dos filósofos*/
   for i := 0; i < numFilosofos; i++ {
      wg.Add(1) //para sabermos que todos terminaram de comer, adicionamos 1 ao wait group para cada filósofo
	  // coilocaremos um go routine para cada filósofo, ou seja, cada filósofo irá tentar comer ao mesmo tempo
	  // e assim q terminar  de comer o defer wg.Done() será chamado para decrementar o wait group
	  go func(filosofo * Filosofo) {
		defer wg.Done() //defer para garantir que o wg.Done() será chamado mesmo que ocorra um panic
		filosofo.tentarComer() //chamando o método tentarComer do filósofo

	  }(filosofos[i])

   }

   /**9º esperando todas as goroutines terminarem ou seja fica preso*/
   wg.Wait()

}
