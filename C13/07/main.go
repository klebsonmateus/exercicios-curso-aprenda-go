// - Atribua uma função a uma variável.
// - Chame essa função.
// - Solução: https://play.golang.org/p/RMHLL3N5Ww

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Hello World!")
	}

	x()
}
