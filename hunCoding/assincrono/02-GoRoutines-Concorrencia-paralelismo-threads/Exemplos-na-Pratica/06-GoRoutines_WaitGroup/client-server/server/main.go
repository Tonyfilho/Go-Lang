package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

/**Imagine que temos um servidor e queremos executar várias tarefas em concorrencia*/
/**Para vermos temos que instanciar mais de 1 terminal*/
/**1 NO server e outro no client do Server*/
func main() {
	fmt.Println("GoRotines WaitGroup - Ponteiros")
	/**Criaremos um Listener TCP na porta 8080 onde ficará escutando conexoes*/
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
/**Enquanto ele receber conexão vai abrir uma nova variavel de conexão e se não tiver erros dará um handle na conexão*/
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConn(conn)
	}

	///fmt.Println("Fim do programa")
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		/**Equanto tivermos conexão sendo usada teremos um FOR infinito e teremos um print Responce From Server*/
		_, err := io.WriteString(c, "responce from server\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}