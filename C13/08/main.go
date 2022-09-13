// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
// - Solução: https://play.golang.org/p/A74rufv6Rs

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x:= func () func() {
		return func ()  {
			fmt.Println("Função secundária")
		}
	}

	x()()
}