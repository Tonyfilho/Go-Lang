#### Aula21H Divergencia Exemplo 02

https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/21_canais/08

https://www.youtube.com/watch?v=8X6eOnSJu5g&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=156

# 2. Com throttling! Ou seja, com um número máximo de go funcs.
Ídem acima, mas a func que lança go funcs é assim:
Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.


Resulmo da Lição: 
# 01º cada Trabalho leva 1S
# 02º e criamos 5 GoRoutines que no daremos resposta
# 03º ou seja no terminal teremos 5 em 5, veja.
*Ou seja na Func Outra, criar grupos de 5 goRoutines para resolver 1 Trabalho por vez
* isto prova que podemos criar um grupo de GoRoutines para execultar 1 trabalho concorrentemente.
* Imagine que temos 1 stream e dividimos o trabalho desta stream em varias Treads ou GoRoutine
