package main

import (
	"fmt"
	"sort"
)

/*Criaremos 1º struct do tipo carro*/
type Carro struct {
	Nome     string
	Potencia int
	Consumo  int
}

/**2º Temos que criar um TIPO que é um Slice de um Struct, ordenarPorPotencia []Carro ordenarPorConsumo**/
type ordenarPorPotencia []Carro
type ordenarPorConsumo []Carro
type ordenarPorLucro []Carro

/**3º Função Len(),Less() e Swap() onde retornaremos para cada TIPO ordenarPorPotencia*/
/**Fazendo da 1º Função type ordenarPorPotencia []Carro */
/** Função tamanho do Slice*/
func (x ordenarPorPotencia) Len() int { return len(x) }

/**Função Less(i, j int) onde temos 2 paramentros e queremos saber se 2º é maior que o 1º e retorna booleano*/
/*temos que compara os valores dentro dos indexes, quem o valor Menor*/
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].Potencia < x[j].Potencia }

/**Função Swap(i, j int) e muda a ordem dos elementos, sem retornar*/
func (x ordenarPorPotencia) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

/**Fazendo da 2º Função type ordenarPorConsumo []Carro */
func (x ordenarPorConsumo) Len() int { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool {return x[i].Consumo > x[j].Consumo}
func (x ordenarPorConsumo) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

/**Fazendo da 3º Função type ordenarPorLucro []Carro , nesta, queremos fazer a logica de Maior e não do Menor */
func (x ordenarPorLucro) Len() int { return len(x) }
func (x ordenarPorLucro) Less(i, j int) bool {return x[i].Consumo < x[j].Consumo}
func (x ordenarPorLucro) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func main() {
	fmt.Println("Aula17 Costumizando nosso proprio Sort")
	/**1ª Temos que  Criar uma Função Sort(data Interface)*/
	fmt.Println()
	fmt.Println("Criaremos 1º struct do tipo carro")
	/**Criação do Objeto de Slices */
	carros := []Carro{
		{Nome: "Civic", Potencia: 180, Consumo: 9},
		{Nome: "Corolla", Potencia: 140, Consumo: 10},
		{Nome: "Gol", Potencia: 80, Consumo: 12},
	}
	fmt.Println()
	fmt.Println("Criaremos a Interface interface de Carro")
	fmt.Println("Temos que criar type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }")
	fmt.Println()
	fmt.Println("Interface ordenarPorPotencia []Carro  com estes suas Funções { Len() int; Less(i, j int) bool; Swap(i, j int) }")
	fmt.Println("ordenarPorPotencia implementa automaticamente Sort.Interface, bastando criar os metodos")
	fmt.Println()
	
	fmt.Println("Como Funciona? invoco sort, converto passando o Type criado e na memoria fica modificado")
	fmt.Println("Imprimindo o Objeto original: ", carros)
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println("Imprimindo o Objeto Modificando por Potencia: \n", carros)
	fmt.Println()
	
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println("Imprimindo o Objeto Modificando por Consulmo: \n", carros)
	fmt.Println()
	sort.Sort(ordenarPorLucro(carros))
	fmt.Println("Imprimindo o Objeto Modificando por Lucro para a Bomba de Gasolina: \n", carros)

}
