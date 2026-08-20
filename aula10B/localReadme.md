


#### Funções, Interfaces & polimorfismo

# 01 Polimorfismo em Go

# Declaração: keyword identifier type → type x interface
Em Go, valores podem ter mais que um tipo.
Uma interface permite que um valor tenha mais que um tipo.
## Em Go um Interface é um conjunto de Metodos, e como se fosse um Ovo , quem herdar implementa o tipo de ovo, Ex: Ovo de Avestruz, de Galinha, de Cordona etc.. cada um em tamanhos diferentes, mas todos tem clara e gema.

Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.

## OBS: Em Go, todos os Tipos Criados (type) implementam automaticamente a Interface{} VAZIA.
Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
Esse tipo será o seu tipo e também o tipo da interface.

## OBS: Em Go não preciso implementar, caso seja contruido os mesmo Metodo que a Interface Ovo tem, Automaticamente fica implementado.


Exemplos:
Os tipos profissão1 e profissão2 contem o tipo pessoa
Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
Implementam a interface gente
Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF
Go Playground ao vivo: 
    https://play.golang.org/p/njiKbTT20Cr
Onde se utiliza?
Área de formas geométricas (gobyexample.com)
Sort
DB
Writer interface: arquivos locais, http request/response
Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.