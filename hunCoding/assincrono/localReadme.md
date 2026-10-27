#### Dados Assícronos em GoLang GouRoutines

#### 01 Aula Ponteiros

https://www.youtube.com/watch?v=FLIWsfS9gTg&list=PLm-xZWCprwYQz7Di62Z93zGHyLC-HkFpb

https://github.com/HunCoding/golang-basics/tree/main/ponteiros

## 02 Aula Go Routines Concorrencia Paralelismo Threads

# Como funcionam as GOROUTINES, concorrência/paralelismo e threads - GoLang Assíncrono #go #async

https://www.youtube.com/watch?v=PMM5Fia9O4g&list=PLm-xZWCprwYQz7Di62Z93zGHyLC-HkFpb&index=3

Hoje falamos sobre as tão usadas e queridas, goroutines. Com elas temos infinitas possibilidades de utilização e de possíveis melhorias em performance ja que quando se trata de goroutines, estamos falando tambem de concorrência

#### 02B Aula Colocando a mão na massa com as GOROUTINES! - GoLang Assíncrono #go #async #goroutine

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


#### Aula03 Mutex 
https://github.com/HunCoding/golang-asynchronous/tree/main/mutex
# imagem 01
/**Usaremos Mutex quando queremos compartilhar a mesma variavel com varias GoRoutines e ter o valor iqual a todas, muito parecido com Ponteiro e GoRoutines*/
/**Imagine 1 variavel normal sem ser o channel e 10 goRoutines ficam escrevendo e lendo concorrentemente, é Mutex q garante a segurança da atruibuição de valores, pois sem ele concorrentemente cada GoRoutine teria um valor diferente*/
/**O que o Mutex faz: Ele fecha(Bloqueia) a variavel impedindo que outra GoRoutine leia enquanto alguem estiver atribuindo valor nesta variavel, com isto é feito um bloqueio onde somente 1 GoRoutine por escrever e ler por vez.*/

# Imagem 01B
/*Qual a diferença entre Channel e Mutex?, pois o Mutex parece com channels, mas não são*/
/**Channels passam copias de valores*/
/**Channels trabalham com GoRoutines*/
/**Channels trabalham com Assícronos usando Buffer entre GoRoutines*/

/**MUTEX, trabalha com Caches*/
/**MUTEX, trabalha com Estado de uma variavel*/
/**MUTEX é um Concurrent Safe fazendo que um varial Mutex poder ser ALTERADA apenas por 1 GoRoutine por vez*/

# Imagem 02
/**Usaremos MUTEX em caso de REUNITLIZAÇÃO de valores entre GoRoutine*/
/**Sync.MUTEX Disponibiliza acesso exclusivo p recursos compartilhados, ou seja mais de uma variavel pode LER ao mesmo tempo. */

/**Seção Critica pode ser usar de 2 formas, uma delas é com defer */
## Sintax 
mu.lock()
localVariavel += amount
mu.Unlock()
# ou Neste ex, o defer será o ultimo a execultar no seu codigo
mu.Lock()
defer mu.Unlock()
localVariavel -= amount
# Imagem 03
/**Sync.RMutexm permite q várias GoRoutine LEIAM ao mesmo tempo, mas somente 1 vai GRAVAR*/
## Sintax 
/**Cria o lock, lê o valor(return) e libera o lock do Mutex*/
/**O Ideal seria criar um GetBalance e comportilhar entre as GoRoutines*/
mu.RLock()
defer mu.RUnlock()
return localBalance 





#### Aula04 Channels 
/**Channels é a forma q temos de compartilhar dados entre as GoRoutines*/
/**Channels  é usado para organizar a execução de GoRoutines*/
/**No Exemplo de GoRoutines e Ponteiros tinhamos a situação q assim que mudarmos o valor do Ponteiro era atualizada a GoRoutines, o problema desta abordagem é que se o dado vem de fora da aplicação Ex: uma API teremos problema de time, ou seja a GoRutines passa a ficar bloqueante*/

/**Em Channels podemos fazer q um GoRoutine ESPERE o valor enquando não foi abastecido, o que acontece é que a GoRoutine q envia, manda o valor para Channel e com isto não precisa esperar e ficar bloqueada*/

/**Na Imagem 01C temos 1 channel de string com buffered 1, a variavel é instanciada com a função MAKE() e enviaremos uma mensagem como Exemplo de "ping". Em outra lugar da aplicação em outra GoRoutine poderemos receber esta mensagem. Ou seja tenho 2 GoRoutines e no meio um Canal() comunicando as 2s*/
#  OBS: Muito parececido com Pipe do RXJS
# Ou Seja para eu passar um valor de uma GoRoutine para outra e o valor ficar sempre atualizado mandamos o valor pelo Canal q no final deste Canal teremos um outra GoRoutine esperando algo. Então temos Entrada de Dados e Saida de Dados.

