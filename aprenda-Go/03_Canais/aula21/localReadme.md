#### Aula21  Canais – 1. Entendendo canais

https://www.youtube.com/watch?v=jF0xuhnPkDg&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=146

# Canais são o Jeito Certo® de fazer sincronização e código concorrente.
# Eles nos permitem trasmitir valores entre goroutines.
# Servem pra coordenar, sincronizar, orquestrar, e buffering.
Na prática:
make(chan type, b)
Canais bloqueiam:
Eles são como corredores em uma corrida de revezamento
Eles tem que "passar o bastão" de maneira sincronizada
Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
Ou se um corredor ficar esperando receber o bastão, mas ninguem entregar...
...não dá certo.
Exemplos:
Poe um valor num canal e faz um print. Block.
Código acima com goroutine.
# Ou com buffer. Via de regra: má idéia; é legal em certas situações, mas em geral é melhor sempre passar o bastão de maneira sincronizada.
Interessante: ref/spec → types
Código: 
Block: https://play.golang.org/p/dClS7vQlYE (não roda!)
Go routine: https://play.golang.org/p/ZbNCwUuiPi
Buffer: https://play.golang.org/p/32vYvCR7qn
Buffer block: https://play.golang.org/p/smeW6vigAT
Mais buffer: https://play.golang.org/p/Pe2pcboGiA