// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
// - Solução: https://play.golang.org/p/X7qm3aWSa6

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}
