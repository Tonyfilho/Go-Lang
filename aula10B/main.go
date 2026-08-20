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
	/*aqui usaremos a mas podemos usar a mesma Interface pois
	temos o mesmo Metodo SerHumano que tem a Mesma Iterface, basta passar o valor Gente */
	serHumano(localArquiteto)
	//Não precisei invocar oiBomDIa(), somente  passei a variavel localArquiteto
	/**A interface usa valores de TIPOS DIFERENTES sem precisar mudar nada, 
	invocando o mesmo Metodo em TIPOS Diferentes*/

}
