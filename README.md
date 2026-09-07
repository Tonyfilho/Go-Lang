#### Go-Lang
Curso de Go
## Guias de pesquisa 
# golang.org/ref/spec#For_statements
# Effective Go

## Operador NIL em GO
# Em resumo: nil em Go é um valor válido que indica "não aponta para nada", enquanto undefined em JS é a ausência de valor propriamente dita. São conceitos diferentes com propósitos diferentes!

Go (nil apenas para):
Ponteiros (*T)
Slices ([]T)
Maps (map[T]T)
Channels (chan T)
Functions
Interfaces


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




#### Aula06 Operadores logicos condicionais e Agrupamentos de dados

# 01 Operadores Lógicos
&&
||
!
Go Playground: https://play.golang.org/p/MFwrt93xlc
Qual o resultado de fmt.Println...
true && true
true && false
true || true
true || false
!true


#### Aula06B    Tabela de Agrupamentos
# Tabela Comparativa Rápida
# Tipo	    Tamanho	     Mutável	Por Valor/Ref	Nil possível?	Mais usado
# Array	    Fixo	      Sim	        Valor	         ❌	         Raro
# Slice	    Dinâmico	  Sim	        Referência	     ✅	         MUITO
# Map	    Dinâmico	  Sim	        Referência	     ✅	         MUITO
# Struct    Fixo	      Sim	        Valor	         ❌         * MUITO
# Ponteiro	Fixo	      Sim	        Referência	     ✅	         Médio
# Interface	Dinâmico	  N/A	        Referência	     ✅	         MUITO
*Struct pode ser nil se for ponteiro para struct

# 02 Agrupamentos de Dados

Estruturas de dados, ou agrupamentos de dados, nos permitem agrupar valores diferentes. Estes valores podem ser ou não do mesmo tipo.
As estruturas que veremos são: arrays, slices, structs e maps.
Vamos começar com arrays. Arrays em Go são uma fundação, e não algo que utilizamos todo dia.
Seu tamanho deve estar presente na declaração: var x [n]int
Atribui-se valores a suas posições com: x[i] = y (0-based)
Para ver o tamanho usa-se: len(x)
ref/spec: "The length is part of the array's type" → [5]int != [6]int
Effective Go: Arrays são úteis para [umas coisas que a gente não vai fazer nunca] e servem de fundação para slices. Use slices ao invés de arrays.
Go Playground: https://play.golang.org/p/Fv-sDF-ryZ

# 03 Slices
O que são tipos de dados compostos? 
Wikipedia: Composite_data_type
Effective Go: Composite literals
ref/spec: Composite literals
Uma slice agrupa valores de um único tipo.
Criando uma slice: literal composta → x := []type{values}
Go Playground: https://play.golang.org/p/W7Cxm8NPZC

# 03B Dividindo um SLICE, fatiando ou deletando de uma fatia
x[:] // SEM RANGE, VEM TODOS OS ITENS é o default, 
x[a:], x[:b], x[a:b]
"a" é incluso;
"b" não é.
Exemplo: cabeça magnética de um disco rígido (relógio, fita).
Off-by-one error.
Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
É fatiando que se deleta um item de uma slice. Na prática:
x := append(x[:i], x[:i]...)
Go Playground: https://play.golang.org/p/xK2HwCqvwd
Exercício: tente acessar todos os itens de uma slice sem utilizar range.
Solução: https://play.golang.org/p/aUC9qVCobH
Como este conteúdo foi criado

# 03C Slice e a Função Append()

Effective Go: append (package builtin)
x = append(slice, ...values)
x = append(slice, slice...)
Todd: unfurl → desdobrar, desenrolar
Nome oficial: enumeration
Go Playground: https://play.golang.org/p/RpkDCTumpT
... o operador  unFurl é o mesmo que rest ou destruction

# 03D Slice Make

Slices são feitas de arrays.
Elas são dinâmicas, podem mudar de tamanho.
Sempre que isso acontece, um novo array é criado e os dados são copiados.
É conveniente, mas tem um custo computacional.
Para otimizar as coisas, podemos utilizar make.
make([]T, len, cap)
"The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
len(x), cap(x)
x[n] onde n é maior que len é out of range. Use append.
Append maior que cap modifica o array subjacente.
pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
Effective Go.
Go Playground: https://play.golang.org/p/e8GWzyEEL8


