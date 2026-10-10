package main

import (
	"fmt"
	"log"
	"net/http"
	"webSocket/web"
)

func main() {
	/**aqui no nosso Main terá papel de controler onde receberemos  o front*/
	/*Criaremos a variavel para ouvir as requisições vindo do front ni arquivo index.html*/
	/**aqui etmos que passar um ROOT ou seja o diretorio onde estão o frontend*/
	fs := http.FileServer(http.Dir("./frontEnd"))
	/**Agora direcionaremos o cliente para frontend*/
	/*Quando alguem acessar o / ele chama o frontend**/
	http.Handle("/", fs)
	/*Quando alguem acessar /talkabout-tony ele chama o HandleConnections**/
	http.HandleFunc("/ws", web.HandleConnections)
	/* O proximo Passo é colocar o HandleMessage**/
	/*Precisamos de Assincrono data e temos que usar um GoRoutine**/
	go web.HandleMessages()

	fmt.Println("Application up and running")
	/**Agora vamos iniciar o servidor na porta 8081**/
	 if  err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
		log.Panic(err)
	}


}
