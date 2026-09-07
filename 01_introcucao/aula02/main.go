package main

import "fmt"


// tipo criados com a palavra reservada type
type Hotdog int
var localTipo Hotdog
// falaremos de conversão que em outras liguagens é chamado de casting
var localInt int

func main() {
	name := "Go Developers"
	//********** Interpreter literal
	fmt.Println("Tony \n"  + name)

	//**********Row string literal, OBS: isto não é a interpolação de string, apenas concatenação
	fmt.Println(`Tony ${name}` + name)


	// Tipo criado e usando printf para mostrar o tipo e valor
	localTipo = 5
	fmt.Printf("%v %T\n", localTipo, localTipo)

	//*********** Conversão de tipo (casting)
	localInt = int(localTipo)
	fmt.Printf("%v %T\n", localInt, localInt)
	// em outras liguagens consigimos ter um booleano, mas em go não conseguimos comparar tipos diferentes
	//eIqual := localInt == localTipo  
	// convertando o tipo para comparar
	eIqual := localInt == int(localTipo)
	fmt.Printf("%v %T\n", eIqual, eIqual)

}	