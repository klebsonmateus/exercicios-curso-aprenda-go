// - Utilize atomic para consertar a condição de corrida do exercício #3.
// - Solução: https://github.com/ellenkorbes/aprend...

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var contador uint64
var totalGoRoutines int = 1000
var wg sync.WaitGroup

func main() {
	fmt.Println("O valor inicial do contador é", contador)
	wg.Add(totalGoRoutines)

	for x := 1; x <= totalGoRoutines; x++ {
		go incrementaContador(x)
	}

	wg.Wait()

}

func incrementaContador(i int) {
	atomic.AddUint64(&contador, 1)
	fmt.Println("Essa é a routine", i, "e o valor atual do contador é", contador)
	wg.Done()

}
//Aparentemente a condição de corrida não foi eliminada, checar posteriormente o motivo