package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processTaskOrError(id int) error {
	delay := time.Duration(rand.Intn(1000)) * time.Microsecond
	time.Sleep(delay)

	if rand.Float32() < 0.2 {
		return fmt.Errorf("Erro aléatorio ao processar %d", id)
	}
	fmt.Printf("Processamento bem sucedido para %d\n", id)

	return nil
}

func main() {
	fmt.Println("Versão SEM errgroup — WaitGroup + chan error + context")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		totalTasks = 20
		maxTask    = 5
	)

	semaphore := make(chan struct{}, maxTask)
	/**Substituindo errGroup------*/
	var wg sync.WaitGroup                 //espera todas as GoRoutines
	errCh := make(chan error, totalTasks) //cada GoRoutine reporta um erro, ou tem erro ou nil
	var mu sync.Mutex
	completedTask := 0

	for i := 1; i <= totalTasks; i++ {
		i := i
		wg.Add(1) //equivale ao g.Go() registra a GoRoutines

		go func() {
			defer wg.Done() //equivale ao g.Go() marca Done ao sair
			// temos o semafaro igual a antes
			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()

			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			}

			// checa cancelamento antes de trabalhar // foi comentado pois não entendi o sentido de repetir
			// select {
			// case <-ctx.Done():
			// 	errCh <- ctx.Err()
			// 	return
			// default:
			// }

			if err := processTaskOrError(i); err != nil {
				cancel() //avisa as outras GoRoutines
				errCh <- err
				return
			}

			mu.Lock()
			completedTask++
			mu.Unlock()
			errCh <- nil //sucesso não houve erro

		}()
	}

	/**02 Substituindi o g.Wait()*/
	// precisamos esperar o WaitGroup e fechar o canal para poder ler tudo.
	go func() {
		wg.Wait()
		close(errCh)
	}()

	/**03 coletando o 1º erro caso não seja  nil*/
	var firstErr error
	for err := range errCh {
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}

	/**Mostrando Resultados**/

	if firstErr != nil {
		fmt.Printf("Processamento interrompido devido a error: %v\n", firstErr)
	} else {
		fmt.Println("Todas as tarefas foram concluidas com sucesso.")
	}

	fmt.Printf("Total de tarefas concluidas: %d/%d\n", completedTask, totalTasks)

}
