package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

func main() {
	fmt.Println("Aula 13 processando imagens com GoRoutines")
	fmt.Println()
	/**1ª coisa é criar as variaveis onde recebem o caminho das imagens*/
	inputDir := "imgs"        //caminho da pasta onde estão as imagens
	outputDir := "thumbnails" //caminho da pasta onde serão salvas as imagens processadas


    /**Veremos quanto tempo leva para processar as imagens*/
	now := time.Now() //pegando o tempo atual
	processImages(inputDir, outputDir) //chamando o método processImages, passando o caminho das imagens e o caminho da pasta onde serão salvas as imagens processadas
   // fmt.Println(time.Now().Sub(now).Microseconds())
    fmt.Println(time.Since(now).Milliseconds()) //mostrando o tempo que levou para processar as imagens em milissegundos total de : 3825
}

/**2º Criaremos um metodo de processamento de imagens*/
/**Não precisa retornar nada, pois procesaremos e vamos por para outra pasta*/

func processImages(inputDir, outputDir string) {
	/**3º temos que pegar os files ou os erros*/
	files, err := os.ReadDir(inputDir) //pegando os arquivos da pasta de entrada
	if err != nil {
		fmt.Errorf("Erro ao ler o diretório de entrada: %v", err)
		log.Panic(err) //se der erro, o programa para
	}
	/**4º Apesar de passar pelo erro, podemos ter uma erro de não existência da pasta*/
	/**Então caso a pasta não exista, vamos criá-la, passaremos o nome da pasta e permissão q sistema vai dar*/
	err = os.MkdirAll(outputDir, os.ModePerm) //criando a pasta de saída caso não exista
	if err != nil {
		fmt.Errorf("Erro ao Criar diretório de saída OutputDir: %v", err)
		log.Panic(err) //se der erro, o programa para
	}

	/*5º pegaremos as imagens que estarão no array de files**/
	for _, file := range files {
		/**Chamaremos o metodo de processamento de imagens*/
		/**Antes temos que checar se o arquivo não é um diretório*/
		/**Caso haja outra pasta, teriamos q ter outro For*/
		if !file.IsDir() {
			/*neste caso tenho que juntar o nome do arquivo com o caminho da pasta de saída,
			ara que o programa reconheça o local onde deve salvar a imagem processada*/
			processThumbnail(filepath.Join(inputDir, file.Name()), outputDir) //passando o caminho do arquivo e a pasta entrada e de saída
		}
	}

	/*Colocaremos uma mensagem para indicar que o processamento foi concluído*/
	fmt.Println("Processamento de imagens concluído!")
}

/**6ª Metodo process Thumbnail, onde receberá um File e a pasta de saída*/
func processThumbnail(fileInput string, outputDir string) {
	/**Temos que abrir o arquivo*/
	file, err := os.Open(fileInput) //abrindo o arquivo e passando o file como paramétro 
	if err != nil {
		fmt.Errorf("Erro ao abrir o arquivo dentro do método processThumbnail: %v", err)
		log.Panic(err) //se der erro, o programa para
	}
	/**Não tendo erro, daremos um defer para fechar o arquivo no final*/
	defer file.Close()
	/**Agora podemos processar a imagem*/
	/**temos 3 retornos image.Image, string, error*/
	imageDecode, _, err := image.Decode(file) //decodificando a imagem, ou seja, transformando o arquivo em uma imagem
	if err != nil {
		fmt.Errorf("Erro ao fazer Decode: %v", err)
		log.Panic(err) //se der erro, o programa para
	}
	
//	fmt.Printf("Images: %v", imageDecode.Bounds().Size()) //mostrando o tamanho da imagem original p debug


	/**Agora q temos modificar o tamanho da imagem*/
	/**100px de max-width, 100px de max-height , temos que passar a imagem e passar um meto q existem no pacote resize*/
	//thumbnail := resize.Resize(100, 100, imageDecode, resize.Lanczos3) //redimensionando a imagem para 100x100 pixels}
	thumbnail2 := resize.Thumbnail(100, 100, imageDecode, resize.Lanczos3) //redimensionando a imagem para 100x100 pixels mantendo a proporção da imagem original
	//fmt.Printf("Thumbnail: %v", thumbnail.Bounds().Size()) //mostrando o tamanho da imagem processada p debug
	//fmt.Printf("Thumbnail2: %v", thumbnail2.Bounds().Size()) //mostrando o tamanho da imagem processada p debug

	/*7º Temos q por para onde a imagem processada deve ser salva*/
	/*Temos que pegar o outputDir e o nome do arquivo q sera gerado*/
	outputPath := filepath.Join(outputDir, filepath.Base(fileInput)) //juntando o caminho da pasta de saída com o nome do arquivo

	/*8º Depois de termos o nome do arquivo e caminho de saida, criaremos o arquivo de saída*/
	/*Ou seja cria este arquivo q vou escrever dentro dele*/
	outputFile, err := os.Create(outputPath) //criando o arquivo de saída
	if err != nil {
		fmt.Errorf("Erro ao criar o arquivo de saída no processThumbnail: %v", err)
		log.Panic(err) //se der erro, o programa para
	}
	/*Colocaremos um defer para fechar o arquivo de saída no final*/
	defer outputFile.Close()

	/**agora temmos que dar um Switch para salvar a imagem processada* se é PNG, JPEG, etc.*/
	/**Nos temos 2 pacotes: png e jpeg*/
	/**1º temos q formar o q vamos receber*/
	switch strings.ToLower(filepath.Ext(fileInput)) { //passando o formato da imagem para minúsculo
	/**iremos descobrir o formato da imagem original*/
	case ".png":
		err = png.Encode(outputFile, thumbnail2) //salvando a imagem processada em PNG
		if err != nil {
			fmt.Errorf("Erro ao salvar a imagem processada em PNG: %v", err)
			log.Panic(err)
		}
	case ".jpeg", ".jpg":
		err = jpeg.Encode(outputFile, thumbnail2, nil)
		if err != nil {
			fmt.Errorf("Erro ao salvar a imagem processada em JPEG: %v", err)
			log.Panic(err)
		}
	}

}
