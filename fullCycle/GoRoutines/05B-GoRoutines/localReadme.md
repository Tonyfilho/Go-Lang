#### #### FullCicles GoRoutines - Entendendo Go Routines minuto 45
# GoRoutine e o Uso do Close()

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