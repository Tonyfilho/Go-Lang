package main

import (
	"fmt"
)

/**Quando crio um Struct, tenho que TIPAR por boa pratica, por isto que usamos o type**/
/**Caso não tiparmos teremos Structs ANONIMOS*/
/**A Posição  deste OBJETO é importante para a inicialização*/

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	fmt.Println("Structs – 1. Struct")
	/**Obs ao tipo de declaração, neste caso estamos usando a sintaxe de inicialização por campos*/
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}
	/**Obs ao tipo de declaração, neste caso estamos usando a sintaxe de inicialização por posição*/
	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

	/**O Struct é uma forma de criar tipos compostos em Go,
	que permitem agrupar dados de diferentes tipos em uma única estrutura*/
	fmt.Print()
	fmt.Println("Structs – 2. Struct, Não preciso Criar uma Classe, o Struct já é uma Classe")

	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}
	/**Podemos declara iqual a um array no C#, ou um Objeto no Typescript*/
	pessoa3 := pessoa{"Mauricio", 40}
	/**Aqui temos a escrita simples ou declaração do struct pessoa1 , mas de forma simples*/
	pessoa4 := profissional{pessoa{"Vanderlei", 50}, "Político", 10000000}

	/**Acessando os campos dos structs*/
	fmt.Println(pessoa1.nome)
	/**Campos Não promovidos, observe que podemos acessar os campos normamente*/
	fmt.Println("Acessando o campo nome do struct pessoa2.pessoa.nome de forma Normal:")
	fmt.Println("Não promovido", pessoa2.pessoa.nome)
	/**Campos promovidos, observe que podemos acessar os campos da struct que estaão dentro de
	outro Structs ou Objeto, desde que não haja ambiguidade de nomes*/
	fmt.Println("Acessando o campo nome do struct pessoa2.nome de forma promovida, tiremos o Objeto PESSOA:")
	fmt.Println("Promovido", pessoa2.nome)

	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)

	fmt.Print()
	fmt.Println("Structs – 3. Struct, Anonimo, Structs Anonimos são úteis para criar tipos compostos de forma rápida e sem a necessidade de criar uma struct nomeada")

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Sou Struct Anonimo",
		idade: 50}

	fmt.Println(x)
}