# 03E Slice Multi-Dimentional

Slices multi-dimensionais são slices que contem slices.
São como planilhas.
[][]type
Go Playground: https://play.golang.org/p/vKyHiG1GtM
Só pra sacanear: https://play.golang.org/p/ZSU_8eJ9Yp

# 03F Slice a surpresa do array subjacente
OBS: Modificação da referencia de memoria
Isso tudo aqui a gente já viu:
Toda slice tem um array subjacente.
Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
Exemplo:
x := []int{...números}
y := append(x[:i], x[:i]...)
pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
Ou seja, y utiliza o mesmo array subjacente que x.
O que nos dá um resultado inesperado.
Ou seja, bom saber de antemão pra não ter que aprender na marra.
Go Playground: https://play.golang.org/p/BBJLuIjU_i





#### Maps

## comparações

## 3. Comparação com outras linguagens
# Linguagem	Estrutura Chave-Valor	Set (Valores Únicos)
# Go	map[K]V	Não tem nativo (simula com map[K]struct{})
# Python	dict	set
# JavaScript	Object / Map	Set
# Java	HashMap<K,V>	HashSet<T>
# C#	Dictionary<TKey,TValue>	HashSet<T>
# Rust	HashMap<K,V>	HashSet<T>

# 01 Maps Introdução

Utiliza o formato key:value.
E.g. nome e telefone
Performance excelente para lookups.
map[key]value{ key: value }
Acesso: m[key]
Key sem value retorna zero. Isso pode trazer problemas.
Para verificar: comma ok idiom.
v, ok := m[key]
ok é um boolean, true/false
Na prática: if v, ok := m[key]; ok { }
Para adicionar um item: m[v] = value
Maps não tem ordem.
Go Playground: https://play.golang.org/p/JXDdJan8Ev


# 02 Maps  range & deletando

Range: for k, v := range map { }
Reiterando: maps não tem ordem e um range usará uma ordem aleatória.
Go Playground: https://play.golang.org/p/6zEMfIP-AE
delete(map, key)
Deletar uma key não-existente não retorna erros!
Go Playground: https://play.golang.org/p/0uuIicU3Zz


#### Struct
# https://go.dev/ref/spec

# 01 Struct inicio https://go.dev/ref/spec#Struct_types
Struct é um tipo de dados composto que nos permite armazenar valores de tipos diferentes.
Seu nome vem de "structure," ou estrutura.
Declaração: type x struct { y: z }
Acesso: x.y
Exemplo: nome, idade, fumante.
Go Playground: https://play.golang.org/p/5i0DqxuBp1

OBs: O Struct é o Objeto do Js e TypeScript

# 02 Structs 

É importante se familiarizar com a documentação da linguagem Go.
Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
Veremos:
ref/spec
Já vimos mais da metade dos tipos em Go!
Struct types.
x, y int
anonymous fields
promoted fields
Go Playground: https://play.golang.org/p/z9UQej4IQT
## OBS: O acesso aos campos é iqual ao Java Ex: pessoa.idade

# 03 Structs Anonimos

São structs sem identificadores.
x := struct { name type }{ name: value }
Go Playground: https://play.golang.org/p/xyhNnSCu1f



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



#### Aula10 Funções
# Obs: como temos mais de 1 File e o package é o mesmo temos que adcionar todos eles na execução Ex: go run main.go pessoa.go

# 01 Funções Inicio

Um método é uma função anexada a um tipo.
Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
Pode-se anexar uma função a um tipo utilizando seu receiver.
Utilização: valor.método()
Exemplo: o tipo "pessoa" pode ter um método oibomdia()
Go Playground: https://play.golang.org/p/tQtoqUBpY5


#### Aula10B Funções, Interfaces & polimorfismo

# 01 Polimorfismo em Go

# Declaração: keyword identifier type → type x interface
Em Go, valores podem ter mais que um tipo.
Uma interface permite que um valor tenha mais que um tipo.
## Em Go um Interface é um conjunto de Metodos, e como se fosse um Ovo , quem herdar implementa o tipo de ovo, Ex: Ovo de Avestruz, de Galinha, de Cordona etc.. cada um em tamanhos diferentes, mas todos tem clara e gema.

Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.

