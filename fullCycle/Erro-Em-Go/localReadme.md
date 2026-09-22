#### Aula Video04 Essa é a técnica para tratar erros em Golang Exemplo 01

https://www.youtube.com/watch?v=f--pS45o_zg&list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg&index=4

# O erro no GoLang é tradado de forma diferente, sem Try and Catch, o erro m Go tem as mesma prioridade do que os dados.

#  No Go o Erro é explicitos na maioria das funções
em Go podemos retornar mais de 1 valor em um função, com isto podemos retornar o valor e o error caso exista.

# O grande diferencial é que podemos escolher o que faremos com erro.
Ex: Posso dar Panic, Logfatal, ou chamar um outro recurso etc

#### Aula Video04 Essa é a técnica para tratar erros em Golang Exemplo 02

# Criando seu proprio error

Criaremos um função simples onde teremos o retorno de um dado e do erro caso haja

Caso não queira usar a variavel de erro, poderá usar Blank Indentifier "_"
ou seja o underscore no lugar da var err

# Ex res, _ := soma(1, 10) com isto não temos a continuação do error, mas como saberá

è uma variavel que exclui o dado para não ser utilizado