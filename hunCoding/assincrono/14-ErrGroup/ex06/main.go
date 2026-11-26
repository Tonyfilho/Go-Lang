package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
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
	fmt.Println("Versão enxuta: errgroup.WithContext + SetLimit")
	const (
		totalTasks = 20
		maxTasks   = 5
	)

	/**1º errGroup com contexto; cancela sozinho no 1º erro*/
	g, ctx := errgroup.WithContext(context.Background())

	/**2º limite de concorrencia: substitui o chan struct{} semaforo*/
	g.SetLimit(maxTasks)

	/**3º Mutex só para contador(O resto o errGroup cuida*/
	var mu sync.Mutex
	completedTask := 0

	for i := 1; i <= totalTasks; i++ {

		i := i
		g.Go(func() error {
			// checa cancelamento antes de começar
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:

			}

			if err := processTaskOrError(i); err != nil {
				return err // erroRoup cancela o ctx automaticamente
			}

			mu.Lock()
			completedTask++
			mu.Unlock()
			return nil // caso esteja tudo correto, retornaremos nil

		})

	}
	if err := g.Wait(); err != nil {
		fmt.Printf("Processamento interrompido devido a error: %v\n", err)
	} else {
		fmt.Println("Todas as tarefas foram concluidas com sucesso.")
	}

	fmt.Printf("Total de tarefas concluidas: %d/%d\n", completedTask, totalTasks)

}
