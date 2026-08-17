
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