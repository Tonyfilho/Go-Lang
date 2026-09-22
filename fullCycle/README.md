#### Aula02 Sobre Packages e importações como funciona

# 📁 Estrutura que vamos criar
pacotes/
├── go.mod                    ← módulo: pacotes
├── main.go                   ← package main
└── math/                     ← subpacote
    └── math.go               ← package math


# 2. Inicialize o módulo (cria o go.mod)
Dentro da pasta pacotes, execute:

bash
go mod init pacotes
Isso vai criar o arquivo go.mod com este conteúdo:

# 💡 O nome pacotes é o nome do módulo. Você pode usar qualquer nome, mas evite espaços e caracteres especiais. O ideal seria algo como github.com/seu-usuario/pacotes, mas para estudo local, pacotes já serve.

Dentro de pacotes/math/, crie o arquivo math.go:

# ⚠️ Atenção: O nome do pacote (package math) é o que você vai usar no import do main.go. Como o pacote se chama math, o caminho do import será pacotes/math


Podemos ter mais de uma arquivo no mesmo Packge, o Go vai acha-lo

## Obs: Podemos ter acesso, a Variaveis, Funções dentro dos packages

# Lembrando que inicio de nome de variaveis , funções e metodos  Upcase = public e DownCase = privite


# Documentação e uso do GoDoc
 Documentação em Go Ex:  da função criada, tem que começar com nome da Função iqualmente a declaração.
# Isto a biblioteca GoDoc reconhece este padrão e faz a documentação para nos
Ex:
// Soma faz a soma de 2 numeros e retorna um int
func Soma(a, b int) int {

	return a + b
}




#### Curso de Go no canal Full Cycle

#### FullCicles GoRoutines - Entendendo Go Routines minuto 45 do video 

##  - Entendendo Go Routines minuto 45
https://www.youtube.com/watch?v=B4NL0rMvXMg

## Aqui entenderemos que: Enquando não for usando o valor que está em uma GoRoutines, não será atribuido um novo valor.

# Se ninguem usar esta Variavel QUEUE em outra GoRoutine, a tribuição fica Bloqueada

📤 Tentando enviar: 0
📥 Canal Usado e Lido:  0
✅ Enviado com sucesso: 0
📤 Tentando enviar: 1
📥 Canal Usado e Lido:  1
✅ Enviado com sucesso: 1
📤 Tentando enviar: 2
📥 Canal Usado e Lido:  2
✅ Enviado com sucesso: 2
...
📥 Canal Usado e Lido:  19
✅ Enviado com sucesso: 19
📤 Tentando enviar: 20
📥 Canal Usado e Lido:  20
🛑 Limite atingido, saindo...
✅ Programa finalizado!

## Note como 📤 Tentando enviar: N sempre vem antes de 📥 Canal Usado e Lido: N — isso prova que:

A goroutine bloqueia no queue <- i

Só desbloqueia quando o consumidor lê

Só então o i++ acontece e o próximo envio é tentado

## SEM CONSUMIDOR:
Produtor: queue <- 0  →  🔒 BLOQUEADO PARA SEMPRE
                         (o programa congela aqui)

## COM CONSUMIDOR:
Produtor:  queue <- 0  →  🔒 bloqueia
Consumidor: <-queue    →  ✅ recebe 0
Produtor:  (desbloqueia) → i++ → queue <- 1 → 🔒 bloqueia
Consumidor: <-queue    →  ✅ recebe 1

##  Tabela de Erros e Correções
Erro no seu código	Problema	Correção
close(queue) após iniciar goroutine	Fecha antes de enviar valores	Não feche se o produtor é infinito
<-queue na mesma goroutine que envia	Não demonstra concorrência	Remova, deixe só o consumidor ler
close(queue) antes de lerGoRoutines	Range termina imediatamente	Feche APÓS o envio, ou não feche
Falta de time.Sleep ou controle	Loop muito rápido para ver	Adicione pausa para visualizar


#### Aula19 Ex02 Explicação de GoRoutine e o Uso do Close()

# neste exemplo estamos compartinhando variavel QUEUE entre o escopo de MAIN{} e o escopo da Go Routine Função expressa {}()


### OBS: Sobre o close()

go
# // ✅ CORRETO: Fechar quando o produtor TERMINAR
go func() {
    defer close(queue)  // Fecha quando a goroutine terminar
    for i := 0; i < 10; i++ {
        queue <- i
    }
    // Aqui a goroutine termina e o close() é chamado
}()

# // ✅ Consumidor usa range (termina quando o canal fechar)
for x := range queue {
    fmt.Println(x)
}

# // ❌ ERRADO: Loop infinito + close
go func() {
    for {  // Loop infinito
        queue <- i
        i++
    }
}()
# close(queue)  // ❌ Nunca deveria fechar (produtor ainda está ativo)

#### Aula Video04 Essa é a técnica para tratar erros em Golang Exemplo 01

https://www.youtube.com/watch?v=f--pS45o_zg&list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg&index=4

# O erro no GoLang é tradado de forma diferente, sem Try and Catch, o erro m Go tem as mesma prioridade do que os dados.

#  No Go o Erro é explicitos na maioria das funções
em Go podemos retornar mais de 1 valor em um função, com isto podemos retornar o valor e o error caso exista.

# O grande diferencial é que podemos escolher o que faremos com erro.
Ex: Posso dar Panic, Logfatal, ou chamar um outro recurso etc

#### Aula Video04 Essa é a técnica para tratar erros em Golang Exemplo 02

# Criando seu proprio error

Criaremos um função simples onde teremos o retorno de um dado e do erro caso haja

Caso não queira usar a variavel de erro, poderá usar Blank Indentifier "_"
ou seja o underscore no lugar da var err

# Ex res, _ := soma(1, 10) com isto não temos a continuação do error, mas como saberá

è uma variavel que exclui o dado para não ser utilizado

