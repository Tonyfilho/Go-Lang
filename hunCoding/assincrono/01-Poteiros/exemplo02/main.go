package main


import (
	"fmt"
	
)

func main() {
	/*Criaremos 2 funcções para dar exemplo de referencia de memoria*/

	fmt.Println("Aula 01 Ponteiros em Go")
	fmt.Println("Dando exemplo de referencia de memoria")

	var testValue string = "Valor original"
	fmt.Printf("Valor da variável testValue: %s\n", testValue)
	
    copiaDeValor(testValue)
	referenciaDeMemoria(&testValue)
	fmt.Printf("Valor da variável testValue depois da função referenciaDeMemoria: %s\n", testValue)
	



}
/**1º é uma função normal onde recebe uma cópia*/
func copiaDeValor(valorCopiado string) {
	fmt.Printf("Valor recebido na função copiaDeValor: %s\n", valorCopiado)

}

func referenciaDeMemoria(valorReferenciado *string) {
	*valorReferenciado = "Valor alterado"
	fmt.Printf("Valor alterado na função referenciaDeMemoria: %s\n", *valorReferenciado)
}

