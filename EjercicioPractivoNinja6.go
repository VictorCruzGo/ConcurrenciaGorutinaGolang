package main

import "fmt"

//------------Ejercicio 1
/*
func main() {
	x := foo()
	fmt.Println(x)
	a, b := bar()
	fmt.Println(a, b)
}

func foo() int {
	x := 1
	return x
}

func bar() (int, string) {
	x := 10
	s := "hola"
	return x, s
}
*/

//------------Ejercicio 2
/*
func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(foo(slice...))
	fmt.Println(bar(slice))
}

func foo(x ...int) int {
	suma := 0
	for _, v := range x {
		suma += v
	}

	return suma
}

func bar(xi []int) int {
	suma := 0
	for _, v := range xi {
		suma += v
	}

	return suma
}
*/

//------------Ejercicio 3
/*
func main() {
	fmt.Println("hoa 0")
	defer diferido()
	fmt.Println("hoa 1")
}

func diferido() {
	fmt.Println("funcion diferido")
}
*/

//------------Ejercicio 4
/*
type persona struct {
	nombre   string
	apellido string
	edad     int
}

//Adjuntar al tipo "persona" el metodo "presentar"
func (p persona) presentar() {
	fmt.Println("hola soy,", p.nombre, p.apellido)
}

func main() {
	p := persona{
		nombre:   "Victor",
		apellido: "Cruz",
		edad:     30,
	}

	p.presentar()
}
*/
//------------Ejercicio 5

/*
type cuadrado struct {
	lado float64
}

type circulo struct {
	radio float64
}

type forma interface {
	area() float64
}

func (cu cuadrado) area() float64 {
	area := cu.lado * cu.lado
	return area
}

func (ci circulo) area() float64 {
	area := math.Pi * ci.radio * ci.radio
	return area
}

func info(f forma) {
	fmt.Println(f.area())
}

func main() {
	fmt.Println("hola forma")
	cu := cuadrado{lado: 12.35}
	ci := circulo{radio: 15}

	info(cu)
	info(ci)
}
*/
//------------Ejercicio 6
/*
func main() {
	//funcion anonima
	func() {
		x := 100
		fmt.Println(x)
	}()
}
*/

//------------Ejercicio 7
/*
var x int
var f func()

func main() {
	v := func() {
		x := 200
		fmt.Println(x)
	}

	v()
	fmt.Println("listo")

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", f)
}
*/

//------------Ejercicio 8
/*
func funcion() func() {
	return func() {
		fmt.Println("funcion interna")
	}
}

func main() {
	f := funcion()
	f()
}
*/

//------------Ejercicio 9
/*
func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			fmt.Println("por cero")
			return 0
		}
		if len(xi) == 1 {
			fmt.Println("por uno")
			return xi[0]
		}
		fmt.Println(xi[0])
		fmt.Println(xi[len(xi)-1])
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}
*/

//------------Ejercicio 10
func main() {
	//Aparte de asignar la funcion, tambien asigna la variable
	fmt.Println("Valores de f")
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("Valores de g")
	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo() func() int {
	x := 0
	//------clousere
	return func() int {
		x++
		return x
	}
	//------clousere
}
