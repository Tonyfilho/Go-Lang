package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	fmt.Println("#### Aula21G   Canais Convergencia Exemplo 02 Rob Pike (palestra Go Concurrency Patterns)")
	fmt.Println()
	
	
	// 03º Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2))
	/**Criaremos canais 2 trabalhos e colocaremos na Função converge vai juntar e retornar um canal novo*/
	canal := converge(trabalho("Maça"), trabalho("Pera"))
	// e usamos um for para receber dados do canal var.
	for v := 0; v < 15; v++ {
		fmt.Println("A Go routina: ", <-canal, " O Mesmo End. de Memoria ", canal)
	}


}

/**01º  criar uma função trabalho para processador*/
// 01 Func trabalho cria um canal, cria uma go func que manda dados pra esse canal,
// e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		// criamos um For para mandar
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Compartinha dados diferentes : % v diz: %v: ", s, i)
			// dar um time para darmos tempo a go routine teremos um tempo do Zero ao 1000ml(1e3)
			time.Sleep(time.Microsecond * time.Duration(time.Duration(rand.Intn(1e3))))
		}

	}(s, canal)
	return canal
}

/*02º criar um função converge**/
// 02 Func converge toma dois canais, cria um canal novo, e cria duas go funcs com
// for infinito que passa tudo para o canal novo. Retorna o canal novo.
func converge (x, y chan string) chan string {
   novo := make(chan string)
    go func ()  {
		for {
			// o 1ª <- é atribuição e 2º <- é saida de valor de um canal é o mesmo q isto: "v := <-x"
			novo <- <-x
		}
	}()
    go func ()  {
		novo <- <-y
	}()


return  novo
}
/*
novo <- x
//     ↑
//     x é enviado para novo
//     Mas x é um canal, não uma string!

novo <- <-x
//     ↑  ↑
//     |  └── Recebe um valor de x (string)
//     └── Envia esse valor para novo
**/

/*
// Opção 1: Separado (mais legível)
v := <-x
novo <- v

// Opção 2: Junto (conciso mas confuso)
novo <- <-x

// Opção 3: Com range (melhor para múltiplos valores)
for v := range x {
    novo <- v
}
**/
