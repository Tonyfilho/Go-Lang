#### Aula21H Divergencia Exemplo 01

https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/21_canais/08

https://www.youtube.com/watch?v=8X6eOnSJu5g&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=156

Divergência é o contrário de convergência :)
Na prática, exemplos:
# 1. Um stream vira centenas de go funcs que depois convergem.
Dois canais.
Uma func manda X números ao primeiro canal.
Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
Trabalho() é um timer aleatório pra simular workload.
Por fim, range canal dois demonstra os valores.

# 01º Começamos com 1 canal e terminaremos com 1 canal, ate aqui não temos divergenci
# 02º Madamos X numeros para 1º canal Func: Mand
# 03º Pega cada numero do 1º canal  e manda 10 Thread ou 10(GoRoutine) que abastece o canal02
* ou seja temos 10 itens no canal01 e com range criaremos 10 Goroutines ou threads no canal02.
* Ou seja dividimos ou espalhamos o TRABALHO do processador em 10 GoRoutines e todas foram processadas 
* concorrentemente e pois colocamos tudo no canal02

# Então imagine q temos um trabalho para fazer onde NÓS determinamos quantas Threads ou GoRoutine irá fazer o mesmo trabalho
  * esta é a moral deste exercicio.
  