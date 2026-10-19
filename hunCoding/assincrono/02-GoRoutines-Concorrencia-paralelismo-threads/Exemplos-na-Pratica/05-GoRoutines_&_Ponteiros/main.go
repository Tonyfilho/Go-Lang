package main

import (
	"fmt"
	"time"
)

/**Uma função parecida com a anterior, mas que recebe um ponteiro para uma string*/
func funcPointer(message *string) {
	/**Ao inves de dar o loop 3x iremos dar um loop infinito ate que seja interrompido*/
	for {
		/**Iremos printar o valor que está no endereço de memoria */
		fmt.Println(*message)
		time.Sleep(1 * time.Nanosecond) // Simula um atraso de 1 segundo
	}
}

func main() {
	fmt.Println("GoRotines-e Ponteiros de Função")
	/**Como fazer com que varias goRoutines em concorrencia tenham o mesmo valor sendo
	que são execudadas em concorrencia e tempos diferentes? **/
	/**Neste exemplo usaremos Ponteiros*/
	var test string = "teste de ponteiro"
	/**Criando um var pointTest q é um Ponteiro de string, onde recebemos o endereço de memoria da variável test*/
	var pointTest *string = &test

	/**Chamaremos a func funcPointer e passaremos o ponteiro como argumento*/
	/*estamos instaciando uma função e passando o ponteiro como argumento*/
	go funcPointer(pointTest)

	/**Faremos a função Main dormir por 1 segundo */
	time.Sleep(time.Second)

	fmt.Println("Fim do programa  Antes de mudar o valor da variável test")

	/**se dermos um go run main.go veremos que todas as go routines estão imprimindo o mesmo valor */
	/**mas queremos mudar o valor apontado pelo ponteiro */

	/**Por esta domindo por 1 Segundo e estarmos usando ponteiros, podemos mudar o valor da variável test e todas as go routines irão imprimir o novo valor */

	*pointTest = "Novo valor do ponteiro"
	/**daremos um sleep para ver o valor mudar */
	time.Sleep(time.Second * 1) // Simula um atraso de 3 segundos
	/**Temos uma saida e inicial e depois temos a mudança de valor*/
	//teste de ponteiro
	//teste de ponteiro
	//teste de ponteiro
	//Novo valor do ponteiro
	// Novo valor do ponteiro
	// Novo valor do ponteiro
	// Novo valor do ponteiro
	fmt.Println("Fim do programa")

	/**NOrmamente usamos Chanels para comunicação entre go routines e ele tem benefícios em termos de sincronização e segurança **/
}