## OBS: Em Go, todos os Tipos Criados (type) implementam automaticamente a Interface{} VAZIA.
Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
Esse tipo será o seu tipo e também o tipo da interface.

## OBS: Em Go não preciso implementar, caso seja contruido os mesmo Metodo que a Interface Ovo tem, Automaticamente fica implementado.


Exemplos:
Os tipos profissão1 e profissão2 contem o tipo pessoa
Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
Implementam a interface gente
Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF


## OBS: Implementação Automatica
A Interface q por sua vez está dentro do  Metodo SerHumano(g gente) e temos 2 STRUCTs que usam o mesmo NOME do Metodo que temos que implementar da Interface
, basta passar o valor que temos na Interface Gente. Observe que: Não tenho ftm.print() sendo invocado pelo metodo serHumando(g gente), mesmo assim conseguimos imprimir, ou seja: 
Em Go não precisamos IMPLEMENTAR as Interfaces, elas são automaticamente Implementadas, basta usar o mesmo nome de Metodo. Ex:Não precisei invocar oiBomDIa(), somente  passei a variavel localArquiteto.

# Temos TIPOS diferentes (Ou seja classes que usam a mesma interface)


#### Aula10C Interfaces Auto Implementáveis E do Switch Case de TIPOs
# Obs: como temos mais de 1 File e o package é o mesmo temos que adcionar todos eles na execução Ex: go run main.go pessoa.go

## Uso da mesma interface em TIPOS diferentes e Auto Implenents 
# em Go um interface Auto Implementa, bastando que tem for usar ter o mesmo nome de seus Metodo(s).
Neste exemplo abaixo veremos que a mesma Interface calcula um circulo e um retangulo, é a aplicação do S do Solid
https://gobyexample.com/interfaces  

https://play.golang.org/p/zGKr7cvTPF
Go Playground ao vivo: 
    https://play.golang.org/p/njiKbTT20Cr
Onde se utiliza?
Área de formas geométricas (gobyexample.com)
Sort
DB
Writer interface: arquivos locais, http request/response
Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.


# Obs: ao uso do ... o que chamo de Destruction ou Rest
type arquiteto struct {
	pessoa           // desta forma tenho acesso ao modo inteligente de Rest ... ou Destruction ... bastando passar variavel.nome
	tipoDeConstrucao string
}

type dentista struct {
	pessoa  pessoa // desta forma o modo inteligente do ... Rest ou Destructions não funciona, tenho q fazer variavel.pessoa.nome
	denteExtraidos int
	salarios  float64
}


#### Aula11 Funções Anônimas

## Muito usado para Go Roroutine e Funções descartaveis

Anonymous self-executing functions → Funções anônimas auto-executáveis.
func(p params) { ... }()
Go Playground: https://play.golang.org/p/Rnqmo6X6jh

# Nas funções anonimas não temos nome e para fazer a invocação usamos (x), onde X é a variavel

Aqui declara e executa ao mesmo tempo
Vamos ver bastante quando falarmos de goroutines.

As Funções anomimas podem ter paramentros Varidico Ex: x ... int e Retorno iqualmente as outras

#### Aula11A Func como expressão

f := func(p params){ ... }
f()
Go Playground: https://play.golang.org/p/cPxhPUbfLy

Podemo usar uma função como se fosse uma variavel, Ou seja estou atribuido uma Varivel um valor
após a execução e invocaremos como uma função nomal 


#### Aula11B Func Retorno Retornando uma função ou Recursividade

Pode-se usar uma função como retorno de uma função
Declaração: func f() return
Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
????: fmt.Println(f()())
Go Playground: https://play.golang.org/p/zPjoWNrCJF

# OBS: Função recursiva
1º func é para criar uma função Ex: retornaumafuncao()
2º func é para criar a tipagem Ex: func(int) int, neste caso recebe um int e retorna um int
3º func é o retorno da função que retorna outra função
1º retorno é o retorno da função criada 
2º retorno é o retorno da função retornada
func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}

#### Aula12 Callbacks, são funções que recebe um Argumento outra Função muito comum em JS

Primeiro veja se você entende isso: https://play.golang.org/p/QkAtwMZU-z
Callback é passar uma função como argumento.
Exemplo:
Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
Go Playground:
Desafio: Crie uma função no programa acima que utilize somente os números ímpares.

