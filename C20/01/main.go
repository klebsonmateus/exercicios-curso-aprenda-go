// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
// - Solução:
//     - https://github.com/ellenkorbes/aprend...
//     - https://github.com/ellenkorbes/aprend...

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Eu sou a primeira func!")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Eu sou a segunda func!")
	}()

	wg.Wait()
}

