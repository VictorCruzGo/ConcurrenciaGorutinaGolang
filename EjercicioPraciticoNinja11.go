package main

import (
	"fmt"
	"log"
)

//-----------Ejercicio 1-----------
/*
type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "Victor",
		Apellido: "Cruz",
		Frases:   []string{"shaeken, not stirred", "Algun ultimo deseo?", "Nunca digas nunca."},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		//fmt.Println(err)
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
*/

//-----------Ejercicio 2-----------
/*
type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "Victor",
		Apellido: "Cruz",
		Frases:   []string{"frase 1", "frase 2", "frase 3"},
	}

	bs, err := aJSON(p1)
	if err != nil {
		log.Fatalln(err) //Mata todo, funciones, gorutinas, etc.
	}
	fmt.Println(string(bs))
}

func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return bs, fmt.Errorf("ocurrio un error al convertir a json %v", err)
		return []byte{}, errors.New(fmt.Sprintf("ocurrio un error al convertir a json %v", err))
	}
	return bs, nil
}
*/
//-----------Ejercicio 3-----------
/*
type errorPer struct {
	info string
}

func (e errorPer) Error() string {
	return fmt.Sprintf("El error es, %v", e.info)
}

func main() {
	e := errorPer{
		info: "-errores info-",
	}
	foo(e)
	fmt.Println(e)
}

func foo(e error) {
	fmt.Println("foo corrio -", e, ":\n", e.(errorPer).info) //asercion, afirmar explicitamente que es de ese tipo
}
*/

//-----------Ejercicio 4-----------
type raizError struct {
	lat  string
	long string
	err  error
}

//error() string metodo del paquet built-in
func (r raizError) Error() string { //Error() en mayusculas
	return fmt.Sprintf("Error matematica: %v %v %v", r.lat, r.long, r.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		//Escribe tu codigo aqui
		e := fmt.Errorf("El valor era %v", f)
		//e := errors.New("El valor es menor que cero")
		return f, raizError{lat: "9999 N", long: "11111 S", err: e}
	}
	return f, nil
}