t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)

## OBS

A diferença é sutil mas importante:

Parâmetro (Parâmetro Formal)
São as variáveis declaradas na definição da função

Servem como "espaços reservados" que receberão valores

Fazem parte da assinatura da função

Argumento (Argumento Real)
São os valores reais que você passa para a função quando a chama

São os dados concretos que preenchem os parâmetros

#### Aula12B Closure
# https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=96

Muito parecido com Callback, sem receber a função com argumento mas retorna valores internos do escopo.
# Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
## Closures é quando capturamos um determinado ESCOPO e usamos este valor para algo
Package-level scope
Function-level scope
Code-block-in-code-block scope
Exemplo de closure:
func i() func() int { x := 0; return func() int { x++; return x } }
Quando fizermos a := i() teremos um scope, um valor para x.
Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
Closures nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código.
Go Playground: https://play.golang.org/p/AdFciYwI2Z

# OBS: Para cada invocação é somando o valor de X, nas funções Callback o retorno fica de fora do ESCOPO da função interna, aqui o retorno fica DENTRO do Escopo, com isto cada invocação somará a invocação anterior neste caso .
Ex: x++  return x
O Retorno fica dentro da função callback e não do lado de fora 

func i() func() int {
	// OU seja esta variavel X terá valores diferentes para cada Referencia de Memoria
	// E ainda somará o valores nas mesma Referencia de memoria.
	x := 0 // será usado uma variavel do escopo Externo, gerando copias diferentes para cada Expressão Criada 

	// aqui em baixo termos o CLOSURE onde retornamos o valor deste escopo interno{...} e não extermo{{...}}
	return func() int {
		///Ou seja para cada Invocação, será somado os valor de X
		x++
		return x
	}
}


#### Aula13 Func Recursividade
## OBS: Podemos usar Looks no lugar de recursividade

https://www.youtube.com/watch?v=1-pop5h5RAs&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=97

# https://pt.wikipedia.org/wiki/Efeito_Droste

# https://pt.wikipedia.org/wiki/Matriosca

# https://pt.wikipedia.org/wiki/Fractal


## è uma função que no seu retorno é chamado ela mesma,
WP: "The most common application of recursion is in mathematics and computer science, where a function being defined is applied within its own definition."
Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
No estudo de funções: é uma função que chama a ela própria.
Exemplo: fatoriais.
4! = 4 * 3 * 2 * 1 (e no zero, deu.)
Com recursividade. Go Playground: https://play.golang.org/p/ujsLnUhRp_
Com loops. Go Playground: https://play.golang.org/p/F2VsUjYVhc




#### O que são Ponteiros

https://www.youtube.com/watch?v=l2YJ-5GpGr8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=110

https://pt.wikipedia.org/wiki/Ponteiro_(programa%C3%A7%C3%A3o)


Em programação, um ponteiro ou apontador é um tipo de dado de uma linguagem de programação cujo valor se refere diretamente a um outro valor alocado em outra área da memória, através de seu endereço. Um ponteiro é uma simples implementação do tipo referência da Ciência da computação.

## ARQUITETURA

Em programação, um ponteiro ou apontador é um tipo de dado de uma linguagem de programação cujo valor se refere diretamente a um outro valor alocado em outra área da memória, através de seu endereço. Um ponteiro é uma simples implementação do tipo referência da Ciência da computação.


## O Que é DeRerence
https://pt.wikipedia.org/wiki/Refer%C3%AAncia_(ci%C3%AAncia_da_computa%C3%A7%C3%A3o)

Em ciência da computação, uma referência é um tipo de dado que contém informação que indica dados armazenados em algum outro local ao invés de conter o próprio dado. Acessar o valor referenciado pela referência chama-se "dereferenciar". Referências são fundamentais para construir várias estruturas de dados (como uma lista ligada), e para transportar informação entre diferentes partes de um programa de computador.

# Analogia de DeReference
Uma referência pode ser comparada ao endereço de uma residência. Ela é um pequeno identificador com o qual é possível encontrar um objeto potencialmente muito maior. Encontrar uma residência a partir de seu endereço é como dereferenciar uma referência.




