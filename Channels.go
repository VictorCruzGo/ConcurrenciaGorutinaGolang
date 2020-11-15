/*
Channel
-Son un tipo de dato que nos ayuda a escribir codigo concurrente eficiente
-Son como un tubo. Una gorutina envia mensajes atraves del canal a otra gourutina
-El mejor utilizar canales en lugar de escribir en una misma variable
-Se puede convertir un canal general a especifico. Al contrario no.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
func main() {
	//Canal sin buffer. Unbuffered. Se necesitan 2 gourutina. Se necesita otra gourutina que lo reciba.
	ca := make(chan int)

	go func() {
		ca <- 42
	}()

	fmt.Println(<-ca)
}
*/
/*
func main() {
	//Canal con buffer. buffered. No se necesita obligatoriamente otra gorutina
	ca := make(chan int, 1)

	ca <- 42

	fmt.Println(<-ca)
}
*/

/*
func main() {
	//Canal con buffer. buffered. Buffer lleno. Genera un error
	ca := make(chan int, 2) //El chanel acepto dos elementos

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
*/

//Canales direccionables
/*
Canales para leer o escribir
*/

/*
func main() {
	ca := make(chan<- int, 1) //canal de solo envio SEND ONLY
	ca <- 42
	fmt.Println(<-ca)
	fmt.Printf("%T\n", ca)
}
*/

/*
func main() {
	ca := make(<-chan int, 1) //canal de solo recepcion RECIVE ONLY
	ca <- 42
	fmt.Println(<-ca)
	fmt.Printf("%T\n", ca)
}
*/

/*
func main() {

	canalEnvio := make(chan<- int)
	canalRecepcion := make(<-chan int)
	canalBidireccional := make(chan int)

	fmt.Printf("canal bidireccional:::\t%T\n ", canalBidireccional)
	fmt.Printf("canal envio::: \t%T\n ", canalEnvio)
	fmt.Printf("canal recepcion::: \t%T\n  ", canalRecepcion)

	//canalBidireccional = canalRecepcion //no funciona
	//canalBidireccional = canalEnvio	//no funciona
	//canalEnvio = canalBidireccional  //funciona

	fmt.Printf("c\t%T\n", (chan int)(canalEnvio))
	fmt.Printf("c\t%T\n", (chan int)(canalRecepcion))

}
*/

/*
Uso de los canales en funciones
*/
/*
	Pasos en la ejecucion:
	1. Se ejecuta gorutina enviar
	2. Se eejecuta la funcion recibir. Muestra el mensaje recibiendo hasta que llegue un valor en el chanel.
	3. Se ejecuta la funcion recibir nuevamente. cuando llega un valor en al chanel.
*/

/*
func main() {
	c := make(chan int)
	//enviar
	go enviar(c)
	//recibir
	recibir(c)

	fmt.Println("Finalizando...")
}

func enviar(c chan<- int) {
	fmt.Println("Enviando...")
	c <- 10
}

func recibir(c <-chan int) {
	//Esta funcion se queda "esperando" recibir algun valor
	fmt.Println("Recibiendo...")
	fmt.Println(<-c)
}
*/

/*
Range
Utilizar range para leer valores de un canal
*/

/*
func main() {
	c := make(chan int)
	go func() {
		fmt.Println("Dentro de la gorutina")
		for index := 0; index < 6; index++ {
			c <- index
		}
		close(c)
	}()
	fmt.Println("Antes de imprimir")
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Cerrando")
}
*/

//------------ Fan in ------------
/*
Fan in, consisten en agarrar valores de multiples canales que haciendo trabajo independiente, agarra esos valores y se lo envia a un solo canal.
*/
/*
func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	go enviar(par, impar)
	go recibir(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("finalizando")
}

func enviar(p, i chan<- int) {
	fmt.Println("dentro de enviar")
	for index := 0; index < 50; index++ {
		if index%2 == 0 {
			p <- index
		} else {
			i <- index
		}
	}
	close(p)
	close(i)
}

func recibir(p, i <-chan int, fanin chan<- int) {
	fmt.Println("dentro de recibir")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("dentro de recibir 1")
		for v := range p {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("dentro de recibir 2")
		for v := range i {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
*/

//------------ Fan out ------------
/*
-Fan out, consite en agarrar un trabajo que se ejecuta secuencial. dividir el trabajo en poriciones.
Una gourutina va estar encargar de ejecutar partes del trabajo.
-Cada vez que se hace range sobre un canal hay que cerrarlo. Sin se bloquea.
*/
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)
	//go fanOutIn(c1, c2)
	go fanOutInLimitado(c1, c2)

	//Como se hace un range sobre el canal c2, se sabe que el canal esta cerrado.
	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func agregar(c chan int) {
	for index := 0; index < 100; index++ {
		c <- index
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	//Por cada valor de c1 se lanza una gourutina (fan out). Y en cada gourutina se ejecuta un trabajo.
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- trabajoConsumeTiempo(v2)
			wg.Done()
		}(v) //La funcion anonima recibe el parametro "v"
	}
	wg.Wait()
	close(c2)
}

func fanOutInLimitado(c1, c2 chan int) {
	var wg sync.WaitGroup
	const gorutinas = 10
	wg.Add(gorutinas)

	for index := 0; index < gorutinas; index++ {
		go func() {
			//-------------------
			for v := range c1 {
				//func(v2 int) {
				//c2 <- trabajoConsumeTiempo(v2)
				c2 <- trabajoConsumeTiempo(v)
				//}(v)
			}
			//-------------------
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func trabajoConsumeTiempo(n int) int {
	time.Sleep(time.Microsecond + time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
