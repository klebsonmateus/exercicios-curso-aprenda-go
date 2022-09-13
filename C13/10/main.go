// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
// - Solução: https://play.golang.org/p/sA7NHpkCCg

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x := exClosure()
	y := exClosure()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}

func exClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
