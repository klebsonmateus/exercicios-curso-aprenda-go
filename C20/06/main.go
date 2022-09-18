// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//     - go run
//     - go build
//     - go install
// - Solução: https://github.com/ellenkorbes/aprend...

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("O seu OS é:", runtime.GOOS)
	fmt.Println("O seu ARCH é:", runtime.GOARCH)
}
