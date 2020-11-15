package main

import (
	"fmt"
	"runtime"
)

//------------Ejercicio 1
/*
var wg sync.WaitGroup

func gourutina1() {
	fmt.Println("soy la gourutina 1")
	wg.Done()
}

func gourutina2() {
	fmt.Println("soy la gourutina 2")
	wg.Done()
}

func main() {
	fmt.Println("numero de NumCPU:", runtime.NumCPU())
	fmt.Println("numero de NumGoroutine:", runtime.NumGoroutine())
	fmt.Println("numero de GOARCH:", runtime.GOARCH)
	fmt.Println("numero de GOOS:", runtime.GOOS)

	fmt.Println("soy la gourutina principal")
	wg.Add(3)

	go gourutina1()
	go gourutina2()

	go func() {
		fmt.Println("soy la gourutina anonima")
		wg.Done()
	}()
	wg.Wait()
}
*/

//------------Ejercicio 2
/*
	(t T) 		T y *T
	(t *T) 		*T
*/

/*
type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p *persona) hablar() {
	fmt.Println("Hola puedo hablar, mi nombre es: ", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p := persona{
		nombre:   "victor",
		apellido: "cruz",
		edad:     30,
	}

	diAlgo(&p)
	p.hablar() //Funciona porque la variable T es direccionable, tiene asociado un *T.
	//diAlgo(p) //Esta linea no corresponde por que el receptor es del tipo puntero.
}
*/

//------------Ejercicio 3 y 4

/*
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	incremento := 0

	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			variable := incremento
			runtime.Gosched() //Cede el procesador a otro hilo. Con el mutex no es necesario hacer yield
			variable++
			fmt.Println(variable)
			incremento = variable
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final del incremento es: ", incremento)
}
*/

//------------Ejercicio 5
/*
atomic, como estoy leyendo y escribiendo sobre una variable es considerado bajo nivel y atomico.
Leer y escribir sobre una variable es considerado atomico
*/
/*
func main() {
	var wg sync.WaitGroup
	var incremento int64

	incremento = 0

	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&incremento, 2)
			fmt.Println(atomic.LoadInt64(&incremento))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final del incremento es: ", incremento)
}
*/

//------------Ejercicio 6

func main() {
	fmt.Println("El sistema operativo es: ", runtime.GOOS)
	fmt.Println("La arquitectura es: ", runtime.GOARCH)
}
