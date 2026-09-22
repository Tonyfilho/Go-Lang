package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Aula do Video04 Errors em Go")
	fmt.Println()
	/*Imagine que queremos fazer um consulta em um site, um request**/
	/**Observe o retorno da func Get func http.Get(url string) (resp *http.Response, err error)*/
	/*Temos o response e o Error , o que temos de fazer é criarmos 2 variaveis, o Get retornará um Erro caso haja **/
	/*Caso haja ERROR a variavel virá com as informações de error, caso não será NIL(ou vazia)**/
	res, err :=  http.Get("https://www.google.com/")
	if err != nil {
		// daremos um Panic para abortar o Get
		panic(err)
		
	}
	fmt.Println("Nosso Get", res.StatusCode)
	fmt.Println()
	fmt.Println()
	fmt.Println("Mas se colocarmos a URL errada veremos o erro")
	fmt.Println()
	
	/*Mas se colocarmos a URL errada veremos o erro**/
	res2, err :=  http.Get("https://www.google.com20/")
	if err != nil {
		// daremos um Panic para abortar o Get
		log.Fatal("Error na 2º chamada: ",err.Error()) //no such host exit status 1
		
	}
	fmt.Println("Nosso Get", res2.Body)

}
