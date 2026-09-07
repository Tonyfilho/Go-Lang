#### Aula10C Interfaces Auto Implementáveis E do Switch Case de TIPOs
# Obs: como temos mais de 1 File e o package é o mesmo temos que adcionar todos eles na execução Ex: go run main.go pessoa.go

## Uso da mesma interface em TIPOS diferentes e Auto Implenents 
# em Go um interface Auto Implementa, bastando que tem for usar ter o mesmo nome de seus Metodo(s).
Neste exemplo abaixo veremos que a mesma Interface calcula um circulo e um retangulo, é a aplicação do S do Solid
https://gobyexample.com/interfaces  

https://play.golang.org/p/zGKr7cvTPF
Go Playground ao vivo: 
    https://play.golang.org/p/njiKbTT20Cr
Onde se utiliza?
Área de formas geométricas (gobyexample.com)
Sort
DB
Writer interface: arquivos locais, http request/response
Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.


# Obs: ao uso do ... o que chamo de Destruction ou Rest
type arquiteto struct {
	pessoa           // desta forma tenho acesso ao modo inteligente de Rest ... ou Destruction ... bastando passar variavel.nome
	tipoDeConstrucao string
}

type dentista struct {
	pessoa  pessoa // desta forma o modo inteligente do ... Rest ou Destructions não funciona, tenho q fazer variavel.pessoa.nome
	denteExtraidos int
	salarios  float64
}