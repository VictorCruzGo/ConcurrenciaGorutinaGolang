package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
-encodig:codificar y decodificar objetos json
-log: para ver logs
-ht/http: escribir el servidor, peticiones y funcionalidad
-mux

https://www.youtube.com/watch?v=ccdB22kmR_s

*/
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city, omitempty"`
	State string `json:"state, omitempty"`
}

var people []Person

func GetPeopleEndpoint(respuesta http.ResponseWriter, solicitud *http.Request) {
	json.NewEncoder(respuesta).Encode(people)
}

func GetPersonEndpoint(respuesta http.ResponseWriter, solicitud *http.Request) {
	params := mux.Vars(solicitud)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(respuesta).Encode(item)
			return
		}
	}
	json.NewEncoder(respuesta).Encode(&Person{})
}

func CreatePeopleEndpoint(respuesta http.ResponseWriter, solicitud *http.Request) {
	params := mux.Vars(solicitud)
	var person Person
	//para entender lo que el cliente nos envia
	_ = json.NewDecoder(solicitud.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	//Devolviendo todo el arreglo de personas
	json.NewEncoder(respuesta).Encode(people)
}

func DeletePeopleEndpoint(respuesta http.ResponseWriter, solicitud *http.Request) {
	params := mux.Vars(solicitud)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(respuesta).Encode(people)
}

func main() {
	//Creando un enrutador.
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Victor", LastName: "Cruz", Address: &Address{City: "La Paz", State: "La Paz"}})

	people = append(people, Person{ID: "2", FirstName: "Grace", LastName: "Cruz"})

	//endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndpoint).Methods("DELETE")

	//escuchar desde el servidor. puerto y enrutador
	log.Fatal(http.ListenAndServe(":3000", router))
}
