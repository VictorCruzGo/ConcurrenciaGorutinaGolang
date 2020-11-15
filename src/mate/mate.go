//Package mate nos ayuda a comprobar que sabe sumar
package mate

//Sum suma una cantidad ilimida de numeros enteros.
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}
