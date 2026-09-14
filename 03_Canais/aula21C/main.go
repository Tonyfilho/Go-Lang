package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula21B. Assignment/conversion ")
	fmt.Println("01 geral pra específico ")
	/**Quando temos um canal bidirecional, podemos convertê-lo para um canal de recepção
	ou envio, mas não o contrário. */
	fmt.Println("Quando temos um canal bidirecional, podemos convertê-lo \n para um canal de recepção 	ou envio, mas não o contrário.")
	fmt.Println()

	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr sou canal de recepção\t%T\n", cr)
	fmt.Printf("cs sou canal de envio\t%T\n", cs)

	// general to specific converts
	fmt.Println("-----")
	fmt.Println("Pegando um canal bidirecional e convertendo para um canal de recepção ou envio")
	fmt.Println()
	fmt.Printf("c\t%T Convertendo para canal de RECEIVE recepção <-chan int; \n", (<-chan int)(c))
	fmt.Println()
	fmt.Printf("c\t%T Convertendo para canal de SENT envio chan<- int; \n", (chan<- int)(c))

	fmt.Println()
	fmt.Println("02 específico pra específico ")
	/*Aqui faremos a tribuição comum de variaveis,e teremos error**/
	fmt.Println("specific to specific doesn't assignment Ex: cs = cr \n cannot use cr (variable of type <-chan int) as chan<- int value in assignment")
	// cs = cr // cannot use cr (variable of type <-chan int) as chan<- int value in assignment

	fmt.Println()
	fmt.Println("03 específico pra geral ou bidirecional ")
	fmt.Println("specific to general doesn't convert")
	/**Não conseguimos converter um canal específico para um canal geral**/
	fmt.Println("ERROR: cannot convert cr (variable of type <-chan int) to type chan int")
	fmt.Println("ERROR (chan int)(cr)")
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	fmt.Println()
	fmt.Println("ERROR (chan int)(cs)")
	// fmt.Printf("c\t%T\n", (chan int)(cs))

	fmt.Println()
	fmt.Println("04 Atribuição de tipos diferentes dda ERROR ")
	fmt.Println("Não conseguiremos atribuir canais de tipos diferentes")
	fmt.Println("ERROR c = cr")
	fmt.Println("ERROR c = cs")
	//c = cr  // specific to general doesn't assign
	//	c = cs // specific to general doesn't assign

}
