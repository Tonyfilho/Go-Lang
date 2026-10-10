package web

import (
	"log"
	"net/http"

	"webSocket/models"

	"github.com/gorilla/websocket"
)

/**03 */
var Broadcast = make(chan models.Message)
var Clients = make(map[*websocket.Conn]bool)

/**04*/
/*definiremos o tamanho dos buffers de leitura e escrita**/
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*05 o connection é que vai gerenciar quem for querer uma rota do websocket**/
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	/*ela recebe a variavel upgrader**/
	/**aqui nos paramentro do upgrader, no lugar do header, passaremos Nulo pois não precisamos de nenhum header especial*/
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		/**tenho que fazer o erro, por enquanto ficará assim*/
		log.Println("error occurred while upgrading connection", err)
		log.Panicln(err) /// é a melhor opção para encerrar a aplicação, pois se não conseguir fazer o upgrade, não tem como continuar
	}
	/* daremos o defer para fechar a conexão, somente quando a função terminar */
	defer ws.Close()
	/*agora colocaremos no Map de clientes e true pois tem dados e ja esta ativa a conexão*/
	/**isto será salvo em um map de conexão**/
	Clients[ws] = true
	/**Criaremos um For infinto ate que a conexão seja fechada ou um erro*/
	/* e pegaremos o valor do Form que vem do front*/
	for {
		var msg models.Message
		/**aqui pegaremos a mensagem do front, e o ReadJSON vai ler o json e colocar na variavel msg*/
		err := ws.ReadJSON(&msg) // passaremos a referencia da variavel msg, pois o ReadJSON vai preencher a variavel msg com os dados do front
		if err != nil {
			log.Println("error occurred while reading json", err)
			/**Caso exista um erro faço um delete do nosso Map de clientes , ou seja tire esta conexão da lista*/
			delete(Clients, ws) // se der erro, removeremos a conexão do map de clientes
			break               // e sairemos do for
		}
		/**se não tivermos erro, processaremos a mensagem no Canal*/
		Broadcast <- msg // enviamos a mensagem para o canal de broadcast

	}
}

/*05**/
/**agora temos que consumir o canal criado*/
func HandleMessages() {
	/**faremos um For infinito onde  o canal Broadcast será consumido**/
	for msg := range Broadcast {
		/**enviamos a mensagem para todos os clientes conectados ou seja o Frontend*/
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				/** a mesma coisa, se erro for diferente de null deletamos o cliente do map */
				log.Println("error occurred while writing json", err)
				/* fechamos a conexão do cliente com problema, antes do delete */
				/**desta forma evitamos q um go routine tenha dados mortos*/
				client.Close()
				/**Removendo o cliente do map */
				delete(Clients, client)
			}
		}
	}

}
