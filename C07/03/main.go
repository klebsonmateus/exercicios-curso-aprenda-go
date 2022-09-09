// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
// - Solução: https://play.golang.org/p/qnFjiDJzLor

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes
package main

import (
	"fmt"
)

var anoQueNasci = 1989
var anoAtual = 2022

func main() {
	for anoQueNasci <= anoAtual {
		fmt.Println(anoQueNasci)
		anoQueNasci++
	}
}
