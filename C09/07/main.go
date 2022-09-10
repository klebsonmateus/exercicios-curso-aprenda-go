// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
// - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
// - Solução: https://play.golang.org/p/Gh81-d5tMi

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	ss := [][]string {
		[]string {
			"Leticia",
			"Loiola",
			"Ver vídeo no Youtube",
		},
		[]string {
			"Miguel",
			"Mateus",
			"Passear",
		},
		[]string {
			"Klebson",
			"Bernado",
			"Assistir futebol",
		},
	}

	for i,v := range ss{
		fmt.Println(i)
		for _,k :=range v{
			fmt.Println("\t",k)
		}
	}
}