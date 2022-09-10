// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
// - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
// - Do quinto ao último item do slice (incluindo o último item!)
// - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
// - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
// - Solução: https://play.golang.org/p/1aPXVeR1mf 

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n",slice)

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2:len(slice)-1])
}