# Channels Directions imagem 02  Canal <-"Dados"  ou chTest := <- Canal
## Quando a SETA esta a DIREITA o Channel é ABASTECIDO ou ENTRA no Channel
## Quando a SETA esta a ESQUERDA o Channel é o ABASTECEDOR ou SAI do Channel p um variavel

# Obs:Podemos ter Channels normais ou especificar o tipo
/**com o operador "<-" podemos determinar se é entrada(Envio) de dados ou saida (Recebimento) QUANDO usamos como PARAMETRO de FUNÇÂO */
# OBS: O Channels tem as funções READ e WRITE
/**Desta forma podemos determinar q o paramentro SO ENVIA ou SÓ RECEBE, imagine q queremos boquear q algume coloque valor mas somente receba!*/

# Seta depois do Channel <- "dados" quer dizer q so pode INSERIR ou ATRIBUIR dados (CHANNEL ENVIAR o VALOR)

# Seta antes do localVariavel <-Channel quer dizer q so pode RECEBER dados (CHANNEL RECEBER o VALOR) ou seja a variavel RECEBE o dado de CHANNEL
/**Imagine q queira criar o Package ou um biblioteca e queira compartilhar onde quem for usar só pode RECEBER ou seja LER nunca Gravar*/

# UnBuffered Channels são canais criados sem por o tamanho
# Obs: Imagem 03 
/**Não existe buffer entre quem envia e quem recebe*/
/**É usado para dados SÍNCRONOS, ou seja a GoRoutine Remetente fica bloqueda ate q a GoRoutine Destinatário recebe os dados.**/
## OBS: o Buffer é a capacidade de informação q é comportada dentro do Channel
/**Ex: Se colocar o Buffer de 3, que dizer que o limite de WRITE da GoRoutine será no máximo 3 GoRoutines escrevendo concorrentemente . Ou 3 valores na FILA de execução deste Channel*/

## OBS: Um canal pode ser Síncrono ou Assíncrono, caso for UnBufferd ele se torna Síncrono
/**Um Channel UnBuffered só a eceita que dados sejam inseridos caso HAJA um consumidor. Ex: Caso coloque uma string em um Channel, este dados ficará bloqueado ate que tenha um GoRoutine para receber ou consumir*/

# Buffered Channels são Assícronos ou seja existe um Buffer entre a Entrada e Saida do Canal
# Imagem 04
## OBS um Buffered Channel é Assíncrono enquanto haja Buffer, caso todos o Buffed fiquem cheios ela passa a ser UnBuffered.

/**Neste caso poder DEFENIR o tamanho do Buffer**/
/**No Buffered Channel é a indicação do numero de elementos q podem ser enviados sem que a GoRoutine DESTINATARIO esteja pronta para receber valores*/
/**Neste caso não precisamos de um consumidor, isto ficará no Buffer do Channel**/
/**Isto ficará em um LISTA chamada SendQuee e ReceverQuee**/

# Channels são Bloqueantes imagem 05
/**O Channel espera uma GoRoutine que seja Remetente e um Destinatário estar pronto sem isto não há fluxo de dados */
/**O Channel é o responsavel de avisar a GoRoutine se por ou não novos dados ou seja Write e receber um novo dado*/

## OBS Cuidados com Channels imagem 06
/**nil é o valor padrão na criação de um Channel*/
/**Caso criarmos um Channels assim: var localChannel chan interface{} ele será nulo*/
/**Se Criamos com a Função Make() para evitar isto*/

# OBS: Se tentarmos Ler ou escrever em um Channel nulo, o codigo fica BLOQUEADO.

## OBS Cuidados com Channels imagem 06B
/**Fechar um Channel "close()" vai quebrar o codigo ao receber um PANIC*/
/**Sempre tenha CERTEZA que um Channel foi inicializado!!!**/

# Função Close(ch)
/**será usada quando não tem mais valores para serem enviados*/

# Value, ok := <-  Imagem 08, podemos pegar o valor de 2 formas, com o sem a variavel OK
/***Neste caso temos o Valor ou seja os dados e o OK*/
/**Se dentro da variavel OK tivermos TRUE quer dizer: Ainda temos dados a receber*/
/**Se dentro da variavel OK tivermos FALSE quer dizer: A função Close() fechou o Canal*/

# OBS:  Podemos ter multiplos retorno no GO, por isto podemos ter Valor, OK*/

