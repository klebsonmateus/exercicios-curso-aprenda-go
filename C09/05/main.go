// - Comece com essa slice:
//     - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// - Utilizando slicing e append, crie uma slice y que contenha os valores:
//     - [42, 43, 44, 48, 49, 50, 51]
// - Solução: https://play.golang.org/p/26bT-UKmJH

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[len(x)-4:]...)
	fmt.Println(y)
}
