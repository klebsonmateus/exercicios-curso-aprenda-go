// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
// - Nome
// - Sobrenome
// - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
// - Solução: https://play.golang.org/p/Pyp7vmTJfY

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

type pessoa struct{
	Nome string
	Sobrenome string
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

	fmt.Println(pessoa1.Nome, pessoa1.Sobrenome)
	for _,v := range pessoa1.SaboresFavoritosSorvete{
		fmt.Println("\t",v)
	}

	fmt.Println(pessoa2.Nome, pessoa2.Sobrenome)
	for _,v := range pessoa2.SaboresFavoritosSorvete{
		fmt.Println("\t",v)
	}
}