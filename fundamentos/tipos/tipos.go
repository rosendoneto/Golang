package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Liberal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (Só positivos)... int8 int16 int32 int64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (float32 float64)
	var x float32 = 49.99
	fmt.Println("O tipo de X é ", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é ", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))

	// string
	s1 := "Olá meu nome é Neto"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String com múltiplas linhas
	s2 := `Olá 
	meu 
	nome 
	é 
	Neto`
	fmt.Println("O tamanho da string é", len(s2))

	// char???
	// var x char = 'b'
	char := 'a'
	fmt.Println("O tipo de bo é ", reflect.TypeOf(char))
	fmt.Println(char)

}
