// - Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//     - Crie um tipo para um struct chamado "pessoa"
//     - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//     - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//     - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//     - Demonstre no seu código:
//         - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//         - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
// - Se precisar de dicas, veja: https://gobyexample.com/interfaces
// - Solução: https://github.com/ellenkorbes/aprend...

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import "fmt"

type humanos interface {
	falar()
}

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p *pessoa) falar() {
	fmt.Println("Olá, meu nome é", p.Nome, p.Sobrenome, "e eu posso falar!")

}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	joao := pessoa{
		Nome: "João",
		Sobrenome: "Silva",
		Idade: 32,
	}

	dizerAlgumaCoisa(&joao)
}
