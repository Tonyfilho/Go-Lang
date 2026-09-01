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

#	1º FUNÇÃO SORT ordena DADOS e recebe como Paramentro a  interface.Interface func Sort(data Interface)
#   2º A interface.Interface tem 3 metodos { Len() int; Less(i, j int) bool; Swap(i, j int) }
#   3º Podemos extender  Automaticamente a FUNÇÃO SORT se implementarmos os Metodos assima. e 
	* com isto podemos costumisar o que queremos, caso Ex: poderiamos receber um [] Slice de uma lista VIP, onde ordenariamos
	* os melhores clientes 
# 4º Temos que criar os TIPOS.
#	5º Criamos os Metodos para os TIPOS , estes Metodos costumisados
	* fazem com que estes TIPOS implemente AUTOMATICAMENTE a interface.Interface tem a FUNçÃO SORTE
#	6º Como os TIPOS Implementam a inteface.Interface podemos usar a Função neles, e desta forma ordenamos de forma customisado
	