package main

import "fmt"
import "golang.org/x/crypto/bcrypt" 


func main() {
	fmt.Println("Aula 18 BCrypt")
	fmt.Println("Para usar bcrypt, instale o módulo com: go get golang.org/x/crypto/bcrypt")
	/**Criando uma senha*/
	senha := "123456"
	/**Gerando o hash da senha*/
	/**1º Criando 2 variaveis , um para o hash e outro para o erro*/
	/**2º  Temos 2 parametros 	 um []Slice de bytes e  o custo do hash , aqui usaremos o padrão
	*/
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Senha:", senha)
	fmt.Println("Hash:", string(hash))
	fmt.Println()
	fmt.Println("A função CompareHashAndPassword compara uma senha com um hash gerado \n, retornando nil se forem iguais ou um erro se forem diferentes.")
	fmt.Println()
	unHash:= bcrypt.CompareHashAndPassword(hash, []byte("1241"))
	if unHash != nil {
		fmt.Println("Senha incorreta:",err)
	}
	fmt.Println("UnHash:", unHash)	

}
