// - Crie e utilize uma função anônima.
// - Solução: https://play.golang.org/p/Kgo6hVr5G5

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	func(x, y int) {
		res := x + y
		fmt.Println(res)
	}(3,4)

	
}
