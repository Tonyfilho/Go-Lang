package main

import (
	"fmt"
	"sync"
)

// 01 Canais par, ímpar, e converge.
// 02 Func send manda pares pra um, ímpares pro outro, depois fecha.
// 03 Func receive cria 2 gos funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
// 04 Por fim um range retira todas as informações do canal converge.

func main() {
	fmt.Println("#### Aula21G   Canais Convergencia")
	fmt.Println()
    /**Resulmo
	* temos 2 canais enviandos dados  e um canal converge q recebe tudo, isto dentro das 2 funções
	* Ou seja de 2 canais foi convergido para 1, e só imprime caso o canal converge Use os dados

	*/


	/*01 Criaremos 3 canais**/
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)
	go enviar(par, impar)
	go recebe(par, impar, converge)
	/*03 temos que ter um Range para USAR as informações, caso contrario nada imprime**/
	/**O Canal só avança, se os dados forem usados*/
	for v:= range converge {
		fmt.Println("Agora funciona, pois estamos usando os dados: ", v)
	}


}

/*02 Criaremos a função send ou Enviar()**/

func enviar(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		// separaremos os partes dos impares
		if n%2 == 0 {
			fmt.Println("Numero Par: ", n)
			// envia p canal Par
			p <- n
		} else {
			fmt.Println("Numero Impar: ", n)
			// envia p canal impar
			i <- n
		}
	}
	//depois do FOR temos q fechar os canais par e impar
	close(i)
	close(p)

}

/*03 Criaremos a função receiver ou receber()*/
func recebe(p, i, c chan int) {
	// como temos 2s go routines, precisaremos de waitGroups
	var wg sync.WaitGroup
	wg.Add(2)
	// criaremos 2 go funcs, e cada uma enviará para converge seus dados ao mesmo "tempo" Concorrencia
	go func() {

		// criaremos 1 For Range
		for v := range p {
			// enviando p converge
			c <- v
			fmt.Println("Converge recebendo chanel nº: ",c)
		}
		wg.Done()

	}()
	go func() {

		for v := range i {
			// enviando p converge
			c <- v
		}
		wg.Done()

	}()
	// temos q esperar pelo WaitGroup e depois fechar os canaal Converge
	wg.Wait()
	/**O coverge usa as informações geradas nos canais par e impar, caso contrario o for n vai abastecer*/
	close(c)

}