Todos os valores ficam armazenados na memória.
Toda localização na memória possui um endereço.
Um pointeiro se refere a esse endereço.
Notações:
&variável mostra o endereço de uma variável
%T: variável vs. &variável
*variável faz de-reference, mostra o valor que consta nesse endereço
????: *&var funciona!
*type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
Exemplo: a := 0; b := &a; *b++
Go Playground: https://play.golang.org/p/gC1qGFUYrV

## OBS todos o valores em um endereço de memoria, podemos acessar estes endereços e ficar manipulandos como se fosse  uma variavel

a = Casa na Rua A, nº 10 (valor: 10)
b = Papel com endereço da casa (valor: &a)

Endereço de a = Localização física da casa (Rua A, nº 10)
Endereço de b = Localização física do papel (na sua mão)

São locais diferentes!



#### Aula15 Documentação JSON
https://www.youtube.com/watch?v=jnnIgvV0_yA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=113

https://pkg.go.dev/encoding/json
https://go.dev/blog/json
https://pkg.go.dev/encoding/json#example-Marshal
# temos o  que converte Json para Go 
https://mholt.github.io/json-to-go/ 


Já entendemos ponteiros, já entendemos métodos. Já temos o conhecimento necessário para começar a utilizar a standard library.
Nesse vídeo faremos uma orientação sobre como abordar a documentação.
Essa aula não foi preparada. Vai ser tudo ao vivo no improviso pra vocês verem como funciona o processo.
golang.org → Documents → Package Documentation 
godoc.org → encoding/json
files
examples
funcs
types
methods

type Message struct {
    Name string
    Body string
    Time int64
}

m := Message{"Alice", "Hello", 1294706395881547000}

we can marshal a JSON-encoded version of m using json.Marshal:

# Nesta instancia recebemos o Bite e o Error
b, err := json.Marshal(m)

# Decoding A Func retorna mutiliplos valores, Data e Error
To decode JSON data we use the Unmarshal function.
func Unmarshal(data []byte, v interface{}) error

# We must first create a place where the decoded data will be stored
var m Message

# and call json.Unmarshal, passing it a []byte of JSON data and a pointer to m
err := json.Unmarshal(b, &m)

# If b contains valid JSON that fits in m, after the call err will be nil and the data from b will have been stored in the struct m, as if by an assignment like:
m = Message{
    Name: "Alice",
    Body: "Hello",
    Time: 1294706395881547000,
}

### Relembrando o operador NIL
## Operador NIL em GO
# Em resumo: nil em Go é um valor válido que indica "não aponta para nada", enquanto undefined em JS é a ausência de valor propriamente dita. São conceitos diferentes com propósitos diferentes!

Go (nil apenas para):
Ponteiros (*T)
Slices ([]T)
Maps (map[T]T)
Channels (chan T)
Functions
Interfaces


#### aula15B Go Marchal (Ordenação) em Json
# De Go para Json
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=114

Exemplo: transformando structs em Go em código JSON.
No improviso tambem.
Go Playground: https://play.golang.org/p/_JvCOlK-H9



#### Aula15 C UnMarshal (Desordenando) Json
# De Json para GO
https://www.youtube.com/watch?v=mcbj-wy8Ro8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116
https://pkg.go.dev/encoding/json#Unmarshal
https://mholt.github.io/json-to-go/

E agora o contrário.
JSON-to-Go
Marshal/unmarshal vs. encoder/decoder
Marshal vai pra uma variável
Tags
Encoder "vai direto"
# `json:"Nome"`  Isto são as Tags Encoder
Ex de uso das Tags Encoder no Campo Profissão que recebe Trabalho, poderia ser o contrario tb
Go Playground: https://play.golang.org/p/l6wbuLu1NS
Com Encoder: https://play.golang.org/p/Pgwr0O07aL

## Precisamos Criar uma função de UnMarshal
# O retorno vem Nulo ou em um Ponteiro `v` , tem que haver os mesmo campos, caso contrario teremos error `errors.ErrUnsupported`

func Unmarshal(data []byte, v any) error

A função `Unmarshal` analisa os dados codificados em JSON e armazena o resultado no valor apontado por `v`. Se `v` for nulo ou não for um ponteiro, `Unmarshal` retorna um erro `InvalidUnmarshalError`.

A função `Unmarshal` utiliza o inverso das codificações usadas por `Marshal`, alocando mapas, fatias e ponteiros conforme necessário, com as seguintes regras adicionais:

