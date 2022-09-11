// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
// - Solução: https://play.golang.org/p/GLK11Q1_x8y

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

type pessoa struct {
	Nome                    string
	Sobrenome               string
	SaboresFavoritosSorvete []string
}

func main() {

	pessoa1 := pessoa{
		"Alfredo",
		"Silva",
		[]string{
			"Morango",
			"Chocolate",
		},
	}
	pessoa2 := pessoa{
		"Robson",
		"Carvalho",
		[]string{
			"Flocos",
			"Creme",
		},
	}

	pessoas := map[string]pessoa{
		pessoa1.Sobrenome: pessoa1,
		pessoa2.Sobrenome: pessoa2,
	}

	for _,v := range pessoas{
		fmt.Println("O nome da pessoa é", v.Nome, "e seus sorvetes favoritos são:")

		for _,sabores := range v.SaboresFavoritosSorvete{
			fmt.Println("-",sabores)
		}
		fmt.Println("")
	}
}
