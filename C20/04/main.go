// - Utilize mutex para consertar a condição de corrida do exercício anterior.
// - Solução: https://github.com/ellenkorbes/aprend...

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes


package main

import (
	"fmt"
	"runtime"
	"sync"
)

var contador int = 0
var totalGoRoutines int = 1000
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("O valor inicial do contador é", contador)
	wg.Add(totalGoRoutines)

	for x := 1; x <= totalGoRoutines; x++ {
		go incrementaContador(x)
	}
	
	wg.Wait()

}

func incrementaContador(i int) {
		mu.Lock()
		defer mu.Unlock()
		cont := contador
		runtime.Gosched()
		cont++
		contador = cont
		fmt.Println("Essa é a routine",i,"e o valor atual do contador é", contador)
		wg.Done()

}
