#### Aula19 Concorrencia Vs Paralelismo

# Concorrência é quando abre uma padaria do lado da outra e as duas quebram :)
Fun facts: 
# 01 O primeiro CPU dual core "popular" veio em 2006
Em 2007 o Google começou a criar a linguagem Go para utilizar essa vantagem
# Go foi a primeira linguagem criada com multi-cores em mente C, C++, C#, Java, JavaScript, Python, etc., foram todas criadas antes de 2006
Ou seja, Go tem uma abordagem única (fácil!) para este tópico
# E qual a diferença entre concorrência e paralelismo?

## Concorrencia com Goroutines & WaitGroups

# O código abaixo é linear. Como fazer as duas funções rodarem concorrentemente?
https://play.golang.org/p/XP-ZMeHUk4
Goroutines!
O que são goroutines? São "threads."
O que são threads? [WP](https://pt.wikipedia.org/wiki/Thread_...)
Na prática: go func.
Exemplo: código termina antes da go func executar.
Ou seja, precisamos de uma maneira pra "sincronizar" isso.
Ah, mas então... não.
Qualé então? sync.WaitGroup:
Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
func Add: "Quantas goroutines?"
func Done: "Deu!"
func Wait: "Espera todo mundo terminar."
Ah, mas então... sim!
Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()

Go Playground: https://play.golang.org/p/8iiqLX4sWc

# hread (em português: fio de execução[1] ou encadeamento de execução) é uma forma como um processo/tarefa de um programa de computador é divido em duas ou mais tarefas que podem ser executadas concorrentemente ("simultâneo"). 

## aparentemente vejo algo parecido com a PROMISES, pois precisamos esperar o fim da execução