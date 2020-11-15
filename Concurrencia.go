package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//--------------------Condition race
/*
-Yield cede el pasa a otro hilo.
-Data race es cuando tenemos varias gorutinas accediendo a una variable. No es bueno tener data race.
>go run -race .\Concurrencia.go
-Mutex, bloquea el acceso de lectura y escritura cuando la gorutina esta haciendo la tarea.
*/

/*
func main() {
	fmt.Println("Numero de CPUs,", runtime.NumCPU())
	fmt.Println("Numero de Gorutinas primera vez,", runtime.NumGoroutine())
	contador := 0
	const gs = 100        //numero de go rutinas
	var wg sync.WaitGroup //

	wg.Add(gs)
	var mu sync.Mutex //
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			//time.Sleep(time.Second * 2)
			runtime.Gosched() //yield
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Numero de Gorutinas----,", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("cuenta: ", contador)
}

*/
//--------------------Condition race
/*
-Atomic, funciones que implementan algoritmos de sincronizacion para evitar el bloqueo.
*/
func main() {
	fmt.Println("Numero de CPUs,", runtime.NumCPU())
	fmt.Println("Numero de Gorutinas primera vez,", runtime.NumGoroutine())
	var contador int64

	const gs = 100        //numero de go rutinas
	var wg sync.WaitGroup //
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched() //yield
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Numero de Gorutinas----,", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("cuenta: ", contador)
}
