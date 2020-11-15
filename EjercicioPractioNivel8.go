package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	fmt.Println("")
	u1 := usuario{
		Nombre: "Eduar",
		Edad:   32,
	}
	u2 := usuario{
		Nombre: "Victor",
		Edad:   35,
	}
	u3 := usuario{
		Nombre: "Maria",
		Edad:   40,
	}

	usuarios := []usuario{u1, u2, u3}

	fmt.Println(usuarios)

	//
	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println(err)
	}

	//Los atributos del struct tienen que iniciar con Mayuscula
	fmt.Println(string(bs))
}
