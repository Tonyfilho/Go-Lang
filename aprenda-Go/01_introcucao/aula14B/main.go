package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros em GO Pra que servem ?")
	fmt.Println("Poteiro servem para 2 coisas:")
	fmt.Println("««««»»»»: 1º Quando lidamos com muitos dados e queremos que todos acessem o mesmo local e Memoria.")
	fmt.Println("Ou seja. Desta forma todos que querem acessar ou mudar algo do mesmo, não precisa Criar novamente.")
	fmt.Println("Em Go tudo é pass by value!, Ou seja não passamos valores e som ponteiros, pois a performace é maior!")
	fmt.Println("Ex. Imagine um Função que recebe um Valor, ao invez de passar o valor, passamos o Ponteiro, ou seja vai lá e pegue naquele local")
	fmt.Println("Desta forma ao invez de fazer COPIA nossa passamos o ENDEREÇO dos dados")
	fmt.Println("««««»»»»: 2º É a quando queremos mudar o valor de um ENDEREÇO, sem fazer COPIAS")
	fmt.Println()
	fmt.Println()	
	fmt.Println()
	
	x := 11
	fmt.Println("Valor de X Antes da Func estaRecebeOValorDoPonteiro(x) : ", x)
	
	estaRecebeOValorDoPonteiro(x)	
	
	fmt.Println("Valor de X mantem se ORIGINAL e continua com 11 veja: ", x)
	
	/**Recebendo valor do Ponteiro, tenho que passar o ENDEREÇO &x*/
	estaRecebeUmPonteiro(&x)
	fmt.Println("Agora estamos incrementando pois usamos a Referencia sendo Incrementado NÃO é o X e sim a func estaRecebeUmPonteiro(&x) : ", x)
	
	fmt.Println("Valor de X Incrementado do Ponteiro: ", x)
	
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Resulmo o Não Ponteiro faz um copia do original, Mas o Ponteiro modifica tudo")
	fmt.Println("Quando passo um poteiro ou seja um Endereço, e fazer mudança, eu mudo todos, pois não estamos passando COPIAS")

}

func estaRecebeOValorDoPonteiro(x int) {
	/*Aqui estamos modificando o X do PARAMETRO */
	fmt.Println("Estamos dentro da Func estaRecebeOValorDoPonteiro(x int), não consigo alterar o X original")

	println("Aqui estamos fazendo uma COPIA do X ORIGINAL, por isto que não altera o valor do x:=11, continua 11")
	x++
	fmt.Println(" COPIA do X incrementada vale 12 veja: ", x)

}

func estaRecebeUmPonteiro(x *int) {
	/**Aqui recebemos o Ponteiro e modificamos valor */
	println("Aqui recebemos o Ponteiro e modificamos valor do Ponteiro já não é uma copia")
	*x++
	fmt.Println("Na função: estaRecebeUmPonteiro(x *int)", *x)
}
