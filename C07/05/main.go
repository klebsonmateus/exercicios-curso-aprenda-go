// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
// - Solução: https://play.golang.org/p/zcEsXqnBr8

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		res := i%4
		fmt.Printf("%v dividido por 4 sobra : %v\n", i, res)
	}
}
