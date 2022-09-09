// - Utilizando a solução anterior, adicione as opções else if e else.
// - Solução: https://play.golang.org/p/_cO7E-yL0o 

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x:= 2

	if x >5 {
		fmt.Println("X é maior que cinco")
	} else if x <5 {
		fmt.Println("x é menor que cinco")
	} else {
		fmt.Println("x = 5")
	}
}
