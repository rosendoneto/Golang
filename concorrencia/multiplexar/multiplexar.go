package main

import (
	"fmt"
	"html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - Misturar mensagens
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main()  {
	var zc = juntar(
		html.Titulo("Https://www.mercadolivre.com.br", "Https://www.mercadopago.com.br")
		html.Titulo("Https//www.google.com", "Https://www.youtube.com.br")
		)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}