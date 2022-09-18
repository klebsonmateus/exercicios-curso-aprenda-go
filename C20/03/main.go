// - Utilizando goroutines, crie um programa incrementador:
// - Tenha uma variável com o valor da contagem
// - Crie um monte de goroutines, onde cada uma deve:
// 	- Ler o valor do contador
// 	- Salvar este valor
// 	- Fazer yield da thread com runtime.Gosched()
// 	- Incrementar o valor salvo
// 	- Copiar o novo valor para a variável inicial
// - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
// - Demonstre que há uma condição de corrida utilizando a flag -race
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

func main() {
	fmt.Println("O valor inicial do contador é", contador)
	wg.Add(totalGoRoutines)

	for x := 1; x <= totalGoRoutines; x++ {
		go incrementaContador(x)
	}
	
	wg.Wait()

}

func incrementaContador(i int) {
		cont := contador
		runtime.Gosched()
		cont++
		contador = cont
		fmt.Println("Essa é a routine",i,"e o valor atual do contador é", contador)
		wg.Done()

}
