package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**OBS: O Objeto STRUCT para ser convertido em JSON tem que iniciar com Upcase*/
/**Em GO TUDO que começa com UPCASE tem PERMISSÃO de Exportação e conversão*/

type FullName struct {
	Nome      string
	SobreNome string
}

type Pessoa struct {
	NomeCompleto  FullName
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {
	/**Criando Objeto para transformar para Json*/
	/** 1º Objeto anotação composta*/
	tony := Pessoa{
		NomeCompleto: FullName{
			Nome:      "Tony",
			SobreNome: "Filho",
		},
		Idade:         40,
		Profissao:     "Developer",
		ContaBancaria: 1000,
	}

	/** 2º Objeto anotação simples, tenho q por o 2º Struct FullName*/
	juan := Pessoa{
		FullName{
			"Juan",
			"Mtz",
		},
		19,
		"Developer",
		1000,
	}

	/**Criaremos as Funções Json.Marshal(), temos que ter 2 variaveis, Data e Err*/

	tonyJson, err := json.Marshal(tony)
	/**Verificação se temos error, ou se a variavel err esta VAZIA*/
	if err != nil {
		fmt.Println("Error no Marshal TonyJson: ", err)
	}
	fmt.Println("O Json by  Func Marshal TonyJson : ", string(tonyJson)) // Imprime [123 125] isto é slice of bites, tenho q fazer o CASTING
	fmt.Println()
	os.Stdout.Write(tonyJson)
	juanJson, err := json.Marshal(juan)
	/**Verificação se temos error, ou se a variavel err esta VAZIA*/
	if err != nil {
		fmt.Println("Error no Marshal JuanJson: ", err)
	}
	fmt.Println("O Json by  Func Marshal JuanJson : ", string(juanJson) ) // Senão fizer o CASTING vai imprimir um Slice of Bites [123 125] 
	fmt.Println()
	os.Stdout.Write(juanJson)

	/**1º Print usando o Print
	O Json by  Func Marshal TonyJson :  [123 125]
	O Json by  Func Marshal JuanJson :  [123 125]*/
}
