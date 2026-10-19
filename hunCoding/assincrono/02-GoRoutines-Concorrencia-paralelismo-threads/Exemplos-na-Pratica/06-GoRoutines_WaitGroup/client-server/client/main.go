package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	/**Criaremos a  conexão Dial para OUVIR a porta 8000*/
	/**faremos 1 conexão tcp em localhost:8000*/
	conn, err := net.Dial("tcp", "localhost:8000")

	/*Caso tenhamos error daremos um fatal**/
	if err != nil {
		log.Fatal(err)
	}

	/**usaremos o defer() para que o close() seja o ultimo a ser executado*/
	defer conn.Close()
	/**passando o terminal e a conexão*/
	mustCopy(os.Stdout, conn)

}

/**Criaremos uma função mustCopy() q temos Writer, quem é o Write é StOut, ou seja a saida
do terminal */
/*Ou seja pega o que vem lado do Server "responce from server" e Lança no Teminal do Cliente**/
func mustCopy(ds io.Writer, src io.Reader) {
if _, err := io.Copy(ds,src); err != nil {
	log.Fatal(err)
}


}