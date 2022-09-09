// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
// - Solução: https://play.golang.org/p/TyIGp4Hi8B

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x := 20

	switch {
	case x > 5:
		fmt.Println("X é maior que cinco")

	case x < 5:
		fmt.Println("x é menor que cinco")

	case x == 5:
		fmt.Println("x = 5")
	}
}
