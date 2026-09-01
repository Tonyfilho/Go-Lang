package main

import (
	"fmt"

)



func main() {
	fmt.Println("Aula 19C - Concorrência – 3.  Discussão: Condição de corrida")
	fmt.Println("Em Go, a condição de corrida ocorre quando duas ou mais goroutines acessam uma variável compartilhada ao mesmo tempo\n, e pelo menos uma delas está escrevendo nessa variável. Isso pode levar a resultados inesperados e comportamentos indesejados no programa.")
	fmt.Println("Em Go no lugar de Mutex, podemos usar o canal para sincronizar o acesso a uma variável compartilhada entre goroutines. Isso garante que apenas uma goroutine possa acessar a variável de cada vez, evitando condições de corrida.")
	

}