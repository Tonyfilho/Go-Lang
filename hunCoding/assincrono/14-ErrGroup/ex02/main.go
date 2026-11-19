package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Aula 14 Gerenciando erros dentre várias funções assíncronas em GoLang")
	fmt.Println("Neste exemplo usaremos channels para gerenciar erros em funções assíncronas em GoLang")

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://x.comdfdsgfSG", // URL inválida para simular erro
	}
	var wg sync.WaitGroup
	/**1º Criaremos canais para receber os resultados das requisições */
	/**2º Mas não podemos somente receber STRINGs*/
	/**post01Ch := make(chan string)*/
	/**2B temos que criar um struct para armazenar o resultado e o erro */
	type Result struct {
		URL        string
		Err        error
		StatusCode int
	}
	urlChannel01 := make(chan Result)
	urlChannel02 := make(chan Result)
	urlChannel03 := make(chan Result)

	/**03º Como estamos dentro de um FOR qual channel deu error*/

	for _, url := range urls {

		wg.Add(1)

		go func(localUrl string) {

			defer wg.Done()

			req, err := http.Get(localUrl)
			if err != nil {
				fmt.Printf("Erro ao fazer requisição para %s: %v\n", localUrl, err)
				log.Panic(err)
			}

			defer req.Body.Close()
			/***04º depois do defer iremos ter que saber qual channel deu error*/
			/**Lembrando que caso esteja tudo bem o erro será nil*/
			/**Escolheremos o channel 03 para ser o default de erros de requisição*/

			switch localUrl {
			case "https://golang.org":
				urlChannel01 <- Result{URL: localUrl, Err: err, StatusCode: req.StatusCode}
			case "https://google.com":
				urlChannel02 <- Result{URL: localUrl, Err: err, StatusCode: req.StatusCode}
			// case "https://jsonplaceholder.typicode.com/posts/3":
			default:
				urlChannel03 <- Result{URL: localUrl, Err: err, StatusCode: req.StatusCode}
			}

		}(url)

	}

	/**05º como temos vários canais, vamos ler todos eles dentro de um For infinito */
	for {
		select {
		case res := <-urlChannel01:
			if res.Err != nil {
				fmt.Printf("Erro será tratado")
				panic(res.Err)
			}
			fmt.Printf("Resposta para %s: %d\n", res.URL, res.StatusCode)
		case res := <-urlChannel02:
			if res.Err != nil {
				fmt.Printf("Erro será tratado")
				panic(res.Err)
			}
			fmt.Printf("Resposta para %s: %d\n", res.URL, res.StatusCode)
		case res := <-urlChannel03:
			if res.Err != nil {
				fmt.Printf("Erro será tratado no X")
				panic(res.Err)
			}
			fmt.Printf("Resposta para %s: %d\n", res.URL, res.StatusCode)

			wg.Wait()
		}

	}

}
