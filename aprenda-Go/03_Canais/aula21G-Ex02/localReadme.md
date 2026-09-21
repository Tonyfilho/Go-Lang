#### Aula21G Exemplo02

https://www.youtube.com/watch?v=VJyryKEMleU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=155
# Exemplo02. Rob Pike (palestra Go Concurrency Patterns):
Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.


/*
novo <- x
//     ↑
//     x é enviado para novo
//     Mas x é um canal, não uma string!

novo <- <-x
//     ↑  ↑
//     |  └── Recebe um valor de x (string)
//     └── Envia esse valor para novo
**/

/*
// Opção 1: Separado (mais legível)
v := <-x
novo <- v

// Opção 2: Junto (conciso mas confuso)
novo <- <-x

// Opção 3: Com range (melhor para múltiplos valores)
for v := range x {
    novo <- v
}
**/