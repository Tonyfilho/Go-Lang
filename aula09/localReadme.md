#### Funções

# 01  Funções  https://go.dev/ref/spec#Function_types
Qual a utilidade de funções?
Abstrair funcionalidade
Reutilização de código
# Temos q ter esta extrutura, func (receiver) identifier(parameters) (returns) { code }
# É algo assim: func (receiver) NomeDaFução(QualquerCoisaOuTipoDeDado) (Tipo de dado Retornado Ex: String, ou uma soma, ou uma outra Função, lembrando q podemos ter mais de um retorno ) { É o que vc vai programar}
A diferença entre parâmetros e argumentos:
Funções são definidas com parâmetros
Funções são chamadas com argumentos
Parâmetro pode ser ...variádico tem que ser o ultimo paramentro
Pass by reference, pass by copy, ... não.
# Obs: Tudo em Go é pass by value. Ou seja é o temos dentro da Variavel e não a Referencia de memoria criada
Exemplos:
Função básica. 
Go Playground: https://play.golang.org/p/FebJblBenP
Função que aceita um argumento. 
Go Playground: 
        https://play.golang.org/p/CE6Ij3U4QB
Função com retorno. 
Go Playground: https://play.golang.org/p/gKxwYe6btP
Função com múltiplos retornos e parâmetro variádico.
Go Playground: https://play.golang.org/p/OcQ1wXwM2c
Mais um: https://play.golang.org/p/8wc2TA9xH_


# 02 Funções Desenrolando (enumerando) uma slice

Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
Exemplos:
Desenrolando uma slice de ints com como argumento para a função "soma" anterior
Go Playground: https://play.golang.org/p/k8O3__8UDa
Pode-se passar zero ou mais valores
Go Playground: https://play.golang.org/p/C238I9n7Vs
O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
Go Playground: https://play.golang.org/p/8wc2TA9xH_
Não roda: https://play.golang.org/p/2qTAnLWfgB


# 03 Funções Defer siguinifica Postergar ou Adiar
Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
# Obs: Uma declaração DEFER chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.
Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
"Deixa pra última hora!", cria assincronismo
ref/spec
Sempre usamos para fechar um arquivo após abri-lo.

Ex: Abrir e fecha conexão de rede.
Go Playground: https://play.golang.org/p/sFj8arw0E_