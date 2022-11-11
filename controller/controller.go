package controller

import (
	//"crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// slice to store our datas (dummy database)
var employees = make(map[string]Employee)

// func for adding some sample datas to our dummy database
func DummyData() {
	fmt.Println(employees)
	employees["1"] = Employee{"123", "anila", "bglr"}
	employees["2"] = Employee{"321", "John", "London"}
	fmt.Println(employees)

	//d.Id = append(Id, id)
	//	d.Name = append(d.Name, name)
	//	d.Location = append(d.Location, location)

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
			fmt.Println(item.Id)
			json.NewEncoder(w).Encode(item)
			return

		}
	}

}

// handler for adding new employee
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	fmt.Println(emp)
	rand.Seed(time.Now().UnixNano())
	key := strconv.Itoa(rand.Intn(1000000000))

	//employees = append(employees, emp)
	employees[key] = emp
	w.Write([]byte("Success"))

}
