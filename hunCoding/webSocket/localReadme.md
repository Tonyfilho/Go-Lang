#### Criação de webSocket em Go
https://github.com/HunCoding/golang-basics/tree/main/websocket_chat_app
https://www.youtube.com/watch?v=GCB30XTcqAQ

# Criaremos uma aplicação onde abriremos navegadores diferente e teremos um Chat online
Ao mandar mensagem, o outro navegador recebe e vice e vesa


# 1º Criar nosso objeto Struct que tem 2 atributos no Models

# 2º fazer o download para projeto do package Gorilla 
no terminal instalar e importar no projeto o Package gorilla
go get github.com/gorilla/websocket


# 3º Criar 2 Variaveis no diretorio Web
1º Broadcast onde cuidará das mensagens será um chanel
2º Clients que terá um Map onde a chave é um *ponteiro e valor é boolean

# 4º Criar uma variavel Upgrader 
temos que definir os buffers de leitura e escrita com tamanho 1024


# 5º Criar 2 funções 
# ela recebe a variavel upgrader
handleMessage()

handleConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header)
em nosso caso não precisamos o Header, então ficará Nil vazio
o connection é que vai gerenciar quem for querer uma rota do websocket


Resumindo
for range canal: itera até o canal fechar. É o que você quase sempre quer.

for {} + <-canal: loop infinito, precisa tratar ok do canal, senão bug.

for range map: iteração normal sobre o map; deletar durante a iteração é permitido e seguro.

# 6 com tudo criado podemo ir agora para Main()

Ela ficará ouvido o front