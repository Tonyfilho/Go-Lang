#### Aula23 Tratamento de Erros – 1. Entendendo erros

Para quem já programa em outras linguagens:
Em Go não temos exceções. → https://golang.org/doc/faq#exceptions
"We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code."
"Go's multi-value returns make it easy to report an error without overloading the return value. A canonical error type, coupled with Go's other features, makes error handling pleasant but quite different from that in other languages."
Aventureiros: https://blog.golang.org/error-handlin...
É interessante criar o hábito de lidar com erros imediatamente, similar a e.g. defer close.
package builtin, type error interface
package errors



https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/23_tratamento-de-erros/02_verificando-erros


# 01 Exemplo fmt.Println() → stdout
# 02 Exemplo log.Println() → timestamp + pode-se determinar onde o erro ficará logado
# 03 log.Fatalln() → os.Exit(1) sem defer
# 04 log.Panicln() → println + panic → funcões em defer rodam; dá pra usar recover
panic()
Recomendação: use log.
Código: 
1. fmt.Println
2. log.Println
3. log.SetOutput
4. log.Fatalln
5. log.Panicln

# RESULMO. O fatalln é como desliga se a tomada, panic ainda roda as defer antes e vai encerrando por partes as funções. Mas podemos usar a função recover() e reasulme o controle.

O Pacic seria.. deu Merlin , mas vamos resolver, diferentemente de os.Exit()
log.Println() é um erro simples
logFatal() é para erros muito graves
logPanic() é para erros medios mas sem o print na tela
logPanicln() é o mesmo do panic mas com o pront, é o mais recomendado