package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Aula 19 Especial - Entendendo Go Routines")
	fmt.Println()
	/**temos um exemplo mais simples dentro do main, é o mesmo do outro codigo*/
	fmt.Println()
	fmt.Println("neste exemplo estamos compartinhando variavel QUEUE entre o escopo de \n MAIN{} e o escopo da Go Routine Função expressa {}()")
	fmt.Println()

	queue := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()

	/**O For da linha 18 só vair funcionar, depois q o Range usar o valor no fmt.Printl()*/
	/**A Cada vez que este range iterá com o X, ou seja consumir, o FOR da linha 21 atribui um valor*/
	for x := range queue {
	//	time.Sleep(time.Second)
		fmt.Println("Valor Consumido do Canal: ", x)
	}

}
