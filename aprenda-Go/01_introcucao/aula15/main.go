package main

import (
	"fmt"
	"os"

	"encoding/json"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	/**O Retorno é Multiplo, temos que ter 2 variaveis 1 para Data e outra para Error**/
	/**b de BITES e err de ERROR, invocamos a Função Marchal e passamos o STRUCT GROUP como Argumento*/
	b, err := json.Marshal(group)

	/**A Variavel ERR só tem atribuição de valores caso tenha ERROR, caso contrario ela fica NIL */
	/**Em resumo: nil em Go é um valor válido que indica "não aponta para nada ou VAZIO", enquanto undefined
	 em JS é a ausência de valor propriamente dita. São conceitos diferentes com propósitos diferentes!*/
	if err != nil {
		fmt.Println("error:", err)
	}
	/** "OS" equivale ao Print, Stdout ou StandOut é o terminal tipo console.log() do Js */
	os.Stdout.Write(b)

	/**No terminal: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}*/

	/**temos o https://mholt.github.io/json-to-go/  que converte Json para Go*/
}
