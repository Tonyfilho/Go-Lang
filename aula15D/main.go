package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome          string
	SobreNome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {
	fmt.Println("Aula15D Função NewEncoder")

	toniFilho := Pessoa{
		"Tony",
		"Filho",
		40,
		"Developer",
		1000.00,
	}

	/**Qual a diferença entre Marshal e a Função NewEncoder e NewDecoder ?*/
	/** 1º Aqui não preciso criar uma variavel V para receber(Salvar em uma varivel intermediaria)
	 a decodificação e nem o Err, tenho que que passar
	uma interface, neste nosso caso Stdout() para imprimir, Lembre da questão da Interface, onde NÃO
	preciso implementar, ela AUTO IMPLEMENTA bastando ter o mesmo nome de Metodos */
	/** 2º não precisei criar IF o erro já é gerenciavel pela Função*/
	/** 3º A Função NewEnconder, ja faz a conversão direto, sem precisar de Variaveis, ela usa Interface q ja existem Ex: StOut()*/
	/*
			 NewDecoder retorna um novo decodificador que lê de r.
		O decodificador introduz seu próprio buffer e pode ler dados de r além dos valores JSON solicitados.
	*/
	/**Resumo, todo Encoder que eu fizer é para IR direto para Interface os.Stdout, sem variaveis*/
	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(toniFilho) //{"Nome":"Tony","SobreNome":"Filho","Idade":40,"Profissao":"Developer","ContaBancaria":1000}
	fmt.Println("Resumo, todo Encoder que eu fizer é para IR direto para Interface os.Stdout para este exemplo, sem variaveis")
	
	
}