# Range Channels
/**É a forma q temos de dar um Loop no Channel, Itera os valor recebidos, o look encerra automaticamente e NÃO permite usar multiplo retorno, ou seja a variavel OK*/

/**O Range não tem a variavel OK, pois só aceita valores diferente de NIL e com isto não faz sentido ter o OK, visto que nunca seria usado */


#### Aula05 Wait Groups
https://github.com/HunCoding/golang-asynchronous/tree/main/wait-groups

# Entendendo o uso de WAIT GROUPS em GO - GoLang Assíncrono #go #async #goroutines
/**WaitGroup é uma função do package async que trabalha junto com as GoRoutines**/
/**A Grande utilidade é fazer aplicação rodar sem bloquear as GoRoutines e sem ficarem sem resposta, pois o codigo parou a execução, algo q ja vimos este comportamento*/

/**Vamos supor q que temos um Backend e vamos fazer 5 chamadas para salvar no Banco de Dados, e cada vez que  chamamos o banco isto demora 2Segundos então temos 10Segundos, algo ruim p uma aplicação*/
/**A solução seria concorrencia usando GoRoutines, mas como ja vimos se colocarmos as 5 GoRoutines p funcionar ao mesmo tempo, veremos q o codigo terminará antes da execução das GoRoutines, Ja vimos este comportamento, pois depois do fim do programa as GoRoutines morrem sem serem execudadas*/

/**Ou Seja tinhamos 1 problema de demorar 10 Segundos e quando colocamos GoRoutines no Lugar e as mesmas morrem sem execução, criamos outro problema*/
# 1º Exemplo faremos sem GoRoutines 
fazendo linha por linha

# Exemplo02 faremos com GoRoutine mas SEM WaitGROUP
/**Onde veremos o comportamento que o metodo main() terminará antes q as GoRoutines exexutem!*/

# Exemplo03 faremos com GoRoutines e WaitGROUP
/**A função WaitGroup falará com a função Main(), olha temos 5 GoRoutines e preciso q espere...*/
/**Algo assim só fecha o Main() quando a ultima GoRoutine falar que pode! É algo assim*/

/**OBS: Caso pessemos um numero ha menos no Add() ex: Add(3), o Wait() so vai esperar 3 GoRoutines**/
/**OBS: Caso passemos um numero há mais no Add() Ex: 10, teremos um erro fatal e um deadLock**/

#### Aula06 Select em GoRoutines
https://github.com/HunCoding/golang-asynchronous/tree/main/select
/**O Select em GoRoutines é bem parecido com SwitchCase*/
/**O SELECT é um keyword de decisões usando o Select com GoRoutines**/
/**Imagine q temos 3 chamadas em GoRoutine ao mesmo tempo e quem responder 1º utilizaremos, colocaaremos as 3 chamadas em Channels e quem receber 1º utiliza e cancela s outras*/

## exemplo01 
/**Entendendo o funcionamento do SELECT*/
/**Ex: Imagine q dentro do Select temos uma GoRoutine que irá receber  time.out de 2S ou outra GouRoutine que receberá um valor de uma API, caso a API demore mais de 2S, o time.Out setará o valor na GoRoutine e fechará o Select e com isto podemos proteger o codigo em situações de demora. */


## exemplo02  com Select  For eterno e fatal.erro
/**Neste exemplo temos o retorno não só do 1º, mas de todos a medida q os canais recebem valores*/ 
/**O Problema nas abordagem que o FOR gerará um error, pois continuará a esperar outras GoRoutines*/

## exemplo03  com Select  For eterno e fazer a comunição Não bloqueante
/**Para resolver a questão bloqueio, podemos criar context.done()*/
/**E Colocar um Context, pois todo Context é um canal dentro de um Case*/


#### Aula07 Atomic em GoLang
https://github.com/HunCoding/golang-asynchronous/tree/main/atomic
# Imagem 01A
/**ATOMICS é uma forma de sincronisar recursos em GoRoutines*/
/**Porque é muito usado em aplicações Assícronas em numeros Inteiros normamente   em Contadores*/
/**Resumidamente o ATOMIC é uma operação q não pode ser feita pela METADE, ou se faz ou não faz! Isto n tem a ver com a operação , e sim com o estado, alterado ou mantido? Ou seja o estado de uma variavel é ATOMICO somente quando é alterado ou não, não meio alterado*/

# Imagem 02
/**Exemplos de Atomico: Ligar a luz! */
/**Exemplos de NÃO Atomico: encher um copo de agua! Mas o copo esta pela metade? Transbordando ? ou Cheio? esta vendo a diferença? */
/**Ou seja ou uma tarefa será feita ou não feita*/