Para desserializar JSON em um ponteiro, `Unmarshal` primeiro trata o caso em que o JSON é um literal JSON nulo. Nesse caso, `Unmarshal` define o ponteiro como nulo. Caso contrário, `Unmarshal` desserializa o JSON para o valor apontado pelo ponteiro. Se o ponteiro for nulo, `Unmarshal` aloca um novo valor para ele apontar.

A entrada JSON é decodificada de acordo com as seguintes regras:

Se o tipo de valor implementar `jsonv2.UnmarshalerFrom`, o método `UnmarshalJSONFrom` será chamado para decodificar o valor JSON. Se o método retornar `errors.ErrUnsupported`, a entrada será decodificada de acordo com as regras subsequentes.

Se o tipo de valor implementar `Unmarshaler`, o método `UnmarshalJSON` será chamado para decodificar o valor JSON, inclusive quando a entrada for um JSON nulo.

Se o valor implementar `encoding.TextUnmarshaler` e a entrada for uma string JSON, o método `UnmarshalText` será chamado com a string sem aspas.

#### Aula15D UnMarshal E Marshal com as Funções NewDecoder NewEncoder
# Continuação do video depois do minuto 10
https://www.youtube.com/watch?v=mcbj-wy8Ro8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116
https://cs.opensource.google/go/go/+/refs/tags/go1.27.0:src/encoding/json/v2_stream.go;l=38


Marshal/unmarshal vs. encoder/decoder
Marshal vai pra uma variável
Encoder "vai direto"

## Resumo, todo Encoder que eu fizer é para IR direto para Interface os.Stdout, sem variaveis
	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(toniFilho) //{"Nome":"Tony","SobreNome":"Filho","Idade":40,"Profissao":"Developer","ContaBancaria":1000}

Com Encoder: https://play.golang.org/p/Pgwr0O07aL

## Aqui podemos fazer Encoder de forma mais dinamica

func NewDecoder(r io.Reader) *Decoder

func NewDecoder(r io.Reader) *Decoder
NewDecoder retorna um novo decodificador que lê de r.

O decodificador introduz seu próprio buffer e pode ler dados de r além dos valores JSON solicitados.

func (*Decoder) Buffered ¶
adicionado em go1.1
func (dec *Decoder) Buffered() io.Reader
Buffered retorna um leitor dos dados restantes no buffer não lido, que pode conter zero ou mais bytes. Esses são os dados já consumidos da entrada io.Reader, mas ainda não lidos por uma chamada Decoder.Decode ou Decoder.Token. Podem conter bytes que não formam um JSON válido, pois ainda não foram validados de acordo com a gramática JSON. A quantidade exata de dados em buffer é um detalhe de implementação do Decoder e pode mudar ao longo do tempo.

É responsabilidade do chamador concatenar este buffer com o restante do leitor de entrada para obter a sequência completa de bytes após o último valor JSON decodificado.

O leitor é válido até a próxima chamada a `Decoder.Decode` ou `Decoder.Token`.

func (*Decoder) Decode ¶
func (dec *Decoder) Decode(v any) error
Decode lê o próximo valor codificado em JSON de sua entrada e o armazena no valor apontado por `v`.

Consulte a documentação de `Unmarshal` para obter detalhes sobre a conversão de JSON em um valor Go.


#### aula16 A Interface WRITER
https://www.youtube.com/watch?v=S4hEdA0RPVI&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116

https://pkg.go.dev/io#Writer

## A interface writer do pacote io.
# OBS O Tipo File Implementa a Interface Writer, por isto que Escrevemos
## Recebe um Slice de [] Bytes e retorna um Numero e um Erro quando houver, senão um nil
# Para implementar precisamos de uma Função que recebe o Metodo Write(p []byte) (n int, err error)
type Writer interface {
	Write(p []byte) (n int, err error)
}

# type Writer interface { Write(p []byte) (n int, err error) } 
Recebemos aqui um WRITE 
pkg os:   func (f *File) Write(b []byte) (n int, err error)

# Ex: pkg json: func NewEncoder(w io.Writer) *Encoder Esta função recebe um io.WRITER e retorna um Ponteiro de *Encoder

