package main

import "fmt"

/*
Punteros
-Con los punteros se puede acceder directamente a la direccion de memoria y cambiar su valor
*/
/*
func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)        //imprimir la direccion de memoria
	fmt.Printf("%T\n", a)  //tipo INT
	fmt.Printf("%T\n", &a) //tipo PUNTERO *INT

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) //accediendo a la direccion de memoria de b, pero sacando el valor
	*b = 50
	fmt.Println(*b)
	fmt.Println(a)
}
*/

/*
Valores que pueden aceptar las funciones.

Receptores   Valores
(t T) 		 T y *T  //T receptor normal. *T receptor del tipo puntero
(t *T)		 *T
*/
func main() {
	fmt.Println("punteros")
}
