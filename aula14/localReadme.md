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