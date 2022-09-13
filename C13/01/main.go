// - Exercício:
// - Crie uma função que retorne um int
// - Crie outra função que retorne um int e uma string
// - Chame as duas funções
// - Demonstre seus resultados

package main

import "fmt"

func main() {
	fmt.Println(retorna1())
	fmt.Println(retorna2())
}

func retorna1() int{
	return 5
}

func retorna2()(int, string){
	

	return 10, "Olá"
}