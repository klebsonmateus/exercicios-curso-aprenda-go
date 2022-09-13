// - Crie um valor e atribua-o a uma variável.
// - Demonstre o endereço deste valor na memória.
// - Solução: https://play.golang.org/p/0jVt1yaoFL

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes
package main

import "fmt"

func main() {
	x := 10

	fmt.Println(&x)
}