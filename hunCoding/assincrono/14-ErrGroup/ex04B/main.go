package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func taskOrErrorGenerated(id int) error {	
	delay := time.Duration(rand.Intn(1000)) * time.Microsecond
	time.Sleep(delay) 
	if rand.Float32() < 0.2 { 
		return fmt.Errorf("erro aleatório ao processar %d", id)
	}
	fmt.Printf("Processamento bem-sucedido para %d\n", id)
	return nil
}


func main() {
	g, ctx := errgroup.WithContext(context.Background())// cancelamento automatico
	sem := make(chan struct{}, 5)
	var mu sync.Mutex
	ok := 0

    for i := 1; i <= 20; i++ {
		i := i
		g.Go(func() error {
			select {
			case sem <- struct{}{}:
				defer func() {<- sem}()

			case <- ctx.Done():
				return ctx.Err()
			}

           if err := taskOrErrorGenerated(i); err != nil {
			return err //o errGroup cancela o ctx automaticamente 
		   }
		   mu.Lock()
		   ok ++
		   mu.Unlock()
		   return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Interrompido: ", err)
	} else {
		fmt.Println("tudo ok")
	}

	fmt.Printf("concluidas: %d/20\n", ok)


}