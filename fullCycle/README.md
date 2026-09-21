#### Curso de Go no canal Full Cycle

#### FullCicles GoRoutines - Entendendo Go Routines minuto 45 do video 

## Aula 19 Especial - Entendendo Go Routines minuto 45
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