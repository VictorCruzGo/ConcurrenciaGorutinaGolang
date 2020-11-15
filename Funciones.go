package main

import "fmt"

//------------Definicio de una funcion------------
/*
//func(r receptor) identificador(parametros) returno(s){codigo}
//Cuando definimos una funcion llevar parametros
//Cuando utilizamos la funcion la llamamos argumentos
func main() {
	fmt.Println("hola victor")
	foo()
	bar("victor")
	s := woo("Juan")
	fmt.Println(s)
	nombre, resultado := saludar("victor", "cruz")

	fmt.Println(nombre, resultado)
}

func foo() {
	fmt.Println("Hola desde foo")

}

func bar(par string) {
	fmt.Println("hola :", par)
}

func woo(par string) string {
	return fmt.Sprint("Hola desde woo ", par)
}

func saludar(nombre string, apellido string) (string, bool) {
	n := fmt.Sprint(nombre, " ", apellido)
	r := true
	return n, r
}
*/

//------------Parametros variables------------
//func (r receptor) identificador(parametros)(retorno(s)){codigo}

//Utilizar el operado "..." para tener parametros ilimitado
/*
func main() {
	x := sumar(1, 3, 5, 6)
	fmt.Println("El valor de la suma es: ", x)
}

func sumar(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) //slice de enteros

	suma := 0

	for _, v := range x {
		suma += v
	}

	fmt.Println(suma)
	return suma
}
*/

//------------Desplegando un slice------------
//La funcion recibe un parametro del tipo int.
//Para obtener los valores un slice, hacerlo con "..."
//Los valores del slice siempre deberian ir al final

/*
func main() {
	xi := []int{2, 1, 4, 5, 6, 2, 7, 7}

	x := sumar("hola", xi...)
	fmt.Println(x)
}

func sumar(nombre string, x ...int) int {
	fmt.Println(nombre)
	suma := 0
	for _, v := range x {
		suma = suma + v
	}
	return suma
}
*/

//------------Metodos e interfaces I------------
//Si un tipo no recibe ningun parametro, entonces recibe la interfaz vacia

/*
type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona persona
	lpm     bool
}

func (p persona) presentar() {
	fmt.Println("hola, soy ", p.nombre, p.apellido)
}

func (a agenteSecreto) presentar() {
	fmt.Println("Hola, soy el agente ", a.persona.nombre, a.persona.apellido)
}

type humano interface {
	presentar()
}

func bar(h humano) {
	//Afirmacion es igual a verificar el typo de objeto en java
	switch h.(type) {
	case persona:
		fmt.Println("fui pasa a la funcion bar (persona)", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("fui pasa a la funcion bar (agenteSecreto)", h.(agenteSecreto).persona.nombre)

	}
	fmt.Println("fui pasado a la funcion bar ", h)
}

type manzana int

func main() {
	ags1 := agenteSecreto{
		persona: persona{
			nombre:   "victor",
			apellido: "cruz",
		},
		lpm: true,
	}

	ags1.presentar()

	p := persona{
		nombre:   "Maria",
		apellido: "illanes",
	}

	p.presentar()

	//Como person y agente secreto implementan implicitamente la interfac humano,
	//se puede pasar a la funcion bar como argumento.
	//Polimorfismo. La funcion bar recibe varios tipos gracias a las interfaces.
	bar(ags1)
	bar(p)

	fmt.Println(ags1)

	//Conversion
	var x manzana = 42
	fmt.Println(x)
	fmt.Println("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Println("%T\n", y)
}
*/

//------------funciones anonimas------------
//Son funciones que no tienen identificacion
/*
func main() {
	foo()

	//func(){}(), los ultimos parentesis es por la firma de la funcion()
	func() {
		fmt.Println("La funcion anonima se ejecuto")
	}()

	func(x int) {
		fmt.Println("El valor de la fun anonima es ", x)
	}(42)
}

func foo() {
	fmt.Println("foo se ejecuto")
}
*/

//------------expresion funcion------------
//Las funciones son un tipo cualquier como int, bool. Se puede asignar a una variable
/*
func main() {
	f := func() {
		fmt.Println("Mi primera expresion funciones")
	}

	f()

	g := func(x int) {
		fmt.Println("El año que se descubirio america fue ", x)
	}

	g(1492)
}
*/

//------------Retornando uan funcion------------
/*
func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

	//Forma simplicada de llamar a una funcion que retorna otra funcion
	fmt.Println(bar()())
}

func foo() string {
	s := "hola mundo"
	return s
}

func bar() func() int {
	return func() int {
		return 1492
	}
}
*/

//------------callback------------
//Son funciones que se pasan como parametro a otras funciones. f(g(x))
//No se recomienda hacer programacion funcional con go
/*
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8}
	resultado := suma(ii...) //ii es un slice de enterno; ii... son los valores del slice
	fmt.Println(resultado)

	s1 := sumaPares(suma, ii...)
	fmt.Println("el total de pares es ", s1)
}

func suma(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total = total + v
	}

	return total
}

//func sumaPares(firmaFuncion, parametrosVariables)
func sumaPares(f func(yi ...int) int, vi ...int) int {
	//Creando un slice de numeros pares
	var y []int

	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}

	//Retornar evaluando la funcion f(y...)
	return f(y...)
}
*/

//------------Clousure------------
//Clousure, encerramiento o clausura
//Ayudan a limitar el alcance de una variables
//Las variables declaradas en scope externos son accesibles desde escope internos

/*
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)

	{
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y)//Aqui se genera una excepcion porque "y" esta fuera de su scope
}

func foo() {
	fmt.Println(x)
}
*/

/*
func main() {
	//retorna un valor del tipo funcion
	//En "a" tambien se llavara una copia de "x"
	a := incremento()
	b := incremento()
	//Al llamar la funcion "a()" se hace uso de la memoria.
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

//Por tener un clouse, la funcion llevara una copia de "x"
func incremento() func() int {
	var x int
	//Desde aqui hacia abajo es clousure --------
	return func() int {
		x++
		return x
	}
	//--------------------------------------------
}
*/

//------------Recursividad------------
//La recursividad es una funcion que se llama a si misma
//La recursividad puede traer problemas de memoria
//Es mejor utilizar ciclos o bucles para tener codigo legible

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(cicloFactorial(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func cicloFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
