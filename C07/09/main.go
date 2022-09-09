// Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
// - Solução: https://play.golang.org/p/4-iTPZwfEz

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	esporteFavorito:="Natação"

	switch esporteFavorito{
	case "Futebol":
		fmt.Println("Mandou bem!")
	case "Natação":
		fmt.Println("Mandou mal!")
	default:
		fmt.Println("Vá se exercitar!")
	}
}