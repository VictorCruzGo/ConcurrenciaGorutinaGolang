package main

//--------Tipo booleano
/*var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
}*/

/*
func main() {
	x := 7
	y := 42
	//fmt.Println(x == y)
	//fmt.Println(x <= y)
	fmt.Println(x != y)
}*/

//--------Tipos numericos
/*var x int

func main() {
	x = 42
	y := 42.32
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
*/

//--------Tipos string
//Un string es una secuencia de bytes de solo lectura.
/*func main() {
	s1 := "hola mundo"
	s2 := `Esta es una  linea
	partida.`
	fmt.Println(s1)
	fmt.Println(s2)

	bs := []byte(s1)
	fmt.Println(bs) //Imprime en ascii lo valores de cada letra del slice
	fmt.Printf("%T", bs)
}*/

//--------Constantes
//Todas la variables son constantes

/*const a = 42
const b = 42.23
const c = "Victor Cruz Gomez"

type nombre string

var d nombre

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	d = c
	fmt.Println(d)
}
*/

//--------Iota
//Constantes usados para incrementar numeros
//Puede ser utilizado para construir un conjunto de constantes relacionadas

/*
const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
*/

//--------bit shifting
//Desplazar binarios a la izquierda o derecha
//Permite mover un bit.

/*func main() {
	a := 4
	a = a << 1
	fmt.Printf("%d\t\t%b", a, a)
}*/

/*
const (
	_  = iota             //iota=0
	kb = 1 << (iota * 10) //iota=1
	gb = 1 << (iota * 10) //iota=2
	tb = 1 << (iota * 10) //iota=3
)

func main() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}*/

//--------bit shifting
