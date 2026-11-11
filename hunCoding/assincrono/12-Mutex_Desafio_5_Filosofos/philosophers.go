package main

import (
	"fmt"
	"sync"
	"time"
)

/**1º criar a quantia de filósofos , tem q ter numero impar*/
const numFilosofos = 5

/**2º criar a quantia de refeições desejadas q cada filósofo deve comer*/
const refeicoesDesejadas = 1

/** 3º construindo o objeto filósofo*/
type Filosofo struct {
	id        int         //identificador do filósofo
	refeicoes int         //quantidade de refeições q o filósofo já comeu Ex: se estamos no filosofo 1 e ele já comeu 2 refeições, então refeicoes = 2 ou não, se ele ainda não comeu nada, então refeicoes = 0
	garfoEsq  *sync.Mutex //ponteiro para o garfo da esquerda q temos que dar um lock nele para que ninguem pegue o garfo
	garfoDir  *sync.Mutex //ponteiro para o garfo da direita q temos que dar um lock nele para que ninguem pegue o garfo
}

/**4º Temos a função pensar*/
/**Toda vez q esta pensando que dizer que ele tentou pegar os garfos e não conseguiu*/
func (f *Filosofo) pensar() {
	/**Iremos printar qual filósofo e q está pensando*/
	fmt.Printf("Filósofo %d está pensando...\n", f.id)
	/**Feito isto temos que dar um time.Sleep(1 * time.Second) para simular o tempo de pensamento, toda vez que o filósofo pensar dura 50 microsegundos*/
	time.Sleep(time.Microsecond * time.Duration(50)) //tempo de pensamento de 50 microsegundos
}

/**5º Temos o metodo comer*/
/**Toda vez q esta comendo que dizer que ele pegou os 2  garfos */
/**Lembrando que está comendo, o restante fica pensando, ouy seja os filósofos da esquerda e direita não conseguem pegar os garfos e comer*/
func (f *Filosofo) comer() {
	/**Iremos printar qual filósofo e q está comendo*/
	fmt.Printf("Filósofo %d está comendo...\n", f.id)
	/**Feito isto temos que dar um time.Sleep(10 * time.Second) para simular o tempo de refeição, toda vez que o filósofo comer dura 10 microsegundos*/
	time.Sleep(time.Second * 10)
	f.refeicoes++ //incrementando a quantidade de refeições do filósofo q esta comendo
	/**falaremos que o filósofo terminou de comer*/
	fmt.Printf("Filósofo %d terminou de comer. Refeições: %d\n", f.id, f.refeicoes)

}

/**6º Como funcionará? Cada  Filosofo irá entrar no metodo tentarComer de uma vez*/
/**usaremos go routines*/

func (f *Filosofo) tentarComer() {
	/**Criaremos um for que enquanto as refeições forem menores que refeicoesDesejadas que neste caso é 1 continuaremos tentando comer*/
	/**Ou seja quando tiver 1, o For acaba*/
	for f.refeicoes < refeicoesDesejadas { 
		f.pensar() // o filosofo irá pensar
		/**Tentando pegar o garfo da esquerda*/
		fmt.Printf("Filósofo %d está tentando pegar o garfo da esquerda...\n", f.id)
		f.garfoEsq.Lock() //dando lock no garfo da esquerda para pegar o garfo da esquerda, 
		//lembrado que se alguem ja deu lock no garfo da esquerda, o filosofo vai ficar esperando até que o garfo da esquerda seja liberado
		// e aposição será o Id
		fmt.Printf("Filósofo %d pegou o garfo da esquerda (posição %d).\n", f.id, f.id)

		/**Tentando pegar o garfo da direita*/
		fmt.Printf("Filósofo %d está tentando pegar o garfo da direita...\n", f.id)
		f.garfoDir.Lock() //dando lock no garfo da direita para pegar o garfo da direita, mas qual é o guarfor da  direita?
		//lembrado que se alguem ja deu lock no garfo da direita, o filosofo vai ficar esperando até que o garfo da direita seja liberado
		// e aposição será o Id
		// temos o Id e operador de Modulo % numFilosofos, ou seja, se o Id do filosofo for 0, então o garfo da direita será 0 +1 
		fmt.Printf("Filósofo %d pegou o garfo da direita (posição %d).\n", f.id, (f.id)%numFilosofos + 1) // no modo desta forma so pega o garfo da direita.

		/**Depois disto colocaremos para comer*/
		f.comer()

		/*depois colocaremos que esta devolvendo os garfos**/
		/**Liberando os garfos*/
		fmt.Printf("Filósofo %d está devolvendo o garfo da esquerda...\n", f.id)
		f.garfoEsq.Unlock() //liberando o garfo da esquerda
		fmt.Printf("Filósofo %d está devolvendo o garfo da direita...\n", f.id)
		f.garfoDir.Unlock() //liberando o garfo da direita

	}

	/**iremos para função main*/

}