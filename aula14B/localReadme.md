#### O que são Ponteiros e Aplicações Praticas
https://www.youtube.com/watch?v=0slBes2RYgc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=110


Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
Não queremos passar grandes volumes de dados pra lá e pra cá
Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
Exemplos:
x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)
Go Playground: https://play.golang.org/p/VZmfWfw76s

# Ponteiros em GO Pra que servem ?"
Poteiro servem para 2 coisas:
««««»»»»: 1º Quando lidamos com muitos dados e queremos que todos acessem o mesmo local e Memoria.
Ou seja. Desta forma todos que querem acessar ou mudar algo do mesmo, não precisa Criar novamente.

Em Go tudo é pass by value!, Ou seja não passamos valores e som ponteiros, pois a performace é maior!
Ex. Imagine um Função que recebe um Valor, ao invez de passar o valor, passamos o Ponteiro, ou seja vai lá e pegue naquele local
Desta forma ao invez de fazer COPIA nossa passamos o ENDEREÇO dos dados
««««»»»»: 2º É a quando queremos mudar o valor de um ENDEREÇO, sem fazer COPIAS

