

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