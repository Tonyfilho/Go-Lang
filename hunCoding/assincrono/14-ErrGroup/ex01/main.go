package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Aula 14 Gerenciando erros dentre várias funções assíncronas em GoLang")
	/**1º temos 3 URL onde vamos fazer requisições assíncronas*/
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}
	/**1B criaremos o waitgroup */
	var wg sync.WaitGroup

	/*2º iremos fazer as requisições de forma assíncrona*/
	/*2B Criaremos um For e teremos acesso as URLs **/
	for _, url := range urls {
		/**Para cada item ou seja ccada GoRoutines daremos um add()*/
		wg.Add(1)
		/*Aqui você pode fazer a requisição HTTP para a URL*/
		go func(localUrl string) {
			/**3º temos q dar um done() quando a requisição for concluída*/
			defer wg.Done()
			/**Aqui você pode fazer a requisição HTTP para a URL*/
			req, err := http.Get(url)
			if err != nil {
				fmt.Printf("Erro ao fazer requisição para %s: %v\n", localUrl, err)
				log.Panic(err)
			}
			/**fechando o corpo da resposta*/
			defer req.Body.Close()
			/**imprimindo somente um status*/
			fmt.Printf("Resposta para %s: %d\n", localUrl, req.StatusCode)
		}(url)

	}
	/**4º temos q esperar todas as GoRoutines terminarem*/
	wg.Wait()
}
