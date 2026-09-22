package main

/**Neste caso tenho de chamar o construtor de imports */
import (
	"fmt"	
	"pacotes/math"
)

/**Veremos o comportamento da importação de Packages no Go*/

func main() {
    varNome := math.Nome
	fmt.Println("Video Aula 03 pacotes ")
	fmt.Println(" No bash ou terminal temos que criar o Arquivo go.Mod go mod init nome-do-diretorio")
	totalSoma := math.Soma(10, 5)

	fmt.Println("A Soma: ", totalSoma)
	fmt.Println()
	
	totalSubtrai := math.Subtrai(20,5)
	fmt.Println("A subtrai: ", totalSubtrai)
	fmt.Println()
	
	fmt.Println("Podemos ter acesso, a Variaveis, Funções dentro dos packages")
	fmt.Println()
	fmt.Println("Sou a Variavel do Package math: ", varNome)


	
}
