// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.
// - Solução: https://play.golang.org/p/GBZcnu0Ajp

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) demonstraNomeEIdade(){
	fmt.Println("O nome completo da pessoa é:", p.nome, p.sobrenome)
	fmt.Println("Esta pessoa tem",p.idade,"anos.")
}

func main() {

	pessoa1 := pessoa{
		nome:      "Mario",
		sobrenome: "Alberto",
		idade:     33,
	}

	pessoa1.demonstraNomeEIdade()
}


