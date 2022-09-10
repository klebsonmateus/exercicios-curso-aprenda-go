// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
// - Solução: https://play.golang.org/p/3fcvHlt8Lm

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes


package main

import "fmt"

func main() {

	map1 := map[string][]string{
		"mateus_klebson": {"futebol", "flamengo"},
		"loiola_miguel":  {"youtube", "carrinho"},
		"loiola_leticia": {"youtube", "escola"},
	}

	map1["mateus_gabriela"] = []string {"Fazer bolo", "ver novela"}


	for key, hobbies := range map1 {
		fmt.Println(key)
		for _, hobby := range hobbies {
			fmt.Println("\t", hobby)
		}
	}
}