"Println [...] writes to standard output."
func Println [...] return Fprintln(os.Stdout, a...)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
Stdout: NewFile(uintptr(syscall.Stdout), "/dev/stdout") (Google: Standard streams)
func NewFile(fd uintptr, name string) *File
func (f *File) Write(b []byte) (n int, err error)
Exemplo:
Println
Fprintln os.Stdout
io.WriteString os.Stdout
Ou:
func Dial(network, address string) (Conn, error)
type Conn interface { [...] Write(b []byte) (n int, err error) [...] }

Resumo
1º Temos a Interface Write que tem o Metodo Write type Writer interface { Write(p []byte) (n int, err error) }
2º temos uma função: func (f *File) Write(b []byte) (n int, err error) que implementa este metodo       da Interface Write
3º Por fim temos uma Função que recebe dados pela Interface io.Write e retorna um Pondeiro      : func NewEncoder(w io.Writer) *Encoder
A Grande questão é que esta Interface não é so usada Imprimir, é Implementada em Muita coisa em GO!
Ex: É a Função de Conexão func DIAL que implementa a Interface Conn , mas a  interface Conn implementa o metodo WRITE:
Função func Dial(network, address string) (Conn, error)
O Seja TUDO que tem o METODO WRITE pode ser recebido como ARGUMENTO
Ate uma conexão de REDE é passado como WRITE

#### Aula17 Função do Pacote Sort
## Conhecendo o Sort e na aula17B, fazendo  modificações
# O Sort é um Função, e com isto podemos costumizar
https://www.youtube.com/watch?v=b67JIGYM6Hc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=117

Sort serve para ordenar slices.
Como faz?
golang.org/pkg/ → sort
godoc.org/sort → examples
Sort altera o valor original!
Exemplo: Ints, Strings.
Go Playground: 
sort.Strings: https://play.golang.org/p/Rs1NVwmg7h
sort.Ints: https://play.golang.org/p/I2_vsHujZa
Aprenda Go explora as funcionalidades do pacote s

#### Aula17B Usando a interface.Interface

## Criando nosso Proprio Sort
https://www.youtube.com/watch?v=0E-q22d3QD4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=118

O sort que eu quero não existe. Quero fazer o meu.
# Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
## type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
E aí posso fazer do jeito que eu quiser.
Exemplo:
struct carros: nome, consumo, potencia
slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
tipo ordenarPorPotencia
tipo ordenarPorConsumo

Go Playground: https://play.golang.org/p/KOIhAsE3OK

#	1º FUNÇÃO SORT ordena DADOS e recebe como Paramentro a  interface.Interface func Sort(data Interface)
#   2º A interface.Interface tem 3 metodos { Len() int; Less(i, j int) bool; Swap(i, j int) }
#   3º Podemos extender  Automaticamente a FUNÇÃO SORT se implementarmos os Metodos assima. e 
	* com isto podemos costumisar o que queremos, caso Ex: poderiamos receber um [] Slice de uma lista VIP, onde ordenariamos
	* os melhores clientes 
# 4º Temos que criar os TIPOS.
#	5º Criamos os Metodos para os TIPOS , estes Metodos costumisados
	* fazem com que estes TIPOS implemente AUTOMATICAMENTE a interface.Interface tem a FUNçÃO SORTE
#	6º Como os TIPOS Implementam a inteface.Interface podemos usar a Função neles, e desta forma ordenamos de forma customisado
	

#### Aula18 BCrypt

https://www.youtube.com/watch?v=4vCb7jmwkzM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=120
https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/16_aplicacao/bcrypt/main.go

É uma maneira de encriptar senhas utilizando hashes.
x/crypto/bcrypt
GenerateFromPassword
CompareHashAndPassword
# Sem Go Playground! No terminal, dentro da pasta do seu projeto, execute:

# Tem que criar o Modulo
# 1. Inicialize o módulo na pasta atual
go mod init aula18

# 2. Agora instale o bcrypt
go get golang.org/x/crypto/bcrypt

# 3. Seu arquivo go.mod será criado/atualizado

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


#### Aula19B Concorrência – 2. Goroutines & WaitGroups 
# Continuação
https://www.youtube.com/watch?v=4jXSU2jw3Ag&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=126

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

# Como ja vimos o main() terminou antes da goroutine go Func01() e foi descartada e precisamos usar o WaitGroup  para controlar a execução


