package main

import (
	"encoding/json"
	"fmt"
)

/**1º Passa temos que  criar um Struct e termos as CRASE `.....` em volta
das dos dados para desserelização, pois ja estamos usando ASPAS DUPLAS "...." */

type Informacoes struct {
	Nome      string `json:"Nome"`
	SobreNome string `json:"SobreNome"`
	Idade int `json:"Idade"`
	Profissao string `json:"Trabalho"`
	ContaBancaria float64 `json:"ContaBancaria"`
}

func main() {
	fmt.Println("Aula15C Recebendo Json e Criando STRUCTs")
	/** 2º Criaremos um SLICE of Bite onde termos DUMMY do Json*/
	/**OBS: Uso de Tags Encoder usando a CRASE  `...` por causa do Uso de ASPAS DUPLAS, evitando conflito.
	* "NomeDoCampo":"Dados Caso seja STRING", se não for Ex: "ContaBancaria":1000
	*/
	/**Obs: Ex de uso das Tags Encoder no Campo Profissão que recebe Trabalho*/
    DummyJsonTony := []byte(`{"Nome":"Tony", "SobreNome":"Filho", "Idade": 40, "Trabalho":"Developer", "ContaBancaria":1000}`)
    DummyJsonJuan := []byte(`{"Nome":Juan, "SobreNome":"Mtz", "Idade": "19", "Trabalho":"Developer", "ContaBancaria":1000}`)
    
	/**3º Criaremos um variavel onde receberemos os dados do UnMarshal, a TIPAGEM tem q ser IQUAL*/
	var TonyUnMarshal Informacoes
	/**4º Criaremos  uma Variavel de Error onde receberemos a Função Unmarshal(data []byte, v any) error
	com 2 paramentros e receberemos 2 argumentos*/
	/**OBS quando passarmos na FUNÇÃO a Variavel q receberá o UnMarshal, receberá um PONTEIRO "&" da Função*/
    err := json.Unmarshal(DummyJsonTony, &TonyUnMarshal)

	/**5º Criaremos o IF de checagem de VAZIO*/
	if err != nil {
		fmt.Println("Error no UnMarshal: ", err)
	}
	fmt.Println("Nosso UnMarshal de Informações TonyUnMarshal: ", TonyUnMarshal)
	fmt.Println("Nosso Nome: ", TonyUnMarshal.Nome)
	fmt.Println("Nosso Idade: ", TonyUnMarshal.Idade)
	fmt.Println()
	fmt.Println()
	
	/**Exemplo 2 com DummyJsonJuan onde CAUSAREMOS Erro no Nome e na Idade,
	por causa "Aplicação Errada das ASPAS"*/
	
	var JuanUnMarshal Informacoes
	err = json.Unmarshal(DummyJsonJuan, &JuanUnMarshal)
	if err != nil {
		fmt.Println("Error no UnMarshal DummyJsonJuan: ", err)
	}
	fmt.Println("Nosso UnMarshal de Informações JuanUnMarshal: ", JuanUnMarshal)
	fmt.Println("Nosso Nome: ", JuanUnMarshal.Nome)
	fmt.Println("Nosso Idade: ", JuanUnMarshal.Idade)




}
