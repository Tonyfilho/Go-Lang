package main


import (
	"fmt"
	
)

func main() {
	/*Ponteiros são super variaveis que armazenam endereços de memória de outras variaveis*/
	/**Quanto não usar ponteiros ? Quando queremos passar uma COPIA sem alterar o valor original*/
	/**Na verdade, normalmente passamos COPIAS e não ponteiros*/
	fmt.Println("Aula 01 Ponteiros em Go")
	var a int = 10
	var b *int = &a //b é um ponteiro que armazena o endereço de memória de a
	fmt.Printf("Valor de a: %d\n", a)
	fmt.Printf("Endereço de a: %p\n", &a)
	fmt.Printf("Valor que tem em b é o endereço de a: %p\n", b)
	fmt.Printf("Valor de b: %d\n", *b)
	fmt.Printf("Endereço de b: %p\n", &b)


}