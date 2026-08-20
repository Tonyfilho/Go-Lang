package main

import (
	"fmt"
)


/**Criação dos Structs */
/**Aqui usamos TIPE + Nome + STRUCT */
type pessoa struct {
	nome  string
	sobrenome string
	idade int
}


type dentista struct {
	pessoa
	dentesArrancados int
	salario float64
}

type arquiteto struct {
	pessoa
	obrasconcluidas string
	salario float64
}

/**Criação das Funções oiBomDiaAula10() que o Arquiteto e o Dentista  irão ter */

func (d dentista) oiBomDia() {
	fmt.Println("Meu nome é: ", d.nome, " e eu já arranquei: ", d.dentesArrancados)
	// ou
	//fmt.Printf("Oi, bom dia! Meu nome é %v e eu já arranquei %d dentes!\n", d.nome, d.dentesarrancados)
}

   // Explicação ******** VERBO ******* Quando usar cada um?
    // Use %d quando quiser garantir que é um número inteiro (mais específico)
    // Use %v quando quiser flexibilidade ou não souber o tipo exato
func (a arquiteto) oiBomDia() {	
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já fiz %d obras!\n", a.nome, a.obrasconcluidas)
}

/**Criação da Interface Gente onde temos o metodo oiBomDia() como casca de Ovo*/
/**Aqui usamos TIPE + Nome + INTERFACE */
type gente interface {
   oiBomDia()
}

func serHumano(g gente) {
	g.oiBomDia()
}





