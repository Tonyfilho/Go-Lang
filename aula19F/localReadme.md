#### Aula19F Atomic
https://www.youtube.com/watch?v=iFlQ2yAYcp4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130

https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/18_concorrencia/07_atomic/main.go

# https://pkg.go.dev/sync/atomic
Agora vamos fazer a mesma coisa, mas com atomic ao invés de mutex.
atomic.AddInt64
atomic.LoadInt64

# O pacote `atomic` fornece primitivas de memória atômica de baixo nível úteis para implementar algoritmos de sincronização.

Essas funções exigem muito cuidado para serem usadas corretamente. Exceto para aplicações especiais de baixo nível, a sincronização é melhor feita com canais ou com os recursos do pacote `sync`. Compartilhe memória comunicando-se; não comunique-se compartilhando memória.

A operação de troca (swap), implementada pelas funções `SwapT`, é o equivalente atômico de:


