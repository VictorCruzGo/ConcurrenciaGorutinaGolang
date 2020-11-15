//Package perrito nos permite entender mejor los perros.
package perrito

import "fmt"

//Edad nos devuelve la edad en años perro.
func Edad(e int32) int32 {
	fmt.Println("tu edad es")
	return e * 7
}
