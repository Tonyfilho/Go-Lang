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

# Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
## Closures é quando capturamos um determinado ESCOPO e usamos este valor para algo
## Muito parecido com Callback, mas retorna valores internos do escopo.
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