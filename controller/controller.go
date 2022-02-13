package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parthverma/CRUDapplication/model"
	"github.com/parthverma/CRUDapplication/service"
)

func InsertOneEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var employee model.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Fatal(err)
	}

	inserted := service.InsertOneEmployee(employee)
	fmt.Printf("Inserted with id: %s.\n", inserted)
	json.NewEncoder(w).Encode(employee)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allEmployees := service.GetAllEmployees()
	fmt.Println("Got all employees.")
	json.NewEncoder(w).Encode(allEmployees)
}

func UpdateOneEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var employee model.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)
	result := service.UpdateOneEmployee(employee, params["id"])
	fmt.Printf("Modified %d records.\n", result)
	json.NewEncoder(w).Encode(employee)
}

func DeleteOneEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	result := service.DeleteOneEmployee(params["id"])
	fmt.Printf("Deleted %d records.\n", result)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	result := service.DeleteAllEmployees()
	fmt.Printf("Deleted %d records.\n", result)
}
