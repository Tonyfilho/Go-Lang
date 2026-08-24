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




