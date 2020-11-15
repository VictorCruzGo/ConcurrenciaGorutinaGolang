package main

import "fmt"

//-------------Arreglo
//Es un bloque constructor
/*func main() {
	var x [10]int
	x[0] = 15
	var y [5]int
	y[3] = 33
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(x))
}*/

//-------------Slice
//Tipo de dato que se construye con el tipo literal compuesto.
//Literal compuesto. Es cualquier ti
//Similar a un arreglo.

/*func main() {
	//Literal compuesto  ":=........"
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	fmt.Printf("%T", x)
}*/

//-------------Slice for range
//Se puede iterar sobre un slice con "range"
/*func main() {
	x := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(x)
	fmt.Println(cap(x)) //cap=capacidad del slice

	for indice, valor := range x {
		fmt.Println(indice, valor)
	}
}*/

//-------------Slice for range
/*func main() {
	x := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(x)
	fmt.Println(x[1:3])
	fmt.Println(x[:])
}*/

//-------------Slice añadiendo un slice
/*func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	y := []int{100, 200, 300}
	fmt.Println(x)
	fmt.Println(y)
	z := append(x, y...)
	fmt.Println(z)
}*/

//-------------Eliminado valores de un slice
/*func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}*/

//-------------Slike make
//Cuando la capacidad de un slice es rabasado, el slice duplica su capacidad.
/*func main() {
	//variable:=[]tipo_dato{valor1,valor2,...}
	//variable:=make([]tipo_dato,longitud, capacidad)
	x := make([]int, 10, 100)
	fmt.Println(cap(x))
	fmt.Println(len(x))

	x[0] = 1
	x[5] = 5

	fmt.Println(x)
}*/

//-------------Slike multidimencional
//multidimencional = fila y columnas (tambien conocidos como matrices)
/*func main() {
	victor := []string{"victor", "cruz", "gomez"}
	grace := []string{"grace adriana", "cruz", "illanes"}

	personas := [][]string{victor, grace}

	fmt.Println(victor)
	fmt.Println(grace)
	fmt.Println(personas)
}*/

//-------------Arreglo subyacente
//Un slice tiene: puntero, logitud y capacidad (subyacente=permanecer oculto debajo de alguna cosa)

/*func main() {
	x := []int{10, 20, 30, 40, 50}
	fmt.Println(x)
	y := append(x[:2], x[3:]...) //Se utiliza el mismo arreglo subyacente
	fmt.Println(y)
	fmt.Println(x)
}*/

//-------------Mapa(map)
//Los mapas son un tipo de almacenamiento llave-valor
//Son tipos de datos desordenados

func main() {
	//mapa=map[tipo_clave]tipo_valor{ llave1:valor1,llave2:valor2,... };
	personas := map[int]string{
		10: "Victor Cruz Gomez",
		20: "Grace Adriana Cruz Illanes",
		30: "Jhon Cruz Calle",
	}

	fmt.Println(personas[10])
	fmt.Println(personas[20])
	fmt.Println(personas[30])
}
