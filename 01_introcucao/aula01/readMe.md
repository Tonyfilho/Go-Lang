Aula01

# criando uma projeto em GO

\*\* go run + nome do projeto
rodando o projeto go run aula01.go


# 01 Operador curto de declaração, populamente chamada de Marmota

Obs: é igual ao const e let no type,
tem que esta dentro de um bloco { := }

localNome .= "Tony"

OBS: vem com typagem automatica, não pode ter valor nulo ou undefine, tem que ter valor,
isto faz parte da linguagem GO

// vendo a typagem usando o operador %T
fmt.Printf("x is of type %T\n", x)
fmt.Printf("y is of type %T\n", y)

// saida
x is of type int
y is of type string

# 02 declaração de variavel Global com a palavra VAR
Para uso de variavel Gobal temos que usar a palavra Var
Variaveis globais pode startar como nulas, mas so podem receber valores depois dentro de um bloco {..}

# 03 Opertadores e Operando e Expressões, Statement
Opertadores e Operando, Iqual as outras liguagens.

Expressões são tudo que produz uma resultado, ex: var local = 10 + 10, o resultado será  a expressão.

Statement, são linha de codigos, ou chamada de funções ex: fmt.Printf("localVariable is of type %T\n", localVariable)

OBS: Um statemente é formado de 1 ou mais expressoes, ou seja conjunto

# chaves (), Parenteses {} e cochetes []

# Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
slice, array, struct, map

# Atribuição de valores a variaveis

Valor zero, é nome do valor inicial de uma variavel antes de ser atribuido valor, não estou
dizendo que uma variavel do tipo string vari ter um ZERO é o nome deste momento.
-int: 0
-float: 0.0
-boleans: false
-string: ""
-poiters, functions, interfaces, slicesm channels e maps: nil

inicialização é o 1º valor q coloco em uma variavel