#### Aula19D  Concorrência – 4. Na prática: Condição de corrida

https://www.youtube.com/watch?v=XxG7qqJzDKk&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=128

Aqui vamos replicar a race condition mencionada no artigo anterior.
time.Sleep(time.Second) vs. runtime.Gosched()
go help → go help build → go run -race main.go
Como resolver? Mutex. 

# Resulmo vais goroutines leram uma variavel compartilhada e não conseguiram salvar na variavel


#### Aula19E  Concorrência –  5. Mutex
https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=129
https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/18_concorrencia/06_mutex/main.go
# https://pkg.go.dev/sync#Mutex


Agora vamos resolver a race condition do programa anterior utilizando mutex.
# Mutex é mutual exclusion, exclusão mútua.
# Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
Na prática:
type Mutex
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
RWMutex

# Um Mutex é um bloqueio de exclusão mútua. O valor zero para um Mutex representa um mutex desbloqueado.

Um Mutex não deve ser copiado após o primeiro uso.

Na terminologia do modelo de memória do Go, a n-ésima chamada a `Mutex.Unlock` "sincroniza antes" da m-ésima chamada a `Mutex.Lock` para qualquer n < m. Uma chamada bem-sucedida a `Mutex.TryLock` é equivalente a uma chamada a `Lock`. Uma chamada malsucedida a `TryLock` não estabelece nenhuma relação de "sincroniza antes".






#### Aula19F Atomic
https://www.youtube.com/watch?v=iFlQ2yAYcp4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130

https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/18_concorrencia/07_atomic/main.go

# https://pkg.go.dev/sync/atomic
Agora vamos fazer a mesma coisa, mas com atomic ao invés de mutex.
atomic.AddInt64
atomic.LoadInt64

# O pacote `atomic` fornece primitivas de memória atômica de baixo nível úteis para implementar algoritmos de sincronização.

Essas funções exigem muito cuidado para serem usadas corretamente. Exceto para aplicações especiais de baixo nível, a sincronização é melhor feita com canais ou com os recursos do pacote `sync`. Compartilhe memória comunicando-se; não comunique-se compartilhando memória.

A operação de troca (swap), implementada pelas funções `SwapT`, é o equivalente atômico de:


##### Aula20 Package Organizations  8. Pacotes
## OBS tem q dar o "go mod init"

https://www.youtube.com/watch?v=SO-RFPSqD3c&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=138
https://rakyll.org/style-packages/ 

https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/19_seu-ambiente-de-desenvolvimento/pacotes


Opção 1: uma pasta, vários arquivos.
package declaration em todos os arquivos
package scope: um elemento de um arquivo é acessível de todos os arquivos
imports tem file scope
Opção 2: separando por packages.
pastas diferentes
requer imports
para usar: package.Função()
Exportado vs. não-exportado, ou seja, visível vs. não-visível
Em Go não utilizamos os termos "público" e "privado" como em outras linguagens
É somente questão de capitalização
Com maiúscula: exportado, visível fora do package
Com minúscula: não exportado, não utilizável fora do package

#### Aula21  Canais – 1. Entendendo canais

https://www.youtube.com/watch?v=jF0xuhnPkDg&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=146

# Canais são o Jeito Certo® de fazer sincronização e código concorrente.
# Eles nos permitem trasmitir valores entre goroutines.
# Servem pra coordenar, sincronizar, orquestrar, e buffering.
Na prática:
make(chan type, b)
Canais bloqueiam:
Eles são como corredores em uma corrida de revezamento
Eles tem que "passar o bastão" de maneira sincronizada
Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
Ou se um corredor ficar esperando receber o bastão, mas ninguem entregar...
...não dá certo.
Exemplos:
Poe um valor num canal e faz um print. Block.
Código acima com goroutine.
# Ou com buffer. Via de regra: má idéia; é legal em certas situações, mas em geral é melhor sempre passar o bastão de maneira sincronizada.
Interessante: ref/spec → types
Código: 
Block: https://play.golang.org/p/dClS7vQlYE (não roda!)
Go routine: https://play.golang.org/p/ZbNCwUuiPi
Buffer: https://play.golang.org/p/32vYvCR7qn
Buffer block: https://play.golang.org/p/smeW6vigAT
Mais buffer: https://play.golang.org/p/Pe2pcboGiA