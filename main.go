package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/an1l4/employeeService_crud/controller"
	"github.com/gorilla/mux"
)

func main() {
	//creating new router using mux
	r := mux.NewRouter()
	controller.DummyData()

	//adding some data to our dummy database to avoid to get empty data
	//controller.Dummydb("fg45", "anila", "Bangaluru")

	//router for getting all employee details from dummy db
	r.HandleFunc("/employees", controller.GetEmployee).Methods("GET")

	//router for getting employee details from dummy db using id
	r.HandleFunc("/employee/{id}", controller.GetById).Methods("GET")

	//router for adding employee details to the dummy db
	r.HandleFunc("/employee", controller.AddEmployee).Methods("POST")

	r.HandleFunc("/employee/{id}", controller.UpdateEmployee).Methods("PUT")

	r.HandleFunc("/employee/{id}", controller.DeleteEmployee).Methods("DELETE")

	fmt.Println("server running at 8080...")

	//starting server
	log.Fatal(http.ListenAndServe(":8080", r))
}
