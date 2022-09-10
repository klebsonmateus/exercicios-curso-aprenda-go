// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.
// - Solução: https://play.golang.org/p/nD3TW8VQmH

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

	fmt.Println(map1)
	for key, hobbies := range map1 {
		fmt.Println(key)
		for _, hobby := range hobbies {
			fmt.Println("\t", hobby)
		}
	}
}
