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