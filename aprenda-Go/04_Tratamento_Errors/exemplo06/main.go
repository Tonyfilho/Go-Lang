package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)





func main() {
	fmt.Println("Exemplo 06 - Panic")
	fmt.Println("http://godoc.org/builtin#panic")

	f, err := os.Create("names.txt")
	if err != nil {
		/**Tratamento do erro onde colocamos o erro na tela e retornamos */
		fmt.Println("Error: ", err)
		return
	}

	defer f.Close()
	defer fmt.Println("Arquivo criado com sucesso!")
	r := strings.NewReader("Guilherme\nRafael\nLucas\n")
	io.Copy(f, r)



}