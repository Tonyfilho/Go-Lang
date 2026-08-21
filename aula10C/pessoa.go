package main

import (
	"fmt"
)

/**Criação da Interface*/
type gente interface {
	oiBomDia()
}

/**Criação do Metodo que usará a Interface indiretamente ou implicitamente no seu Argumento*/
/**Podemos criar o switch por tipo  Ex: queso salvar no Db, será um tipo, quero salvar no Hd, será um tipo, etc*/
func serHumano(g gente) {
	g.oiBomDia()
	switch g := g.(type) {
	/**Por causa do AutoImplemento temos acesso ao structs*/
	case dentista:
		fmt.Println("Eu ganho: ", g.salarios)
	case arquiteto:
		fmt.Println("Eu construo: ", g.tipoDeConstrucao)

	default:
		fmt.Println("Profissão não achada!")
	}
}

/**Criação do Structs*/
type pessoa struct {
	nome      string
	sobreNome string
	idade     int
}
type arquiteto struct {
	pessoa           // desta forma tenho acesso ao modo inteligente de Rest ... ou Destruction ... bastando passar variavel.nome
	tipoDeConstrucao string
}

type dentista struct {
	pessoa         pessoa // desta forma o modo inteligente do ... Rest ou Destructions não funciona, tenho q fazer variavel.pessoa.nome
	denteExtraidos int
	salarios       float64
}

/**Criação dos Metodos para imprimir Dentista e Arquiteto*/

func (a arquiteto) oiBomDia() {
	fmt.Println("Meu nome é", a.nome, "e ouve só: Bom dia!")
}
func (d dentista) oiBomDia() {
	fmt.Println("Meu nome é", d.pessoa.nome, "e eu já arranquei", d.denteExtraidos, "dentes, e ouve só: Bom dia!")
}
