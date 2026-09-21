#### Aula21C Assignment/conversion

# Assignment/conversion: 
de geral para específico
de específico para geral não
Exemplos:
# geral pra específico: https://play.golang.org/p/H1uk4YGMBB

# específico pra específico: https://play.golang.org/p/8JkOnEi7-a 
// cs = cr // cannot use cr (variable of type <-chan int) as chan<- int value in assignment

# específico pra geral: https://play.golang.org/p/4sOKuQRHq7
// fmt.Printf("c\t%T\n", (chan int)(cs))
// fmt.Printf("c\t%T\n", (chan int)(cr))

# atribuição tipos != (Diferentes) https://play.golang.org/p/bG7H6l03VQ 
fmt.Println("teremos error de tipagem ")

Em funcs podemos especificar:
receive channel
Parâmetro receive channel: (c ←chan int)
No scope dessa função, esse canal só recebe
Não podemos fechar um receive channel
send channel 
Parâmetro send channel: (c chan← int)
No scope dessa função, esse canal só envia
Podemos fechar um send channel
Exemplo: passando informação de uma função para outra.
Código: https://play.golang.org/p/TlcSm8bHkW (replay)