// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"
// - Solução: https://play.golang.org/p/qLY-q3vffQ

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q := quadrado{lado: 3}
	c := circulo{raio: 2}

	fmt.Println("O quadrado tem lado", q.lado, "e área", info(q))
	fmt.Println("O círculo tem raio", c.raio, "e área", info(c))
}
