package router

import (
	"github.com/gorilla/mux"
	"github.com/parthverma/CRUDapplication/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/employees", controller.GetAllEmployees).Methods("GET")
	router.HandleFunc("/employee", controller.InsertOneEmployees).Methods("POST")
	router.HandleFunc("/employee/{id}", controller.DeleteOneEmployee).Methods("DELETE")
	router.HandleFunc("/employee/{id}", controller.UpdateOneEmployees).Methods("PUT")
	router.HandleFunc("/employeesDelete", controller.DeleteAllEmployee).Methods("DELETE")

	return router
}
