package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Aula17 O Pacote Sort")
	/**1º Criando o Slice of String*/
	ss := []string{"Go", "Java", "Js", "TypeScript", "C#"}
	fmt.Println("Slice String Deserdenado: ", ss) //[Go Java Js TypeScript C#]
	si := []int{10, 9, 8, 7, 5, 6, 4, 1, 2, 3}
	fmt.Println("Slice Int Deserdenado: ", si) //[10 9 8 7 5 6 4 1 2 3]

	/**2º Ordenando o Slice*/
	sort.Strings(ss)
	sort.Ints(si)
	fmt.Println()
	/**3º Slice Ordenado*/
	fmt.Println()
	fmt.Println("Slice Ordenando: ", ss) //[C# Go Java Js TypeScript]
	fmt.Println("Slice Ordenando: ", si) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println()

}
