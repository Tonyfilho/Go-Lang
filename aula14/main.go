package main

import (
	"fmt"
)

func main() {

    fmt.Println("Ponteiros em GO ") 
    fmt.Println("OBS: ", "& Este Simbolo em GO serve para Linkar um Ponteiro de uma Variavel a Outra") 
    fmt.Println("OBS: ", "* Este Simbolo em GO serve para mostrar o valor que tem dentro do Ponteiro") 
    fmt.Println("OBS: ", "*b Chamamos esta operação de DEReference, desta forma temos o VALOR e não a REFERENCIA") 
    fmt.Println("O Ponteiro é um VALOR localizado na referencia de memoria, que pode ser CAPITURADO e feito Operações") 

	a := 10
    fmt.Println("Endereço e  Valor de A: ", &a, a) 
	b := &a
    fmt.Println("Endereço Diferente B, MAS mesmo Valor de A: ", &b, *b) 
	/*Mundando o valor de A**/
	a = 11
    fmt.Println("Endereço Diferente MAS mesmo Valor de B Atualizado: ", &b, *b) 
	/**Mudando o valor de B e fazendo o Icremento no valor armazenado na variavel A*/
	*b++
    fmt.Println("Endereço A é ligado ao B, auterando o valor de B, muda o valor de A ",  a, " É o mesmo valor de B ", *b) 

  
	
	fmt.Println("Fazendo The Reference *A em B ",*b) //The Reference é saber o que temos no Endereço usando "*"
	/**T é formatador de TYPE*/
	fmt.Printf("%T, %T\n", a, b) //int, *int  Um é um int e outro é um PONTEIRO de int *int
	fmt.Printf("int, *int  Um é um INT e outro é um PONTEIRO de *int ,") 
	fmt.Printf("Ponteiro é genericamente falando uma variavel que armazena um endereço de memoria") 

	fmt.Println(a, b)

}
