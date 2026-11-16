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
	"sync"
	"time"

	"github.com/nfnt/resize"
)

/*
*NEste exemplo 1 imagem não tem relação uma com a outra, dentro do for
processThumbnail(filepath.Join(inputDir, file.Name()), outputDir)
*/
func main() {
	fmt.Println("Aula 13 processando imagens com GoRoutines")
	fmt.Println()

	inputDir := "imgs"
	outputDir := "thumbnails"
	now := time.Now() 
	processImages(inputDir, outputDir)
	fmt.Println(time.Since(now).Milliseconds()) // saida de 948

}

func processImages(inputDir, outputDir string) {
	/**Como sabemos a quantidade de arquivos usaremos esta abordagem**/
	/**1º É criar um wg e passar o tamanho do slice de arquivos, mas temos q por depois da verificação de pasta*/
	files, err := os.ReadDir(inputDir)
	var wg sync.WaitGroup

	if err != nil {
		fmt.Errorf("Erro ao ler o diretório de entrada: %v", err)
		log.Panic(err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Errorf("Erro ao Criar diretório de saída OutputDir: %v", err)
		log.Panic(err)
	}

	for _, file := range files {

		/**2º Aqui toda a vez q entrar no if será feito um sync waitgroup , ou seja para cada item NO FOR*/
		if !file.IsDir() {
			wg.Add(1)
			/**2b Aqui temos que usar a GoRoutine*/
			go processThumbnail(filepath.Join(inputDir, file.Name()), outputDir, &wg)
		}
	}
	/**3º Aqui esperaremos por todos os goroutines terminarem*/
	wg.Wait()

	fmt.Println("Processamento de imagens concluído!")
}

/**4º Temos que receber como paramentro do método wg.waitGroup como ponteiro*/

func processThumbnail(fileInput string, outputDir string, wg *sync.WaitGroup) {
	/**5º Aqui temos que chamar o wg.Done() no final do método*/
	/**Apesar de esta no incio com o defer() ele será chamado quando a função retornar ou seja no final*/
	defer wg.Done()

	file, err := os.Open(fileInput)
	if err != nil {
		fmt.Errorf("Erro ao abrir o arquivo dentro do método processThumbnail: %v", err)
		log.Panic(err)
	}

	defer file.Close()
	imageDecode, _, err := image.Decode(file)
	if err != nil {
		fmt.Errorf("Erro ao fazer Decode: %v", err)
		log.Panic(err)
	}
	//thumbnail := resize.Resize(100, 100, imageDecode, resize.Lanczos3) //redimensionando a imagem para 100x100 pixels}
	thumbnail2 := resize.Thumbnail(100, 100, imageDecode, resize.Lanczos3)

	outputPath := filepath.Join(outputDir, filepath.Base(fileInput))

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Errorf("Erro ao criar o arquivo de saída no processThumbnail: %v", err)
		log.Panic(err)
	}

	defer outputFile.Close()

	switch strings.ToLower(filepath.Ext(fileInput)) {
	case ".png":
		err = png.Encode(outputFile, thumbnail2)
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
