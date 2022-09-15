// - Partindo do código abaixo, utilize marshal para transformar []user em JSON.
// - https://play.golang.org/p/U0jea43X55
// - Atenção! Tem pegadinha aqui.
// - Solução: https://play.golang.org/p/gUOHXruFGs

// → Compartilhe sua solução e converse com outros estudantes em: https://github.com/ellenkorbes/aprend...

// → Segue lá: http://twitter.com/veekorbes

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(usersJson))
}
