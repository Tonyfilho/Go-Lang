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
# 2. Com throttling! Ou seja, com um número máximo de go funcs.
Ídem acima, mas a func que lança go funcs é assim:
Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.