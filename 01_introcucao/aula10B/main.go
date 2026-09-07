package main

import (
	"fmt"
)

func main() {

	/**Criamos o Objeto Pessoa */
	tony := pessoa{"Tony", "Filho", 30}
	juan := pessoa{"Juan", "Mtz", 19}

	/**Criando */
	localDentista := dentista{
		tony,
		20,
		1000,
	}

	fmt.Println("Nome do dentista: ", localDentista.pessoa.nome)
	//cada objeto é de um tipo diferente
	localDentista.oiBomDia()
	/*aqui usaremos a mas podemos usar a mesma Interface pois
	temos o mesmo Metodo SerHumano que tem a Mesma Iterface, basta passar o valor Gente */
	serHumano(localDentista)
	//Não precisei invocar oiBomDIa(), somente  passei a variavel localDentista

	localArquiteto := arquiteto{
		juan,
		"Predios",
		2000,
	}

	fmt.Println("Nome do Arquiteto: ", localArquiteto.pessoa.nome)
	//cada objeto é de um tipo diferente
	localArquiteto.oiBomDia()
	/*aqui usaremos a Interface q por sua vez está dentro do  Metodo SerHumano(g gente)
	 e temos 2 STRUCTs que usam o mesmo NOME do Metodo que temos que implementar da Interface
	 , basta passar o valor que temos na Interface Gente. Observe que: Não tenho ftm.print() sendo invocado pelo 
	 Metodo serHumando(g gente), mesmo assim conseguimos imprimir, ou seja: Em Go não precisamos
	  IMPLEMENTAR as Interfaces , elas são automaticamente Implementadas, basta usar o mesmo nome 
	  de Metodo.
	Não precisei invocar oiBomDIa(), somente  passei a variavel localArquiteto*/
	serHumano(localArquiteto)
	
	/**A interface usa valores de TIPOS DIFERENTES sem precisar mudar nada, 
	invocando o mesmo Metodo em TIPOS Diferentes*/

}
