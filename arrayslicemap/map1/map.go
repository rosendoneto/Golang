package main

import "fmt"

func main() {
	// var aprovados map[int] string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456] = "Maria"
	aprovados[456789] = "Pedro"
	aprovados[789345] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[456789])
	delete(aprovados, 456789)
	fmt.Println(aprovados[456789])
}
