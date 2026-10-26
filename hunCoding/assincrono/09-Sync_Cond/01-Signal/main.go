package main

import (
	"fmt"
	"math/rand"
	"sync"

	"time"
)

func main() {

	fmt.Println("Aula 09 Sync.Signal")
	fmt.Println()
	fmt.Println("Neste Exemplo usaremos o Signal()")
	/*Imagine 2 GoRoutines q compartilham a mesma Var**/
	/**E uma destas GoRoutine troque de VALOR e receba um Sinal quando ela trocar de valor*/
	/**Caso não recebe o Signal, ela fica presa em um FOR infinito fazendo validações*/
	/**A solução melhor a GoRoutine ser avisada quando valor for trocado*/
	/**O Sync.Condidion +  Mutex com Lock() cria esta situação e o  Signal faz o aviso*/

	/*01º criaremos valores randomicos**/
	rand.Seed(time.Now().UnixNano())
	/*02º teremos um Array q tem tamanho 10 e receberá 10 letras**/
	const N = 10
	/**para por valore neste array teremos um GoRoutine pondo valores ALEATÓRIOS*/
	/*Ex: a letra D, pode entrar depois da letra E, pois virá em horarios diferentes**/
	var values [N]string

	/*3º a criação da Var COND**/
	/**Lembrado q estamos passando o Lock() pelo paramentro da Interface Mutex*/
	/**Lembrando q poderiamos criar um Lock() proprio e não usar o Mutex, mas é reiventar a roda*/
	cond := sync.NewCond(&sync.Mutex{})

	/**4º Criaremos o FOR de 0 a N para startar as GoRoutines*/
	for i := 0; i < N; i++ {
		/**4B pegaremos um time aleatorio q será de uma tempo de 0 ate 10*/
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		/**4C com este tempo aleatorio iniciaremos um função GoRoutine em Func anonima*/
		/**4D A Func anomima rece o (i) por paramentro*/
		go func(i int) {
			/**4E Simularemos um delay na GoRoutine por X time que não sabemos*/
			time.Sleep(d)
			/**4D iniciaremos o Cond.lock() para trocar valores em uma Var q está sendo compartilhada entre 2 GoRoutines*/
			/**5 A GoRoutine q está sendo executada e a que está lendo o valor*/
			/**5B Enquanto esta variavel esta alterando valor, outras GoRoutines estarão lendo*/
			/**5C Para evitar rancing Condition damos o Lock()*/
			/**Desta forma evitaremos conflito de rancing condition*/
			cond.L.Lock()
			values[i] = string('a' + i) // a + 0 = a; a + 1 = b; a + 2 = c ....
			/**6 Depois de atribuido o valor daremos um UnLock() p q outras GoRoutines Gravem ou Leiam*/
			cond.L.Unlock()
			/**7 Enviaremos o Signal, dizendo já pode usar o valor para Rescrever ou Ler*/
			/**Foi posto um valor não sei se é o ultimo... o Signal()  falará isto, valide o valor!!*/
			cond.Signal()
			/**Toda vez q mandar um valor enviaremos 1 Signal() mas cabe que for usar VALIDAR*/

		}(i)
	}
	/**8 temos um Var q recebe uma func q irá validar para nós..*/
	/**Criamos um validador*/
	/**Toda vez q recebermos um Signal, checaremos o Array se tem espaço**/
	checkCondition := func() bool {
		fmt.Println("Valores: ", values)
		/**para cada execução checaremos todo o Array*/
		for i := 0; i < N; i++ {
			/*Se alguma posição for vazia, retornal false e vai verificar novamente*/
			/**Ou seja enquanto tivemos posições livres, ficaremos no loop*/
			if values[i] == "" {
				return false
			}
		}
		return true
	}
	/**Como ele saberá q temos ou não espeço no array??? */
	/**9, temos q dar um Cond.Lock() para poder LER a */
	cond.L.Lock()
	/**10 daremos um Defer em Unlook() q será o ultimo a executar, ou seja depois q o For terminar*/
	defer cond.L.Unlock()
	/**Daremos um For infinito, onde quando for False ele entra neste FOR e executa o Wait()*/
	for !checkCondition() {
		/**Enquanto recebermos Signal(), verificaremos Valores e enquanto o array for vazio, teremos o Wait()*/
		/**Para cada Signal, validaremos se temos espaço*/
		/**11 dentro da checagem daremos Wait()*/
		cond.Wait()

		/**A ideia do cond() é validar somente quando temos Signal(), mas não quer dizer q a condição é ideal ou satisfatoria*/
	}




	fmt.Println("Fim do Main().......")
	fmt.Println()
}
