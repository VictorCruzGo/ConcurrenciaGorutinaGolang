package main

import (
	"fmt"
	"perrito"
)

type canino struct {
	nombre string
	edad   int32
}

//La funcion main deberia estar fuera del paquete perrito
func main() {
	c1 := canino{
		nombre: "Chester",
		edad:   perrito.Edad(2),
	}

	fmt.Println(c1)
}
