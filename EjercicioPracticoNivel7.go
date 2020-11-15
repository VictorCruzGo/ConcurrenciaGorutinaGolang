package main

import "fmt"

//-----------Ejercicio 1
/*
func main() {
	x := 32
	fmt.Println(x)
	fmt.Println(&x)
}
*/

//-----------Ejercicio 2
type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p *persona) cambiame() {
	(*p).nombre = "Juan"
	(*p).apellido = "Ali"
	(*p).edad = 50
}

func cambiame(p *persona) {
	(*p).nombre = "Juan Carlos"
}

func main() {
	p := persona{
		nombre:   "Victor",
		apellido: "Cruz",
		edad:     30,
	}

	fmt.Println(p)
	fmt.Printf("%T\n", p)

	//p.cambiame()
	(&p).cambiame()

	fmt.Println(p)
	fmt.Printf("%T\n", p)

	cambiame(&p)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
}