# Imagem 03
/**UINT que dizer valores sem SINAL, e so podem ser valores positivos*/
/**Temos variaões Uint8, 16, 32 e 64*/
/**São Muito usado para IDs numericos*/

# Imagem 04
/**O Package Atomic em GOLang:*/
/**Em Go, o package Atomic fornece formas primiticas de baixo nível*/
/** Usaremos quando precisarmos comprtilhar uma variavel entre GoRoutines*/

# Imagem 05
/**Operações atomicas de baixo nivel em memoria para GoRoutines*/
/**Operações não bloqueantes para GoRoutines*/
/**Contadores Atomicos para GoRoutines*/

## Sintaxe onde passamos o Ponteiro da Variavel
# atomic.AddUint64(&ops,1)
# value := atomic.LoadUint65($ops)

# Exemplo SemAtomic 
/**temos as saidas variadas*/
// Counter: 50000
	// Counter: 50000
	// Counter: 49442
	// Counter: 49994
# Exemplo ComAtomic

/**O Problema é corrigo, so temos  Counter: 50000*/

#### Aula08 Sync Once
https://github.com/HunCoding/golang-asynchronous/tree/main/once

/** Sincronizando execução de funções em GoLang*/
/**Sinc Once é uma função que garante q uma função seja executada somente uma VEZ*/
# Imagem 01B
/**Tudo que tiver dentro do escopo once.Do() somente é executado 1vez*/

# Imagem 01C
/**Utilize ONCE caso queira IDEMPOTÊNCIA, mesmo q tenhamos vais GoROutines Usando a mesma FUNÇÃO, o ONCE permitirar que somente seja executada 1VEZ*/

# Imagem 02 Onde usar ?
/**Imagine q tenhamos uma função. Ex: Init da sua aplicação q é usada por varias GoRoutines, para que sua aplicação seja iniciada somente 1 VEZ*/
/**Mesmo que varias GoRoutines use esta função, somente a 1ª que chamar terá permissão de execução*/
## OBS: É Marcado uma Flag na função q bloqueia a execução , por baixo dos panos é usado o MUTEX, mas aqui de forma facilitada 

# Exemplo 01 Sem Once
/**O Codigo será executado 20x*/
# Imagem 03
Executando o Codigo do Init....
Executando o Codigo do Init....
Executando o Codigo do Init....


# Exemplo 02 Com Mutex no lugar do Once
# Imagem 04
/**Aqui para fazer funcionar, temos que criar uma Flag, um If etc, é so um exemplo como o ONCE faz por debaixo dos panos*/

# Exemplo 03 Com Mutex no lugar do Once
/**A forma mais FÁCIL é usar o ONCE!!*/

# Imagem 05
/**Dentro da classe Once, veremos algums funções*/
# done que é uint32 que será a nossa Flag
# m que é Mutex

# Imagem 05B e 05C
/**A função Do() recebe como paramentro uma Função*/


#### Aula09 Sync.Cond()
## YNC.COND - SINCRONIZANDO EXECUÇÃO DE GOROUTINES | Golang Assincrono #go #sync #conditinal
https://www.youtube.com/watch?v=M9hQdtJfkfE&list=PLm-xZWCprwYQz7Di62Z93zGHyLC-HkFpb&index=10

https://github.com/HunCoding/golang-asynchronous/tree/main/signal
# Imagem 01B

/**O Cond é um Sinal q ajuda a execução de algumas GoRoutines ou não*/
/*Hoje aprendemos um assunto bem legal a respeito de sincronizar execução de goroutines nas quais precisamos fazer com que compartihe valores e condições**/
/**E desta forma podemos sincronizar uma ou varias GoRoutines*/
/**Ex: Como se tivessemos um conteiner de >GoRoutines q ficam esperando um condição ser verdadeira para serem executadas.*/

/**Imagine um GoRoutine q só pode executar aparti do momento q uma codição for satisfeita. Ex: 1 GoRoutine chama uma API  ou MicroService e a Outra GoRoutines pega este valores e salva em BD*/
/**Neste caso teriamos q fazer com que a GoRoutine q vai salvar no DB fique esperando a GoRoutine que vai na API buscar os dados*/
/**Temos varias forma de fazer, desde usar um Select ou IF*, mas a forma mais otimizada é usar o Sync.Cond()/

# Imagem 02 temos um exemplo usando o Mutex, muito verboso.
/**Poderiamos usar o Mutex ou Channels*/
/**Uma das formas de fazer um For Infinito e comparar tamanho enquanto ele for vazio, depois dar um Lock() somente para 1 GoRoutine Gravar */
/**Ou seja acada 100 miliS ficaremos checando dentro do FOR se alguma GoRoutine gravaou algo*/

