package main

import (
	"fmt"
	//"sync"
)

// v, ok := ←chan
// Se receber valor: v, true
// Canal fechado, nada, etc.: zero v, false
// Agora vamos resolver o problema do exercício anterior usando comma ok.
func main() {
	fmt.Println("#### Aula21F   Canais – v, ok := ←chan")
	fmt.Println("Lembrando que o ComaOK é usando no Make tb!, Aqui veremos o comportamento dele nos Canais")
	fmt.Println()
	fmt.Println("Se receber valor: v, true")
	fmt.Println()
	fmt.Println("Canal fechado, nada, etc.: zero v, false")
	fmt.Println()
	fmt.Println("Agora vamos resolver o problema do exercício anterior usando comma ok.")
	fmt.Println()

	canal := make(chan int)

	go func ()  {
		canal <- 42
		close(canal)
		
	}()

	/**agora depois da função faremos o ComaOK*/
	v, ok := <-canal
   /**Precisamos usar a variavel*/
   fmt.Println("Nosso ComaOK!: ", v , ok) // 42 true
   
   /**Mas se tentarmos atribuir um valor depois da 1º atribuição teremos false*/
   v, ok = <- canal
   fmt.Println("Nosso ComaOK depois da reAtribuição!: ", v , ok) // 0 false
   /**Na 2ª Vez temos ZERO , pois o Zero é valor de VAZIO e temos False para mostrar isto*/
   fmt.Println("Na 2ª Vez temos ZERO , pois o Zero é valor de VAZIO para Int e temos False para mostrar isto") // 0 false




}


