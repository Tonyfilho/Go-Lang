package main

import (
	"fmt"
)

const x = 10 // so terá tipagem no momento da compilação, se for usado em uma operação com float64, ele será convertido para float64.
var y = 20   // terá tipagem no momento da execução, se for usado em uma operação com float64, ele será convertido para float64.

func main() {
	fmt.Println("Hello, World!")

	a := "e"
	b := "é"
	c := "香"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v", d, e, f)

	s := "ascii éøâ 香"
	/**
	_ que dizer NADA
	For de 3 clausulas
	o Range vai percorrer a string ao invez de retornar bits,
	ele vai retornar o valor do tipo rune, caracteres unicode.
	*/

	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	///outro tipo de FOR que retorna por Bytes, ou seja, retorna o valor do tipo byte, que é um alias para uint8.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

	/// 02 constantes
	/*
	 Constantes são valores que não mudam durante a execução do programa.
	 o Tipo fica indefiniado, mas ele é definido no momento da compilação, e não no momento da execução.

	**/

	///podemos declara assim tb
	const (
		a1 = 10
		b1 = 20
		c1 = 30
	)

	/// 03 Iota
	fmt.Println("03 Iota Numa declaração de constantes, o identificador iota representa números sequenciais, começando com zero. A cada linha de declaração de constante, o valor de iota é incrementado em 1.")
	const (
		a10 = iota + 10000000 //repetindo o iote e somando 10000000, o valor de a10 será 10000000
		_
		c10
		x10
		_
		z10
	)
	fmt.Println(a10, c10, x10, z10)

	//// 04 Deslocamento de Bytes
	fmt.Println("04 Deslocamento de Bytes.")
	/*
		Os Operadores de deslocamento de bits (<< e >>) podem ser usados para multiplicar ou dividir números
		inteiros por potências de dois. Por exemplo, deslocar um número inteiro para a esquerda por 1 bit é equivalente a multiplicá-lo por 2, enquanto deslocá-lo para a direita por 1 bit é equivalente a dividi-lo por 2.
	**/
	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10) ou seja 1 * 2^10 = 1024
		MB = 1 << (iota * 10) // 1 << (2 * 10) ou seja 1 * 2^20 = 1048576
		GB = 1 << (iota * 10) // 1 << (3 * 10) ou seja 1 * 2^30 = 1073741824
		TB = 1 << (iota * 10) // 1 << (4 * 10) ou seja 1 * 2^40 = 1099511627776
	)
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

}
