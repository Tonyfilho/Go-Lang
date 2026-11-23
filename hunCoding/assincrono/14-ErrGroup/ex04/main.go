package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)


/*1º Criaremos uma função para de gerar acerto ou erros aleatórios onde é tipada com interface de error*/
/**Resulmo deste exemplo ao invez de criar um Channel para cada process Task retorna-se um error*/
/**E pegarmos o error deste Channel e fossemos para select para retonar o valor no select e criar outro
channel para printar e saber quantas task deram erradas*/
/**Usaremos o g.Wait() */
func processTask(id int) error {
	//fmt.Println("erro aleatório %d", id)
	/**Aqui criaremos um DELAY randonico, teando simular um request onde acada x numeros teremos um erro*/
	delay := time.Duration(rand.Intn(1000)) * time.Microsecond
	time.Sleep(delay) // Simula algum processamento

	/**Caso o valor randomico seja menor q 0.2 retornaremos um erro com Id*/
	if rand.Float32() < 0.2 { // 20% de chance de gerar um erro
		return fmt.Errorf("erro aleatório ao processar %d", id)
	}
	/**Caso seja maior é colocado a tarefa com concluida */
	fmt.Printf("Processamento bem-sucedido para %d\n", id)
	return nil

}

/*temos q fazer uma importação da biblioteca "golang.org/x/sync/errgroup"**/
func main() {
	fmt.Println("Aula 14 Gerenciando erros dentre várias funções assíncronas em GoLang")
	fmt.Println("Neste exemplo 04 usaremos ErrGroup para gerenciar erros em funções assíncronas em GoLang")
	/*2º temos que ter uma contexto com cancelamento, para gerenciar as funções assíncronas com tempo ou parada**/
	ctx, cancel := context.WithCancel(context.Background())// Cria um contexto com cancelamento
	defer cancel()// Garante que o cancelamento seja chamado ao final da função main
	/**2ºB teremos 5 execuções simultâneas, desta saberemos quantas foram concluídas* ou deram erro*/
	maxTasks := 5
	semaphore := make(chan struct{}, maxTasks) // Limita a 5 goroutines simultâneas
	/**3º Erro Group e Mutex para tratar das funções Assíncronas*/
	var g errgroup.Group // Cria um grupo de goroutines para gerenciar funções assíncronas
	var mu sync.Mutex // Mutex para proteger o acesso à variável completedTasks
	totalTasks := 20 // Total de tarefas a serem processadas
	completedTasks := 0 // Variável para contar o número de tarefas concluídas

	/*4º Agora vamos criar um loop para gerar 20 tarefas assíncronas*/
	for i := 1; i <= totalTasks; i++ {
		i := i // Copiar a variável local para evitar problemas de concorrência
		/**5ª começaremos com o ErrGroup  para cada função*/
		g.Go(func() error {
			semaphore <- struct{}{} // Adquire um slot no semáforo vazio
			defer func() { <-semaphore }()// fica esperando a Struct terminar
            /**6º Aqui dentro do semafarto colocaremos um selec*/
              select {
				/**Caso recebermos um sinal de cancel do context terminaremos e retornaremos erro()*/
			  case <- ctx.Done():
				return ctx.Err()

			  default:
				/**Caso não haja Erro chamaremos a função de processTask()*/
				/** com o valor de 0 a 19*/
				/**E se caso aconteça algum error, chamaremos o cancel e retornaremos o erro*/
				/**Lembrando q dentro desta task temos um valor randomico de dar erro*/
				if  err := processTask(i); err != nil {
					cancel()
					return err
				}

				/**7º Por fim, caso esta tudo correto, daremos mu.Lock()*/
				mu.Lock()
				completedTasks++ //adicionaremos valor na variavel aumentando a quantidade de tasks corretas
				mu.Unlock() //terminaremos com muUnLock() da variavel

				return nil //retornaremos nil

			  }

			
		})
	}

	/*8º Aqui usaremos g.Wait() onde saberemos quem deu errado ou certo **/
	if err := g.Wait() ; err != nil {
		fmt.Printf("Processamento  interrompido devido a error: %v\n", err)
	} else {
		fmt.Println("Todas as tarefas foram cncluidas com sucesso. ")
	}
	
	fmt.Printf("Total as tarefas concluidas: %d/%d\n, " , completedTasks, totalTasks)


}
