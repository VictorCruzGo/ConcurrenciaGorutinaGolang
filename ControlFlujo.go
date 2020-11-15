package main

import "fmt"

//-------------Ciclo for
/*func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}*/

//-------------Ciclo for anidados
/*func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("El valor del ciclo externo es %d\n", i)
		for j := 0; j <= 5; j++ {
			fmt.Printf("\tEl valor del ciclo interno es %d\n", j)
		}
	}
}*/

//-------------Ciclo declaracion for infinito
/*func main() {
	i := 1

	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}*/

//-------------Break y continuo
/*func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
			//hace un reset de la iteracion
		}

		fmt.Println(x)
	}
}*/

//-------------Ciclo imprimiendo el valor ascii
/*func main() {
	for i := 33; i <= 123; i++ {
		fmt.Printf("%d\t%x\t%#U\n", i, i, i)
	}
}*/

//-------------Declaracion if
/*func main() {
	if true {
		fmt.Println(1000)
	}

	if false {
		fmt.Println(2000)
	}

	if 10 == 10 {
		fmt.Println(3000)
	}

	if x := 20; x == 20 {
		fmt.Println(4000)
	}
}*/

//-------------Declaracion if-elseif-else
/*func main() {
	x := 41
	if x == 40 {
		fmt.Println("El valor es 40")
	} else if x == 41 {
		fmt.Println("El valor de x es 41")
	} else {
		fmt.Println("El valor de x no es 40")
	}
}*/

//-------------Declaracion switch
/*func main() {
	x := 100
	switch x {
	case 1:
		fmt.Println("El valor es 1")
	case 10:
		fmt.Println("El valor es 10")
	default:
		fmt.Println("Por defecto")
	}
}*/

//-------------Operadores logicos
func main() {
	fmt.Println(true == true)
	fmt.Println(false == false || true == true)
	fmt.Println(false && true)
}
