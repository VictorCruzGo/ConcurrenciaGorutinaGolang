package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	//http.HandleFunc("/login", login)
	http.ListenAndServe(":3000", nil)
	fmt.Println("fin de la solicitud")
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
