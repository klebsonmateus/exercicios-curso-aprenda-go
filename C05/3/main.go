// - Crie constantes tipadas e não-tipadas.
// - Demonstre seus valores.
// - Solução: https://play.golang.org/p/eWnKI59ual

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

const a = 10
const b int = 20
const c = "Hello"
const d = false

func main() {
	fmt.Printf("%v-%T\t%v-%T\t%v-%T\t%v-%T\t\n", a,a,b,b,c,c,d,d)
}
