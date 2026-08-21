package main

func main() {
	tony := pessoa{"Tony", "filho", 40}
	juan := pessoa{"Juan", "Mtz", 19}

	localDentista := dentista{
		pessoa:         tony,
		denteExtraidos: 50,
		salarios:       1000,
	}

	localArquiteto := arquiteto{
		juan,
		"Igrejas",
	}

	/**invocando o Metodo SerHumano que usa a Interface*/

	serHumano(localArquiteto)
	serHumano(localDentista)

}
