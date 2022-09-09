// - Crie um programa que:
// - Atribua um valor int a uma variável
// - Demonstre este valor em decimal, binário e hexadecimal
// - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// - Demonstre esta outra variável em decimal, binário e hexadecimal
// - Solução: https://play.golang.org/p/IiwgT0v3Mp

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

var x int = 200

func main() {
	fmt.Printf("%d\t%b\t%#x\t\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\t\n", y, y, y)

}
