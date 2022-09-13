// - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
// - Solução: https://play.golang.org/p/b5Ua2kNN8a

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")

}