#### Aula Canais – 6. Convergência

https://www.youtube.com/watch?v=VJyryKEMleU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=152
https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/21_canais/07

# Observamos convergência quando informação de vários canais é enviada a um número menor de canais.
Interessante:
Na prática, exemplos:
# Exemplo01. Todd: 
Canais par, ímpar, e converge. 
Func send manda pares pra um, ímpares pro outro, depois fecha.
Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
Por fim um range retira todas as informações do canal converge.
