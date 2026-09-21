package main

import (
	"fmt"
)

func main() {

	/**01 Operadores Lógicos*/

	fmt.Println("Aula 06 - Operadores Lógicos")
	x := 9

	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println("é múltiplo de dois e tambem de treis")
	}

	if x%2 == 0 || x%3 == 0 {
		fmt.Println("é múltiplo de dois ou de treis")
	}

	/* 02 Agrupamentos de dados*/
	fmt.Println("\n 02 Agrupamentos de dados")
	fmt.Println("\n Agrupamentos de dados - Arrays")

	var xArray [5]int
	var yArray [6]int

	xArray[0] = 1
	xArray[1] = 10
	fmt.Println(xArray[0], xArray[1])
	fmt.Println(xArray)
	fmt.Printf("%T\n", xArray)
	fmt.Printf("%T\n", yArray)
	fmt.Println(len(xArray))

	/* 03 Agrupamentos de dados*/
	fmt.Println("\n Agrupamentos de dados - Slices")

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	/**NO SLICE podemos deixar o tamanho dinâmico*/
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
	/**Error Criado*/
	// slice[20] = 1
	// fmt.Println(slice[20])

	/* 03A Agrupamentos de dados*/
	fmt.Println("\n Agrupamentos de dados - Slices, fatiando ou deletando de uma fatia ")
	/**Vai fazer o  Slice entre o indice 2 e 4*/
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia1 := sabores[2:4] //[abacaxi quatroqueijos]
	fatia2 := sabores[2]   //abacaxi
	fatia3 := sabores[:4]  //[pepperoni mozzarela abacaxi quatroqueijos]

	fmt.Println("fatia1: ", fatia1)
	fmt.Println("fatia2: ", fatia2)
	fmt.Println("fatia3: ", fatia3)

	fmt.Println("\n Agrupamentos de dados - Slices, fatiando ou deletando de uma fatia ")

	//		    0.           1.           2.         3.               4.
	sabores2 := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	sabores2 = append(sabores2[:2], sabores2[4:]...)

	fmt.Println("sabores2: ", sabores2)

	fmt.Println("Outro exemplo de fatiamento")

	//			    0.           1.           2.         3.               4.
	sabores3 := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia := sabores3[:]

	fmt.Println(fatia)

	fmt.Println()
	fmt.Println("Por Index", sabores3[0], sabores3[1], sabores3[2], "\n")

	for i := 0; i < len(sabores3); i++ {

		fmt.Println("Item", i, ":", sabores3[i])
	}

	/**03C Slice e a Função Append()*/
	fmt.Println()
	fmt.Println("03C Slice e a Função Append()")

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)
	/**Lembrando que o operador ... é usado para desestruturar slices é o mesmo que fazer 5,6,7,8, ...*/
	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	/**Obs: ao operador ... tem a mesma função que o operador ... destruturante no typescript*/
	/**Desta forma passamos os ITENS da slice outraslice*/
	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)

	/**03D Slice e a Função MAKE()*/
	fmt.Println()
	fmt.Println("03D Slice e a Função MAKE()")

	sliceMake := make([]int, 5, 10)
	/*Tribuindo valores ao slice em linha*/
	sliceMake[0], sliceMake[1], sliceMake[2], sliceMake[3], sliceMake[4] = 1, 2, 3, 4, 5

	/* de acordo com a capacidade do slice,
	o append vai aumentar a capacidade do slice para 10,
	pois o slice foi criado com capacidade de 10.*/

	sliceMake = append(sliceMake, 6)
	/* o append serve para aumentar a capacidade do slice,
	   caso o slice seja criado com capacidade
	    menor que a quantidade de elementos que serão adicionados.*/

	sliceMake = append(sliceMake, 7)
	sliceMake = append(sliceMake, 8)
	sliceMake = append(sliceMake, 9)
	sliceMake = append(sliceMake, 10)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake)) //[1 2 3 4 5 6 7 8 9 10] 10 10

	sliceMake = append(sliceMake, 10)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake)) //[1 2 3 4 5 6 7 8 9 10 10] 11 20

	/**03E Slice e a Multidimensional ou Matrizes*/
	fmt.Println()
	fmt.Println("03E Slice Multidimensional ou Matrizes")

	ss2 := [][]int{
		// Índice: 0  1  2                   // Índice:
		[]int{1, 2, 3, 4, 5, 6},       // 0
		[]int{7, 8, 9, 10, 11, 12},    // 1
		[]int{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println("Matriz [2][4]:", ss2[2][4])

	ss4 := [][][][]int{

		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				[]int{10, 20, 30, 40, 50},
			},
		},

		[][][]int{
			[][]int{
				[]int{2, 4, 6, 8, 10},
			},
			[][]int{
				[]int{3},
			},
		},
	}
	fmt.Println("Matriz 4D [1][0][0][2]:", ss4[1][0][0][2])

	
	/**03F Slice a surpresa do array subjacente*/
	fmt.Println()
	fmt.Println("03F Slice a surpresa do array subjacente")

	
	primeiroslice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(primeiroslice) //[1 2 3 4 5]
	/**usando append para pegando dados do índice 0 a 2 e do índice 4 ao final*/
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice) //[1 2 5]
	/**O append acima alterou o array subjacente do slice primeiroslice, pois o slice segundoslice e o slice primeiroslice compartilham o mesmo array subjacente.*/

	fmt.Println(primeiroslice) 
	/** Modifica a referência da memoria, saida [1 2 5 4 5] - o valor do índice 2 foi alterado para 5, 
	pois o slice segundoslice compartilha o mesmo array subjacente
	 do slice primeiroslice.*/

}
