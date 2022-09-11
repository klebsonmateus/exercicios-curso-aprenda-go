// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
// - Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
		map1 map[string]string
		slice1 []string
	}{
		nome:  "Eu",
		idade: 10,
		map1: map[string]string{"chave":"valor"},
		slice1: []string{"string1","string2","string3"},
	}

	fmt.Println(x)
}
