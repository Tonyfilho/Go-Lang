package main

import (
	"fmt"
	"sync"
	"time"
)

/**1º começaremos criando uma variavel global para receber o compartlhamento entre as GoRoutines*/
var sharedRsc = false

func main() {

	fmt.Println("Aula 09 Sync.COND BROADCAST.")
	fmt.Println()
	fmt.Println("Neste Exemplo  usaremos o Cond Broadcast() ")
	fmt.Println()
	/**2º Teremos um WaitGroup, pois temos varias GoRoutines em Execução*/
	var wg sync.WaitGroup
	/**3º E dentro do Add(2), teremos 2 GoRoutines executando ao mesmo tempo e esperar 2 GoRoutines */
	wg.Add(2)
	/**4º Criaremos uma Var q receba a interface do Mutex*/
	m := sync.Mutex{}
	/**5º Criaremos a Var q recebe Condition e o paramentro da interface Mutex*/
	c := sync.NewCond(&m)

	/**6º Instaciaremos varias GoRoutines e para cada teremos 1 Lock()*/
	go func() {
		c.L.Lock()
		/**Aqui validamos o valor se for false a GoRoutine fica esperando*/
		/**Enquando não receber o Sinal de BroadCast ficará validando!!! */
		for sharedRsc == false {
			fmt.Println("GoRoutines01 wait")
			c.Wait()
		}
		/**vendo time do recebimento*/
		fmt.Println(time.Now().String())
		/***Assim que emitimos o Broadcast sairemos do For*/
		fmt.Println("Sinal BroadCast emitido em SharedRSC GoRoutines01", sharedRsc)
		/**7º daremos umUnlock() fora do For*/
		c.L.Unlock()
		/**8º daremos um Done(), pois ja terminamos*/
		wg.Done()

	}()
	go func() {
		c.L.Lock()
		/**Aqui validamos o valor se for false a GoRoutine fica esperando*/
		/**Enquando não receber o Sinal de BroadCast ficará validando!!! */
		for sharedRsc == false {
			fmt.Println("GoRoutines02 wait")
			c.Wait()
		}
		/**vendo time do recebimento*/
		fmt.Println(time.Now().String())
		/***Assim que emitimos o Broadcast sairemos do For*/
		fmt.Println("Sinal BroadCast emitido em SharedRSC GoRoutines02", sharedRsc)
		/**7º daremos umUnlock() fora do For*/
		c.L.Unlock()
		/**8º daremos um Done(), pois ja terminamos*/
		wg.Done()
	}()
	go func() {
		c.L.Lock()
		/**Aqui validamos o valor se for false a GoRoutine fica esperando*/
		/**Enquando não receber o Sinal de BroadCast ficará validando!!! */
		for sharedRsc == false {
			fmt.Println("GoRoutines03 wait")
			c.Wait()
		}
		/**vendo time do recebimento*/
		fmt.Println(time.Now().String())
		/***Assim que emitimos o Broadcast sairemos do For*/
		fmt.Println("Sinal BroadCast emitido em SharedRSC GoRoutines03", sharedRsc)
		/**7º daremos umUnlock() fora do For*/
		c.L.Unlock()
		/**8º daremos um Done(), pois ja terminamos*/
		wg.Done()
	}()

	/**8º Temos q dar time Sleep para dar tempo de atulizar os valores da Var SharedRSC */
	time.Sleep(2 * time.Second)
	/**9º Neste momento temos que dar  lock(), pois queremos alterar o valor da var SharedRSC*/
	c.L.Lock()
	fmt.Println("Main()....GoRoutine ...Mundando Valor da SharedRSC.....")
	sharedRsc = true
	/**10º Emitido o sinal de Broadcast() 2 GoRoutnes q estão no Add() receberão o Sinal*/
	c.Broadcast()
	fmt.Println("Main()....GoRoutine Broadcast().....")
	/**11º dando o Unlock()*/
	c.L.Unlock()
	/**12º usando Wait() mas do WaitGroup e não do COND para finalizar */
	wg.Wait()

	fmt.Println("Fim do Main().......")
	fmt.Println()
}
