#### Go-Lang
Curso de Go
## Guias de pesquisa 
# golang.org/ref/spec#For_statements
# Effective Go


#### aula 01
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



#### aula 02

# 01 estudaremos sobre a função FMT que é de imprimir

https://pkg.go.dev/fmt

temos 3 principais categorias
01 Grupo #1: Print → standard out
func Print(a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
Format verbs. (%v %T)
02 Grupo #2: Print → string, pode ser usado como variável
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
02 Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

# 02 temos 2 tipo de strings em Go: Interpreted literal e Row Literal

# Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte.

Interpreted Literal: são valores fixos, int , boolean ; Usamos \N e cria uma linha nova

Ex: name := "Go Developers"
fmt.Println("Tony \n" + name)

Ou seja quando a função vai ler a lina ela interpreta que o \n é para criar uma nova linha ou \t que será interpretado com um TAB

Row Literal (Ou seja CRU)
será feito o que está escrito `...`

fmt.Println(`Tony \n` + name)

Obs: cada caracter no Go é considerado um Rune Literal

# 03 Criação do proprio tipo, usaremos palavra reservada TYPE, é parecido com TypeScript

Revisando: tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
Tem uma história que Bill Kennedy dizia que se um dia fizesse uma tattoo, ela diria "type is life."
Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
type hotdog int → var b hotdog (main hotdog)
Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog.

OBS: em GO tipo são imutaveis, ou seja se inicio uma variavel como string ela NÃO MUDA

Ex:
type Hotdog int
var localTipo Hotdog
localTipo = 5
fmt.Printf("%v %T\n", localTipo, localTipo)

# 04 conversão de Types, é o mesmo que CASTING de outras liguagens
Conversão de tipos é o que soa.
Em Go não se diz casting, se diz conversion.
a = int(b)
ref/spec#Conversions

OBS eIqual := localInt == localTipo  Erro no compilador, não permite retornar um boolean em GO


#### Aula03 



# 01 como funcionam os computadores

Isso é importante pois daqui pra frente vamos falar de ints, bytes, e etc.
Não é necessário um conhecimento a fundo mas é importante ter uma idéia de como as coisas funcionam por trás dos panos.
ASCII: https://en.wikipedia.org/wiki/ASCII
Filme: Alan Turing, The Immitation Game.

# 02 Falaremos sobre boolean
Agora vamos explorar os tipos de maneira mais detalhada. golang.org/ref/spec. A começar pelo bool.
O tipo bool é um tipo binário, que só pode conter um dos dois valores: true e false. (Verdadeiro ou falso, sim ou não, zero ou um, etc.)
Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.
Na prática:
Zero value
Atribuindo um valor
Bool como resultado de operadores relacionais
Go Playground: https://play.golang.org/p/7joj615nZw

OBS: são uados operadores relacionais == ; <= ; >= ;  < ;  > ;
Sempre que você ver operadores relacionais, o resultado da expressão será um valor booleano.


# 03 Inteiro vs Frações
int vs. float: Números inteiros vs. números com frações.
golang.org/ref/spec → numeric types

*****Integers:
Números inteiros
int & uint → “implementation-specific sizes”
Todos os tipos numéricos são distintos, exceto:
byte = uint8
rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
Tipos são únicos
Go é uma linguagem estática
int e int32 não são a mesma coisa OBS: int fica no automatico, caso o processador precise ele muda para 32 ou 64.
Para "misturá-los" é necessário conversão
Regra geral: use somente int

******Floating point:
Números racionais ou reais
Regra geral: use somente float64, é automatico, caso somente coloque float, ele vai usar o 64
Na prática:
Defaults com :=
Tipagem com var
Dá pra colocar número com vírgula em tipo int?
Overflow
Go Playground: https://play.golang.org/p/dt2x1ies5b
“implementation-specific sizes”? Runtime package. Word.
GOOS
GORUNTIME
https://play.golang.org/p/1vp5DImIMM

# 04 Overflow
Um uint16, por exemplo, vai de 0 a 65535.
Que acontece se a gente tentar usar 65536?
Ou se a gente estiver em 65535 e tentar adicionar mais 1?
Playground: https://play.golang.org/p/t7Z4m127F2t




#### aula 04

# 01 Strings são sequencias de bytes.

String são Imutáveis.
Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
Na prática:
%v %T
Raw string literals
Conversão para slice of bytes: []byte(x)
%#U, %#x
Go Playground: https://play.golang.org/p/dt2x1ies5b & https://play.golang.org/p/PpDnspiyA_7
https://blog.golang.org/strings

OBS: cada item é um byte


# 02 Constantes

São valores imutáveis.
Podem ser tipadas ou não:
const oi = "Bom dia"
const oi string = "Bom dia"
As não tipadas só terão um tipo atribuido a elas quando forem usadas.
Ex. qual o tipo de 42? int? uint? float64?
Ou seja, é uma flexibilidade conveniente.
Na prática: int, float, string.
const x = y
const ( x = y )

# 03 IOTA

golang.org/ref/spec
Numa declaração de constantes, o identificador iota representa números sequenciais.
Na prática.
iota, iota + 1, a = iota b c, reinicia em cada const, _
Go Playground: https://play.golang.org/p/eSrwoQjuYR

# 04 Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
https://play.golang.org/p/7MOnbhx4R4
  / bit-hacking-with-go  
Fim da sessão. Massa!


#### aula 05

## 01 Fluxos de Controle

####  Fluxos de Controle


Computadores lêem programas de uma certa maneira, do mesmo jeito que nós lemos livros, por exemplo, de uma certa maneira.
Quando nós ocidentais lemos livros, lemos da frente pra trás, da esquerda pra direito, de cima pra baixo.
Computadores lêem de cima pra baixo.
Ou seja, sua leitura é sequencial. Isso chama-se fluxo de controle sequencial.
Alem do fluxo de controle sequencial, há duas declarações que podem afetar como o computador lê o código:
Uma delas é o fluxo de controle de repetição (loop). Nesse caso, o computador vai repetir a leitura de um mesmo código de uma maneira específica. O fluxo de controle de repetição tambem é conhecido como fluxo de controle iterativo.
E o outro é o fluxo de controle condicional, ou fluxo de controle de seleção. Nesse caso o computador encontra uma condição e, através de uma declaração if ou switch, toma um curso ou outro dependendo dessa condição.
Ou seja, há três tipos de fluxo de controle: sequencial, de repetição e condicional.

Nesse capítulo:
Sequencial
Iterativo (loop)
for: inicialização, condição, pós
for: hierarquicamente
for: condição ("while")
for: ...ever?
for: break
for: continue
Condicional
declarações switch/case/default
não há fall-through por padrão
criando fall-through
default
múltiplos casos
casos podem ser expressões
se resultarem em true, rodam
tipo
if
bool
o operador "!"
declaração de inicialização
if, else
if, else if, else
if, else if, else if, ..., else


# 01 Loops
OBS: falando do ponto e virgula do FOR, em Go o compilador coloca automaticamente um ";" no fim de cada instrução ou statement o compilador coloca automaticamente para nós
For
Inicialização, condição, pós
Ponto e vírgula?
gobyexample.com
# OBS: Não existe while!

# 02 Fluxo de Controle – 3. Loops: nested loop (repetição hierárquica)
For
Repetição hierárquica
Exemplos: relógio, calendário

# 03 For: inicialização, condição, pós
For: inicialização, condição, pós
For: condição ("while")
For: ...ever? (http servers)
For: break
golang.org/ref/spec#For_statements, Effective Go
(Range vem mais pra frente.)

# 04 For Break e Continue
Operação módulo: %
For: break
For: continue
Go Playground: https://play.golang.org/p/gpKMP1wAEM & https://play.golang.org/p/8erMGEbZQix

# 05 Desafio surpresa!
Format printing:
Decimal       %d
Hexadecimal   %#x
Unicode       %#U
Tab           \t
Linha nova    \n
Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
Solução: https://play.golang.org/p/REm2WHyzzz

# 06 Declaração IF
If: bool
If: o operador não → "!"
If: declaração de inicialização
Go Playground: https://play.golang.org/p/6nq2Tjb07i

If, else.
If, else if, else.
If, else if, else if, ..., else.
Go Playground: https://play.golang.org/p/18VrRX2pec


# 07 Declaração Switch

Switch:
pode avaliar uma expressão 
switch statement == case (value)
default switch statement == true (bool)
Obs não há fall-through por padrão temos que por a palavra reservada Fallthrough.
criando fall-through
default
cases compostos

# 07 Declaração Switch

Switch:
pode avaliar uma expressão 
switch statement == case (value)
default switch statement == true (bool)
Obs não há fall-through por padrão temos que por a palavra reservada Fallthrough.
criando fall-through
default
cases compostos

OBS: POdemos por a variavel do switch como Generica, usadoa palavra reservada Interface, como isto ela vai filtrar de a cordo com o Tipo. 