# Image 02B
/**Criamos assim um problema performatico, pois estamos dando um look infinito, acada 100 Ms tem que checar um Map etc*/

# Imagem 02C
/**Precisamos suspender um execução de uma GoRoutine ate q ela recebe um Signal ou uma condição seja especifica*/
/**Com isto precisamos de alguma forma avisar a GoRoutine suspendida já é hora de voltar a executar*/

# Imagem 02D
/**O CHANNEL ajudaria , pois só vai gravar a entrada se houver saida!*/
/**Ja que um CHANNEL pode ter Buffed, posto ficar esperando varios valores*/

# Imagem 02E
/**Se tivermos muitos CHANNEL e muitas GoRoutines? Fica inviavel dentro do codigo*/

# Imagem 03 
/**Criaremos uma Var que recebe um valor de um Ponteiro Sync:Cond**/
/**Utilizamos o construtor sync.NewCond() q recebe uma inteface Sync.Locker()
*/
/**Criamos uma var q recebe a interface do Mutex*/
# m := sync.Mutex{}
/**Criamos o construtor newCond() q recebe a Ref. da  interface*/
# c:= sync.NewCond(&m)

# Imagem 04 e 04B
/**Criada a Variavel c:= sync.NewCond(&m) temos agora 3 Metodos*/
/**1º e o mais Importante*/
# C.Wait() 
/**Suspende a execução da GoRoutine na qual foi executada e libera o Lock() do Mutex automaticamente, por isto que recebe como paramentro do Mutex*/
/**Não preciso me preocupar quando ficará bloqueado pelo Lock()*/
/**Só Retorna o valor somente quando for "Acordado usando as funções BroadCast ou Signal"*/
/**A GoRoutine bloqueada, retorna para Loop para validar novamente o valor logo depois do fim do Wait()*/

# C.Signal()
https://github.com/HunCoding/golang-asynchronous/tree/main/signal
## OBS: Aqui no Signal a 1º GoRoutine q chamou o c.Wait() será a 1º q será notificada pelo Signal
# Imagem 04C
/*Acorda uma GoRoutines q espera em C**/
/*Signal procura a GoRoutine q tiver esperando pro mais tempo e notifica/avisa ela 1º, o Signal vai usar FiFo**/
/*Signal é ideal para poucas GoRoutines, quando for muitas é melhor o Broadcast**/
/*1º Saida **/
## OBS: Os tempo de atribuição são Randomicos
Valores:  [         ]
Valores:  [      g   ]
Valores:  [  c    g   ]
Valores:  [  c d   g   ]
Valores:  [  c d   g h  ]
Valores:  [ b c d   g h  ]
Valores:  [a b c d   g h  ]
Valores:  [a b c d   g h  j]
Valores:  [a b c d   g h i j]
Valores:  [a b c d e  g h i j]
Valores:  [a b c d e f g h i j]

/*2º Saida , será diferente da 1º  **/
## OBS: Os tempo de atribuição são Randomicos
Valores:  [         ]
Valores:  [   d      ]
Valores:  [ b  d      ]
Valores:  [ b  d   g   ]
Valores:  [ b  d   g   j]
Valores:  [a b  d   g   j]
Valores:  [a b  d   g  i j]
Valores:  [a b  d  f g  i j]
Valores:  [a b  d e f g  i j]
Valores:  [a b  d e f g h i j]
Valores:  [a b c d e f g h i j]
/**No SIGNAL() para cada letra inseriada, ele PRINTA e CHECA se tem espaço e manda um SINAL*/

# Imagem 05 com Signal
/**Exemplos de codigo*/





# C.BroadCast()
## OBS: Aqui no BroadCast caso haja 100 GoRoutine com c.Wait() serão as 100º q serão notificada pelo Signal ao mesmo tempo.
https://github.com/HunCoding/golang-asynchronous/tree/main/broadcast
# Imagem 04D
/**A diferença do Signal() é que em Broadcast(), todas as GoRoutines utilizam a var C e executaram o C.Wait() */
/**É permitido mas não necessario da GoRoutine que chamar o BroadCast() ficar com o lock ativo*/
/**Imagine q tenha 100 GoRoutines. O Signal vai 1 por uma procurando q está em lock(), mas BroadCast acorda a todas */
## OBS: o UnLock() é automatico, pois ja passamos a interface por Paramentro.


# Imagem 06 com BroadCast