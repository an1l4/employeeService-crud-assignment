package controller

import (
	"encoding/json"
	"net/http"

	"github.com/an1l4/employeeService_crud/db"
	"github.com/gorilla/mux"
)

// slice to store our datas (dummy database)
var employees []db.Employee

// func for adding some sample datas to our dummy database
func Dummydb() {
	employees = append(employees, db.Employee{Id: "fg45", Name: "anila", Location: "Bangaluru"})
	employees = append(employees, db.Employee{Id: "123", Name: "John", Location: "London"})

}

// handler for getting all employees
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)

}

// handler for getting employee by id
func GetById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range employees {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}

}

// handler for adding new employee
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp db.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)

	employees = append(employees, emp)
	w.Write([]byte("Success"))

}
