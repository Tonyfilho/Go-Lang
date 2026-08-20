package main

import (
	"fmt"
)

/** aqui temos um type e um struct , ou seja um Objeto(struct) e um tipo (type)
 */
type pessoa struct {
	nome  string
	idade int
}

/**Obs onde fica a TIPAGEM (receiver) */
func (p pessoa) oiBomDia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {

	mauricio := pessoa{"Maurício", 30}
	mauricio.oiBomDia()
	/**Em Go não conseguimos modificar o valor do receiver, apenas acessá-lo */
	/**E nem invocar o OIBOMDIA, pois tenho q criar um Objeto e invocá-lo
	Ex: oiBomDia() vai dar erro, pois não é um função e sim um método que tem uma função
	 especifica para valor tipados */

}

// func (receiver) identifier(parameters) (returns) { code }
