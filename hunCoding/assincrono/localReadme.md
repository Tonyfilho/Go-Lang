#### Dados Assícronos em GoLang GouRoutines

## 01 Aula Ponteiros

https://www.youtube.com/watch?v=FLIWsfS9gTg&list=PLm-xZWCprwYQz7Di62Z93zGHyLC-HkFpb

https://github.com/HunCoding/golang-basics/tree/main/ponteiros

## 02 Aula Go Routines Concorrencia Paralelismo Threads

# Como funcionam as GOROUTINES, concorrência/paralelismo e threads - GoLang Assíncrono #go #async

https://www.youtube.com/watch?v=PMM5Fia9O4g&list=PLm-xZWCprwYQz7Di62Z93zGHyLC-HkFpb&index=3

Hoje falamos sobre as tão usadas e queridas, goroutines. Com elas temos infinitas possibilidades de utilização e de possíveis melhorias em performance ja que quando se trata de goroutines, estamos falando tambem de concorrência

## 02B Aula Colocando a mão na massa com as GOROUTINES! - GoLang Assíncrono #go #async #goroutine

https://www.youtube.com/watch?v=jDb_cinrDYA&list=PLm-xZWCprwYQz7Di62Z93zGHyLC-HkFpb&index=3

Depois do video passado super teorico, hoje aprendemos a usar as goroutines no nosso código. Com elas temos infinitas possibilidades de utilização e de possíveis melhorias em performance ja que quando se trata de goroutines, estamos falando tambem de concorrência.

https://github.com/HunCoding/golang-asynchronous

# 01Exemplo sem GoRoutines

# 02Exemplo com GoRoutines

# 03Exemplo >GouRoutines com diferentes variantes For Functions Call

# 04Exemplo Função anonima com GoRoutines

# 05Exemplo como fazer com q Varias GoRoutnes printasse o mesmo Valor ?

Neste caso usaremos ponteiros

# 06Exemplo de Client-Server

/**Imagine que temos um servidor e queremos executar várias tarefas em concorrencia ou seja varias conexoes\*/
/**Para vermos temos que instanciar mais de 1 terminal*/
/\*\*1 NO server e outro no client, e veremos o Server trabalhar*/

# No Client

/**Criaremos uma função mustCopy() q temos Writer, quem é o Write é StOut, ou seja a saida
do terminal */
/*Ou seja pega o que vem lado do Server "responce from server" e Lança no Teminal do Cliente**/

# saida

responce from server  
responce from server  
responce from server  
responce from server  
responce from server

E se tivermos VARIOS CLIENTES conectados no terminal ?

/**Se abrimos ouitro terminal no Cliente e rodarmos , ele rodará mais imprimirá Porque?\*/
/**Porque temos 1 Serve q emite 1 conexão e codigo fica travado\*/


Caso, paremos o terminal onde esta printando, o que n\ao printa passa a printar. Ou seja como so tinhamos 1 conexão e assim que ela pára, ele automaticamente cuida da outra.

# NO serve o HangleConn() esta preso dentro do FOR por Não ser uma GoRoutine
# Obs. Imagem 07
handleConn(conn)

## Como fazer com que haja mais de 1 conexão ao mesmo tempo? É ai que entra a GoRoutine
/**Agora para cada conexão criada bo SERVER dentro FOR, ele cria um GoRoutine e Espera para Criar Outra GoRoutine*/
# Obs. Imagem 07B e 7C mostra a Concorrencia funcionando pela GoRoutine
    go handleConn(conn)


## Aula03 Mutex


## Aula04 Channels 
/**Channels é a forma q temos de compartilhar dados entre as GoRoutines*/
/**Channels  é usado para organizar a execução de GoRoutines*/
/**No Exemplo de GoRoutines e Ponteiros tinhamos a situação q assim que mudarmos o valor do Ponteiro era atualizada a GoRoutines, o problema desta abordagem é que se o dado vem de fora da aplicação Ex: uma API teremos problema de time, ou seja a GoRutines passa a ficar bloqueante*/

/**Em Channels podemos fazer q um GoRoutine ESPERE o valor enquando não foi abastecido, o que acontece é que a GoRoutine q envia, manda o valor para Channel e com isto não precisa esperar e ficar bloqueada*/

/**Na Imagem 01C temos 1 channel de string com buffered 1, a variavel é instanciada com a função MAKE() e enviaremos uma mensagem como Exemplo de "ping". Em outra lugar da aplicação em outra GoRoutine poderemos receber esta mensagem. Ou seja tenho 2 GoRoutines e no meio um Canal() comunicando as 2s*/
#  OBS: Muito parececido com Pipe do RXJS
# Ou Seja para eu passar um valor de uma GoRoutine para outra e o valor ficar sempre atualizado mandamos o valor pelo Canal q no final deste Canal teremos um outra GoRoutine esperando algo. Então temos Entrada de Dados e Saida de Dados.

# Channels Directions imagem 02  Canal <-"Dados"  ou chTest := <- Canal
# Obs:Podemos ter Channels normais ou especificar o tipo
/**com o operador "<-" podemos determinar se é entrada(Envio) de dados ou saida (Recebimento) QUANDO usamos como PARAMETRO de FUNÇÂO */
# OBS: O Channels tem as funções READ e WRITE
/**Desta forma podemos determinar q o paramentro SO ENVIA ou SÓ RECEBE, imagine q queremos boquear q algume coloque valor mas somente receba!*/

# Seta depois do Channel <- "dados" quer dizer q so pode INSERIR ou ATRIBUIR dados (CHANNEL ENVIAR o VALOR)

# Seta antes do localVariavel <-Channel quer dizer q so pode RECEBER dados (CHANNEL RECEBER o VALOR) ou seja a variavel RECEBE o dado de CHANNEL
/**Imagine q queira criar o Package ou um biblioteca e queira compartilhar onde quem for usar só pode RECEBER ou seja LER nunca Gravar*/