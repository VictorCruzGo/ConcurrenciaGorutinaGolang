package mate

import "fmt"

/*
-Con Example se genera un ejemplo para la documentacion, a la vez se pueden realizar pruebas.
-Correr todos los test que se encuentra en un directorio
>go test ./...
*/
func ExampleSum() {
	fmt.Println(Sum(2, 4, 5))
	//Output:
	//11
}
