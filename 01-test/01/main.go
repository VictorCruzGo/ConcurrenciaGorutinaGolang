package main

import "fmt"

/*
Testing,
-Es un paquete.
-La logica del testing ira haciendo un chequeo manual del codigo.
-El codigo evoluciona. Si la prueba falla, hubo un fallo.
-Los test terminann en : xxxx_test.go
-Los test deben estar en el mismo paquete.
-Los test deben estar en una funcion con la firma: TestXxxx(t *testing.T)
-Para correr las pruebas: >go test.
-Los errores tratar con t.Error
*/

func main() {
	fmt.Println("La suma de 2 + 4 es: ", miSuma(2, 4))
	fmt.Println("La suma de 1 + 5 es: ", miSuma(1, 5))
	fmt.Println("La suma de 6 + 7 es: ", miSuma(6, 7))
}

func miSuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

//--------------Tablas de pruebas
/*
Tablas de pruebas
-Agarrar un pruebas y repetirla varias veces.

*/
