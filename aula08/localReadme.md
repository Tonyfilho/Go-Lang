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