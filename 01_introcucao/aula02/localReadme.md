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