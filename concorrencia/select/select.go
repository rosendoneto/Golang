package main

import (
	"html"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// Estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	default:
		return "Sem resposta ainda!"

	}
}

func main() {
	campeao := oMaisRapido(
		"Https://www.mercadolivre.com.br",
		"Https://www.mercadopago.com.br",
		"Https://www.google.com",
	)
	println(campeao)
}
