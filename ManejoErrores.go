package main

import "fmt"

/*
Porque go no tiene excepciones?
-Los dev suelen etiquetar errores comunes como excepcionales. Ej. no abrir un archivo.
-Los retornos multivalores en go facilitan el reporte de un error sin sobrecargar la respuesta de la funcion.
-La verificacion del error se realiza justo despues de ejecutar la misma.
-El paquete built-in se encuentra un tipo error.
-Un error es un tipo cualquiera.
*/

//---------------Chequeando errores
/*
Chequea siempre si hay errores.
*/

/*
func main() {
	n, err := fmt.Println("hola")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
*/

/*
func main() {
	var respuesta1, respuesta2, respuesta3 string

	fmt.Println("Nombre: ")
	_, err := fmt.Scan(&respuesta1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Comida favorita: ")
	_, err = fmt.Scan(&respuesta2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deporte favorito: ")
	_, err = fmt.Scan(&respuesta3)
	if err != nil {
		panic(err)
	}

	fmt.Println(respuesta1, respuesta2, respuesta3)
}
*/

/*
func main() {
	f, err := os.Create("nombres.txt") //Creando un archivo
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond") //Escribiendo en el archivo
	io.Copy(f, r)
}
*/

/*
func main() {
	f, err := os.Open("nombres.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f) //Convierte a slice de byte lo que tiene el archivo nombres.txt
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
*/

//---------------Printing & Logging
/*
log.Println(), se puede especificar que el log se guarde en un archivo.
log.Fatallin()
log.Panicln()
panic()
*/

/*
func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		fmt.Println("ocurrio un error", err)

	}
}
*/

/*
func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("un erro ocurrio", err)
	}
}
*/

/*
func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	log.SetOutput(f) //Configurando el archivo para escribir los mensajes.

	f2, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("un erro ocurrio", err) //no se imprimira en pantalla el menajes, se imprimira en el archivo.
	}
	defer f2.Close()

	fmt.Println("Revisar el archivo log.txt en el directorio")
}
*/

/*
func main() {
	defer foo()
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err) //Todo muere. Llama  os.exit(1)
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}
*/

/*
Panic
-Para la funcion normal de la gorutina.
-Termina cualquier ejecucion que sea diferida.
-
*/

/*
func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Panicln(err)
		//panic(err)
	}
}
*/

//---------------Recover
/*
defer - panic - recover
-Son mecanismos de control de flujo.
----defer----
-Utilizado para realizar tareas de limpieza, como cerrar archivos o conexiones.
----panic----
-Una especie de excepcion en tiempo de ejecucion. La excepcion la podemos interpretar con recover y obtener el valor.
-Con la ayuda de defer podemos interceptar la llamada a panic.
-En el caso de producirse un panic el programa se detendra desde la funcion donde se origino hacia arriba.
-Detiene la ejecucion de la misma y termina tambian con la ejecucion de la funcion main.
----recover----
-Es una funcion interna que toma el control de una gorutina.
-Al convinar con panic, podemos evitar que termine la ejecucion de nuestro programa.

Ej:

func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()

    _, err := os.Create("/file")
    if err != nil {
        panic(err)
    }
}

*/
func main() {
	fmt.Println("hola mundo")
}
