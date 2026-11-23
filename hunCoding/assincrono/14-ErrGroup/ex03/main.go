package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

/*temos q fazer uma importação da biblioteca "golang.org/x/sync/errgroup"**/
func main() {
	fmt.Println("Aula 14 Gerenciando erros dentre várias funções assíncronas em GoLang")
	fmt.Println("Neste exemplo usaremos ErrGroup para gerenciar erros em funções assíncronas em GoLang")
	var g errgroup.Group
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://x.comnn", // URL inválida para simular erro
	}

	/**Não precisaremos de waitgroup usaremos a variavel G*/
	/**1º temos que usar a variavel G e a função g.Go()*/
	/**2º Esta função tem q retornar um erro*/
	/**OBS: Não precisamos criar GoRoutines Anonimas, ela já é criada internamente*/

	for _, url := range urls {
		localUrl := url // Captura a variável local para evitar problemas de concorrência
		g.Go(func() error {
			req, err := http.Get(localUrl)
			if err != nil {
				fmt.Printf("Erro ao fazer requisição Get para %s: %v\n", localUrl, err)
				/*3º Caso tenhamos erro basta por um return err**/
				/*pois é justamente é o que precisamos pegar o error**/
				return err
			}
			fmt.Printf("requisição Get bem-sucedida para %s: %v\n", localUrl, req.StatusCode)
			defer req.Body.Close()
			/**4º Caso não tenhamos erro, retornamos nil**/
			return nil
		})

	}

	/**5º Agora precisamos pegar o erro de todas as funções assíncronas**/
	/**Usaremos g.Wait() para fazer o trabalho de waitGroup**/
	err := g.Wait()
	if err != nil {
		fmt.Printf("Erro encontrado: %v\n", err)
		return
	}
		fmt.Println("Todas as requisições foram concluídas com sucesso.")
	
   /**6º Ao inves de ficar usando Channels, usaremos ErrGroup*/

}
