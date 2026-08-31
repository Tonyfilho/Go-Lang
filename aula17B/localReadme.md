#### Aula17B Usando a interface.Interface

## Criando nosso Proprio Sort
https://www.youtube.com/watch?v=0E-q22d3QD4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=118

O sort que eu quero não existe. Quero fazer o meu.
# Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
## type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
E aí posso fazer do jeito que eu quiser.
Exemplo:
struct carros: nome, consumo, potencia
slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
tipo ordenarPorPotencia
tipo ordenarPorConsumo

Go Playground: https://play.golang.org/p/KOIhAsE3OK