package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API del grupo 16!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/productos", returnAllProducts)

	myRouter.HandleFunc("/producto", createNewProduct).Methods("POST")
	myRouter.HandleFunc("/producto/{id}", deleteProduct).Methods("DELETE")
	myRouter.HandleFunc("/producto/{id}", updateProduct).Methods("PUT")
	myRouter.HandleFunc("/producto/{id}", returnSingleProduct)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Productos = []Producto{
		Producto{Id: "1", Nombre: "Leche", Descripcion: "Leche descremada 1L", Valor: "950", Expiracion: "19/07/2022"},
		Producto{Id: "2", Nombre: "Torta", Descripcion: "Torta de chocolate 10 p", Valor: "7990", Expiracion: "29/12/2022"},
	}
	handleRequests()

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Productos)
}

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]


	for _, producto := range Productos {
		if producto.Id == key {
			json.NewEncoder(w).Encode(producto)
		}
	}
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var producto Producto
	json.Unmarshal(reqBody, &producto)
	Productos = append(Productos, producto)

	json.NewEncoder(w).Encode(producto)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	id := vars["id"]

	for index, producto := range Productos {
		if producto.Id == id {
			Productos = append(Productos[:index], Productos[index+1:]...)
		}
	}

}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEvent Producto
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, producto := range Productos {
		if producto.Id == id {

			producto.Nombre = updatedEvent.Nombre
			producto.Descripcion = updatedEvent.Descripcion
			producto.Valor = updatedEvent.Valor
			producto.Expiracion = updatedEvent.Expiracion
			Productos[i] = producto
			json.NewEncoder(w).Encode(producto)
		}
	}
}

type Producto struct {
	Id          string `json:"Id"`
	Nombre      string `json:"Nombre"`
	Descripcion string `json:"descripcion"`
	Valor       string `json:"valor"`
	Expiracion  string `json:"expiracion"`
}


var Productos []Producto