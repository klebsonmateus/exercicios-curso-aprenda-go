// - Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
// - ==
// - !=
// - <=
// - <
// - >=
// - >
// - Demonstre estes valores.
// - Solução: https://play.golang.org/p/BMYEch6_s8

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes
package main

import "fmt"

func main() {
	a := 10==10
	b := 10!=10
	c := 10<=10
	d := 10<10
	e := 10>=10
	f := 10>10

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n",a,b,c,d,e,f)
}