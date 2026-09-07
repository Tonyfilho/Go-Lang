package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula16 Interface Writer, entendedo a Interface Writer")
	fmt.Println("File Implementa a interface Writer")
	fmt.Println("Fmt implementa a interface Writer")
	fmt.Println("O Println(w io.Writer, a ...Interface),  ou seja Recebe um Writer")
	fmt.Println("Por isto que podemos passar o FILE para  a Função NewEnCoder() func NewEncoder(w io.Writer) *Encoder")
	fmt.Println("Ou Seja recebemos um (Ponteiro) *Encoder, para um io.Writer como Argumento e isto um fmt.Stdout().")
	fmt.Println("Com este  fmt.Stdout() conseguimos por na Tela.")
	fmt.Println()
	fmt.Println()
	fmt.Println("Resumo")
	fmt.Println("1º Temos a Interface Write que tem o Metodo Write type Writer interface { Write(p []byte) (n int, err error) }")
	fmt.Println()
	fmt.Println("2º temos uma função: func (f *File) Write(b []byte) (n int, err error) que implementa este metodo	da Interface Write")
	fmt.Println()
	fmt.Println("3º Por fim temos uma Função que recebe dados pela Interface io.Write e retorna um Pondeiro	: func NewEncoder(w io.Writer) *Encoder")
	fmt.Println()
	fmt.Println("A Grande questão é que esta Interface não é so usada Imprimir, é Implementada em Muita coisa em GO!")
	fmt.Println()
	fmt.Println("Ex: É a Função de Conexão func DIAL que implementa a Interface Conn , mas a  interface Conn implementa o metodo WRITE: ")
	fmt.Println()
	fmt.Println("Função func Dial(network, address string) (Conn, error)")
	fmt.Println()
	fmt.Println("O Seja TUDO que tem o METODO WRITE pode ser recebido como ARGUMENTO")
	fmt.Println()
	fmt.Println("Ate uma conexão de REDE é passado como WRITE")

	/**1º Temos a Interface Write que tem o Metodo Write type Writer interface { Write(p []byte) (n int, err error) } */
	/**2º temos uma função: func (f *File) Write(b []byte) (n int, err error) que implementa este metodo
	da Interface Write*/
	/**3º Por fim temos uma Função que recebe dados pela Interface io.Write e retorna um Pondeiro
	: func NewEncoder(w io.Writer) *Encoder*/

}
