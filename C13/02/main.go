// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.
// - Solução: https://play.golang.org/p/Tgv3wwxKV-

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 82}
	si2 := []int{1, 22, 3, 4, 76, 82}


	fmt.Println(soma(si...))
	fmt.Println(somaSlice(si2))
}

func soma(x...int)int{
	total:=0
	for _,value:= range x{
		total+=value
	}
	return total
}

func somaSlice(x []int)int{
	total:=0
	for _,value:= range x{
		total+=value
	}
	return total
}