// - Callback: passe uma função como argumento a outra função.
// - Solução: https://play.golang.org/p/2epLD_Yyap

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes
package main

import "fmt"

func main() {
	recebeArgumento(argumento)
}

func argumento(){
	fmt.Println("Olha eu aqui!")
}

func recebeArgumento(f func()){
	fmt.Println("Olá é o seguinte:")
	f()
}