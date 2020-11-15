package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/*
Json
-Formato ligerod de intercambio de datos.
-marshal. Para codificarlo en Json
-unmarshal. Para decodificarlo en formato Go.
-La interfaz vacias es cualquier tipo de dato.
-Buena practica es controlar el error despues de la invocacion del metodo.
-Para convertir a objeto Json un Struct los atributos tienen que iniciar en mayusculas.
-Pagina para convertir json a go: https://mholt.github.io/json-to-go/
-Etiquetas `json:"Nombre"`. Se utiliza para mapear
*/

//--------------Marshal
/*
type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := persona{
		Nombre:   "Victor",
		Apellido: "Cruz",
		Edad:     32,
	}

	p2 := persona{
		Nombre:   "Grace",
		Apellido: "Cruz",
		Edad:     7,
	}

	personas := []persona{p1, p2}

	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

}
*/
//--------------Unmarshal
/*
type persona struct {
	Nombre   string `json:"Nombrecito"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombrecito":"Victor","Apellido":"Cruz","Edad":32},{"Nombre":"Grace","Apellido":"Cruz","Edad":7}]`
	bs := []byte(s)
	fmt.Printf("%T\n", bs)

	var personas []persona

	err := json.Unmarshal(bs, &personas)
	if err != nil { //Si es distinto es porque paso algo
		fmt.Println(err)
	}

	//Al ejecutar el unmarshal, el atribujo Nombre del json no se mapea
	fmt.Println("Toda la data ", personas)

}
*/

//--------------Interfaz writer
/*
-Utilizamos marsha y unmarshal cuando se desea manipular el json. (variables)
-Utilizamos enconder y decoder cuando se quiere manipular json directo en cable. (directo)

Interfaz writer
-

Paquete os
-Create, para crear archivos.
-Un metodo de os implementar la interfaz write por lo tanto directamente se puede hacer encdeo

Println
-Escribe la salida estandar
-En la implementacion utiliza Fprintln(os.Stdwon)

Fprintln
-Recibe como primer parametro io.Writer

Os.Stdout
-Stdout es un variable que tiene como valor una funcion NewFile
-El Stdout implementa la interfaz writer

NewFiles
-Retorna un puntero a archivo

io
-

-Funciones que retorna valores de tipo writer, o que implementen la interfaz writer
-Como se conectan unas a otras, hay que ver en que interfaz se encuentra la interfaz.
*/

/*
func main() {
	fmt.Println("hola")
	fmt.Fprintln(os.Stdout, "hola 2")
	io.WriteString(os.Stdout, "hola 3")

}
*/

//--------------Sort
/*
func main() {
	xi := []int{4, 5, 6, 42, 99, 18}
	xs := []string{"a", "z", "b", "x", "y"}
	//sort no retorna nada. Ints ordena en memoria.
	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)
}
*/

//--------------Sort personalizado
/*
type Persona struct {
	Nombre string
	Edad   int
}

type PorEdad []Persona

type PorNombre []Persona

func (a PorEdad) Len() int {
	return len(a)
}

func (a PorEdad) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PorEdad) Less(i, j int) bool {
	return a[i].Edad < a[j].Edad
}

func (a PorNombre) Len() int {
	return len(a)
}

func (a PorNombre) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PorNombre) Less(i, j int) bool {
	return a[i].Nombre < a[j].Nombre
}

func main() {
	p1 := Persona{Nombre: "Eduar", Edad: 32}
	p2 := Persona{Nombre: "Maria", Edad: 25}
	p3 := Persona{Nombre: "Victor", Edad: 56}
	p4 := Persona{Nombre: "Grace", Edad: 60}

	personas := []Persona{p1, p2, p3, p4}
	fmt.Println("No ordenado-----------")
	fmt.Println(personas)

	fmt.Println("Por edad-----------")
	sort.Sort(PorEdad(personas))
	fmt.Println(personas)

	fmt.Println("Por nombre-----------")
	//Sort implementa el metodo quickSort que a su vez implementa
	//los metodos de la interfaz 	Len(), Swap() y Less()
	sort.Sort(PorNombre(personas))
	fmt.Println(personas)
}
*/

//--------------Bcrypt
/*
Bcrypt
Es una de la herramientas que ayudan a encriptar
*/

func main() {
	s := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	clave := `clave2014`
	err = bcrypt.CompareHashAndPassword(bs, []byte(clave))
	if err != nil {
		fmt.Println("NO te puede loguear")
		return
	}
	fmt.Println("Te haz logueado")

}
