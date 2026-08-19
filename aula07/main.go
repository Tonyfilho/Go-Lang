package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps - Aula 07 inicio")
	/**Temos uma chave do tipo string e um valor do tipo int*/
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"], "\n\n")

	/***Quando não existe a chave o valor será o zero value do tipo do valor*/

	// comma ok idiom
	/**Nesse caso, a variável será receberá o valor da chave "fantasma"
	e a variável ok receberá o valor booleano indicando se a chave existe
	Ou seja temos q ter 2 variaveis, uma para o valor e outra para o booleano
	*/
	if será, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!") //não tem!
	} else {
		fmt.Println(será)
	}

	fmt.Println()
	fmt.Println("Maps - Aula 02 Ranges e Delete")

	/**Range em Maps*/
	for chave, valor := range amigos {
		fmt.Printf("Chave: %s, Valor: %d\n", chave, valor)
	}

	/**Delete em Maps*/
	delete(amigos, "alfredo")
	fmt.Println(amigos)


	
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	total := 0

	/**O range em maps permite iterar sobre todas as chaves e valores*/
	/**poderia ficar assim key, value := range qualquercoisa, usar _ é uma forma de ignorar o valor*/
	for key, _ := range qualquercoisa {
		total += key
	}
	
	fmt.Println(total)

	/**Usamos a palavra reservada DELETE e passamos o Map e a chave que queremos remover*/
	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)

}
