package main

import (
	"fmt"
	"time"
)

/*01 Criaremos uma Função Worker, ou seja algo q de trabalho ao processador**/
/* Imaginaremos como se fosse um ServidorWEB processando uma pagina**/
/**Teremos um ID e a Mensagem que será um Canal, o Worker vai ler e imprimir na tela*/
func worker(id int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker ID: ", id, " Canal de Mensagem Nº : ", res)
		// daremos um time.sleep() para dar tempo da leitura do canal
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Go Routines - Entendendo Go Routines")
	fmt.Println()
	fmt.Println("Entenderemos como um webservice consegue rodar \n com tantas requisiões e ser mais rapido do que outros , é por causa das GoRoutines!")
	fmt.Println()
	/**Cada vez que fizermos uma nova requisição, será gerado um nova GoRoutine(Thread)*/
	
	/*02 *Criaremos o Canal e chamaremos a func Worker*/
	msg := make(chan int)
	/*03 invocaremos a Func. Worker, observe q: para cada invocação o tempo de execução cai!**/
    /*Iniciaremos rodando um GoRouter e a func. Worker funcionará em Background**/
	fmt.Println("Para cada invocação o tempo de execução cai! Quanto mais Thread usarmos menor o tempo de execução")
    go worker(1, msg)
	fmt.Println("Será que se por outro Worker o tempo fica menor, Observe a minha Linha o Worker é Assyncrono!!!")
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)
	go worker(6, msg)
	go worker(7, msg)
	go worker(8, msg)
	// go worker(9, msg)


	
	/*04 temos q fazer um loop para usar estes recursos**/
	fmt.Println("Sem o For não tem entra e com isto não  saida")
	fmt.Println()
// Cada vez q o Look rodar e atribuirmos um novo valor p o canal e Worker vai rodar tb!
	for i := 0; i < 10; i ++ {
		// aqui dentro passaremos o valor do I para o canal, ou seja a entrada de dados
        msg <- i

	}

	fmt.Println("✅ O tempo ficou Menor?")
	fmt.Println("✅ Programa finalizado!")

	/*Resumo, este For de entrada, pode fazer o Look quando o canal estiver Vazio, ou seja o for fica´ra esperando **/
   /**O Worke pegará a a mensagem e esperará 1 segundo, somente depois dará permissão  for de abastecimente*/
   /*Mas a medida q formos adicionando Worker teremos mais Threads e com mais Threads mais Processamento**/
   /*Um Thread em outras liguagem custa 1 Mega de memoria, aqui no Go custa  apenas 2kbits, ou seja podemos ser 500x mais eficientes**/

}
