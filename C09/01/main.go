// - Usando uma literal composta:
// - Crie um array que suporte 5 valores to tipo int
// - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.
// - Solução: https://play.golang.org/p/tpWDzzlO2l

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	array := [5]int {1,2,3,4,5}
	for i,v := range array {
		fmt.Println(i,"-",v)
	}
	fmt.Printf("%T",array)
}